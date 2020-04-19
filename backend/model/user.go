package model

import (
    "net/http"
    "encoding/json"
    "github.com/labstack/echo"
    "golang.org/x/crypto/bcrypt"
    "github.com/yellia1989/tex-go/tools/util"
    cm "github.com/yellia1989/tex-web/backend/common"
)

var users *cm.Map

func init() {
    bs, _ := util.LoadFromFile("data/users.json")
    items := make([]*User,0)
    json.Unmarshal(bs, &items)

    items2 := make([]cm.Item,0)
    for _, item := range items {
        items2 = append(items2, item)
    }
    users = cm.NewMap("data/users.json", items2)
}

type User struct {
    Id uint32           `json:"id"`
    UserName string     `json:"username"`
    Password string     `json:"password"`
    Role uint32         `json:"role"`
}

func (u *User) GetId() uint32 {
    return u.Id
}
func (u *User) SetId(id uint32) {
    u.Id = id
}
func (u *User) IsAdmin() bool {
    return u.Role == 1
}
func (u *User) CheckPermission(path string, method string) error {
    if u.IsAdmin() {
        return nil
    }

    r := GetRole(u.Role)
    if r == nil {
        return &echo.HTTPError{
            Code:   http.StatusForbidden,
            Message: "请给用户指定一个角色",
        }
    }

    for _, p := range r.Perms {
        perm := GetPerm(p)
        if perm == nil {
            return &echo.HTTPError{
                Code:   http.StatusForbidden,
                Message: "用户的角色不存在",
            }
        }

        if perm.checkPermission(method, path) {
            return nil
        }
    }

    return &echo.HTTPError{
        Code:   http.StatusForbidden,
        Message: "没有对应的权限",
    }
}
func (u *User) ComparePwd(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}
func (u *User) EncodePwd(password string) bool {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return false
    }

    u.Password = string(hash)
    return true
}

func GetUser(id uint32) *User {
    if users == nil {
        return nil
    }

    u := users.GetItem(id)
    if u == nil {
        return nil
    }
    // 复制一份防止原始值被修改
    u2 := *(u.(*User))
    return &u2
}

func GetUserByUserName(username string) *User {
    if users == nil {
        return nil
    }

    // username不能为空
    if len(username) == 0 {
        return nil
    }

    items := users.GetItems(func (key, v interface{})bool{
        u := v.(*User)
        return u.UserName == username
    })
    if len(items) == 0 {
        return nil
    }
    if len(items) != 1 {
        panic("username duplicate")
    }

    // 复制一份防止原始值被修改
    u2 := *(items[0].(*User))
    return &u2
}

func GetUsers() []*User {
    if users == nil {
        return nil
    }

    items := users.GetItems(func (key, v interface{})bool{
        return true
    })

    if len(items) == 0 {
        return nil
    }

    us := make([]*User,0)
    for _, item := range items {
        // 复制一份防止原始值被修改
        u := *(item.(*User))
        us = append(us, &u)
    }
    return us
}

func AddUser(username string, password string, role uint32) *User {
    if users == nil {
        return nil
    }

    // username,password不能为空
    if len(username) == 0 || len(password) == 0 {
        return nil
    }
    // role必须存在
    if GetRole(role) == nil {
        return nil
    }
    // username不能相同
    if GetUserByUserName(username) != nil {
        return nil
    }

    u := &User{UserName: username, Role: role}
    if !u.EncodePwd(password) {
        return nil
    }
    if !users.AddItem(u) {
        return nil
    }

    // 复制一份防止原始值被修改
    u2 := *u
    return &u2
}

func DelUser(u *User) bool {
    if users == nil {
        return false
    }

    return users.DelItem(u)
}

func UpdateUser(u *User) bool {
    if users == nil {
        return false
    }

    u2 := *u
    return users.UpdateItem(&u2)
}

func DelAllUser() bool {
    if users == nil {
        return false
    }

    return users.DelAllItem()
}

func DelUserRole(roles []uint32) {
    items := users.GetItems(func (key, v interface{})bool{
        u := v.(*User)
        if util.Contain(roles, u.Role) {
            return true
        }
        return false
    })

    if len(items) == 0 {
        return
    }
    for _, item := range items {
        u := *(item.(*User))
        u.Role = 0
        users.UpdateItem(&u)
    }
}
