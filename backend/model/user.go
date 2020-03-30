package model

import (
    "sync"
    "sync/atomic"
    "sort"
    "encoding/json"
    "github.com/casbin/casbin"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/util"
)

var ce *casbin.Enforcer
var users userMap
var uid uint32

func init() {
    var err error
    ce, err = casbin.NewEnforcer("data/auth_model.conf", "data/auth_policy.csv")
    _ = err

    users.init("data/users.json")
}

type User struct {
    Id uint32            `json:"id"`
    Name string         `json:"name"`
    UserName string     `json:"username"`
    Password string     `json:"password"`
}

func (u *User) CheckPermission(path string, method string) error {
    if ce == nil {
        return nil
    }

    pass, err := ce.Enforce(u.UserName, path, method)
    if err != nil {
        return err
    }
    if !pass {
        return echo.ErrForbidden
    }
    return nil
}

type userMap struct {
    users sync.Map
    path string
}

func (um *userMap) getUser(id uint32) *User {
    v, ok := um.users.Load(id)
    if !ok {
        return nil
    }
    // 复制一个防止返回的user被修改
    u := *(v.(*User))
    return &u
}

func (um *userMap) getUserByUserName(username string) *User {
    var user *User
    um.users.Range(func (key, v interface{}) bool {
        u := v.(*User)
        if u.UserName == username {
            user = u
            return false
        }
        return true
    })
    return user
}

type userById []*User
func (a userById) Len() int           { return len(a) }
func (a userById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a userById) Less(i, j int) bool { return a[i].Id < a[j].Id }

func (um *userMap) getUsers() []*User {
    us := make([]*User,0)
    um.users.Range(func (key, v interface{}) bool {
        us = append(us, v.(*User))
        return true
    })

    sort.Sort(userById(us))
    return us
}

func (um *userMap) addUser(u *User) bool {
    // id不能相同
    if um.getUser(u.Id) != nil {
        return false
    }

    um.users.Store(u.Id, u)
    return um.save() == nil
}

func (um *userMap) delUser(u *User) bool {
    _, ok := um.users.Load(u.Id)
    if !ok {
        return false
    }
    um.users.Delete(u.Id)
    return um.save() == nil
}

func (um *userMap) delAllUser() bool {
    us := um.getUsers()
    for _, u := range us {
        um.users.Delete(u.Id)
    }
    return um.save() == nil
}

func (um *userMap) updateUser(u *User) bool {
    if um.getUser(u.Id) == nil {
        return false
    }
    um.users.Store(u.Id, u)
    return um.save() == nil
}

func (um *userMap) save() error {
    us := um.getUsers()

    bs, err := json.Marshal(us)
    if err != nil {
        return err
    }
    return util.SaveToFile(um.path, bs, false)
}

func (um *userMap) init(path string) error {
    um.path = path
    bs, err := util.LoadFromFile(path)
    if err != nil {
        return err
    }
    users := make([]*User,0)
    err = json.Unmarshal(bs, &users)
    if err != nil {
        return err
    }
    for _, u := range users {
        um.users.Store(u.Id, u)
        if u.Id > uid {
            uid = u.Id
        }
    }
    return nil
}

func GetUser(id uint32) *User {
    return users.getUser(id)
}

func GetUserByUserName(username string) *User {
    return users.getUserByUserName(username)
}

func GetUsers() []*User {
    return users.getUsers()
}

func AddUser(name string, username string, password string) *User {
    // username不能相同
    if users.getUserByUserName(username) != nil {
        return nil
    }

    u := &User{Id: atomic.AddUint32(&uid, 1), Name: name, UserName: username, Password: password}
    if !users.addUser(u) {
        return nil
    }
    return u
}

func DelUser(u *User) bool {
    return users.delUser(u)
}

func UpdateUser(u *User) bool {
    return users.updateUser(u)
}

func DelAllUser() bool {
    return users.delAllUser()
}
