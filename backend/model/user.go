package model

import (
    "fmt"
    "github.com/labstack/echo/v4"
    "github.com/yellia1989/tex-web/backend/cfg"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

type User struct {
    Id uint32           `json:"id"`
    UserName string     `json:"username"`
    Password string     `json:"password"`
    Role uint32         `json:"role"`
    NeedReLogin uint32
    AllowGmCmd string    `json:"allowGmCmd"`
    TerminalUser string  `json:"terminalUser"`
    TerminalKey string   `json:"terminalKey"`
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

func (u *User) IsNeedLogin() bool {
    return u.NeedReLogin != 0
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

func (u *User) CheckGmPermission(cmd string) bool {
    if u.IsAdmin(){
        return true
    }
    reg := regexp.MustCompile(`^([a-zA-z_]*)\s*`)
    res :=reg.FindStringSubmatch(cmd)
    if len(res) < 2 {
        return false
    }
    cmd = res[1]
    cmdArr := strings.Split(u.AllowGmCmd,"\n")
    for _,v := range cmdArr {
        if v == cmd{
            return true
        }
    }
    return false
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
    db := cfg.StatDb
    if db == nil {
        return nil
    }
    user := &User{}
    if err:=db.QueryRow("select id,username,password,role,need_login,allow_gm_cmd,terminal_user,terminal_key from sys_user where id = ?",id).Scan(&user.Id,&user.UserName,&user.Password,&user.Role,&user.NeedReLogin,&user.AllowGmCmd,&user.TerminalUser,&user.TerminalKey);err!=nil{
        return nil
    }
    return user
}

func GetUserByUserName(username string) *User {
    // username不能为空
    if len(username) == 0 {
        return nil
    }
    db := cfg.StatDb
    if db == nil {
        return nil
    }
    user := &User{}
    if err:=db.QueryRow("select id,username,password,role,need_login,allow_gm_cmd,terminal_user,terminal_key from sys_user where username = ?",username).Scan(&user.Id,&user.UserName,&user.Password,&user.Role,&user.NeedReLogin,&user.AllowGmCmd,&user.TerminalUser,&user.TerminalKey);err!=nil{
        return nil
    }
    return user
}

func GetUsers() []*User {
    db := cfg.StatDb
    if db == nil {
        return nil
    }
    rows, err := db.Query("select id,username,password,role,need_login,allow_gm_cmd,terminal_user,terminal_key from sys_user")
    if err!=nil{
        return nil
    }
    us := make([]*User,0)
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Id,&user.UserName,&user.Password,&user.Role,&user.NeedReLogin,&user.AllowGmCmd,&user.TerminalUser,&user.TerminalKey);err!=nil{
            return nil
        }
        us = append(us,&user)
    }
    return us
}

func AddUser(username string, password string, role uint32,allowGmCmd string) *User {
    db := cfg.StatDb
    if db == nil {
        return nil
    }
    // username,password不能为空
    if len(username) == 0 || len(password) == 0 {
        return nil
    }
    // role必须存在
    if role != 0 && GetRole(role) == nil {
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
    _,err := db.Exec("insert into sys_user(username,password,role,allow_gm_cmd) values(?,?,?,?)",u.UserName,u.Password,u.Role,allowGmCmd)
    if err!=nil{
        return nil
    }
    return u
}

func DelUser(u *User) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    _,err := db.Exec("delete from sys_user where id = ?",u.Id)
    if err!=nil{
        return false
    }
    return true
}

func ResetUserNeedReLogin(id uint32) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    _,err := db.Exec("update sys_user set need_login = 0 where id = ?",id)
    if err!=nil {
        return false
    }
    return true
}

func UpdateUser(u *User) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    _,err := db.Exec("update sys_user set password = ?,role = ?,need_login = 1,allow_gm_cmd=?,terminal_user=?,terminal_key=? where id = ?",u.Password,u.Role,u.AllowGmCmd,u.TerminalUser,u.TerminalKey,u.Id)
    if err!=nil {
        return false
    }
    return true
}

func DelAllUser() bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    _,err := db.Exec("delete from sys_user")
    if err!=nil {
        return false
    }
    return true
}

func DelUserRole(roles []uint32) {
    db := cfg.StatDb
    if db == nil {
        return
    }
    roleStr := make([]string,len(roles))
    for _,v := range roles{
       roleStr = append(roleStr,strconv.FormatUint(uint64(v),10))
    }
    sql := "update sys_user set role = 0 where role in (%s)"
    ids := strings.Join(roleStr,",")
    sql = fmt.Sprintf(sql,ids)
    _,err := db.Exec(sql)
    if err!=nil {
        return
    }
}
