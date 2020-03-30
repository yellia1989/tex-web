package model

import (
    "encoding/json"
    "github.com/casbin/casbin"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/util"
    cm "github.com/yellia1989/tex-web/backend/common"
)

var ce *casbin.Enforcer
var users *cm.Map

func init() {
    ce, _ = casbin.NewEnforcer("data/auth_model.conf", "data/auth_policy.csv")
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
    Name string         `json:"name"`
    UserName string     `json:"username"`
    Password string     `json:"password"`
}

func (u *User) GetId() uint32 {
    return u.Id
}
func (u *User) SetId(id uint32) {
    u.Id = id
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

func AddUser(name string, username string, password string) *User {
    if users == nil {
        return nil
    }

    // username,password不能为空
    if len(username) == 0 || len(password) == 0 {
        return nil
    }
    // username不能相同
    if GetUserByUserName(username) != nil {
        return nil
    }

    u := &User{Name: name, UserName: username, Password: password}
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
