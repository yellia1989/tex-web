package sys

import (
    "github.com/gorilla/sessions"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/session"
    "github.com/yellia1989/tex-web/backend/cfg"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "time"
)

func UserLogin(c echo.Context) error {
    ctx := c.(*mid.Context)
    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    u := model.GetUserByUserName(username)
    if u == nil || !u.ComparePwd(password) {
        return ctx.SendError(-1, "用户名或密码输入错误")
    }

    if ctx.GetUserId() != 0 && !u.IsNeedLogin() {
        return ctx.SendError(-2, "已经登陆不用重新登录")
    }
    if u.IsNeedLogin() {
       success := model.ResetUserNeedReLogin(u.Id)
       if !success{
           return ctx.SendError(-3, "重置登录状态失败")
       }
    }
    expire := time.Hour * 24 * 7
    sess, err := session.Get("texweb_session", c)
    if err != nil {
        return err
    }
    sess.Options = &sessions.Options{
        Path : "/",
        MaxAge : int(expire.Seconds()),
    }

    sess.Values["userid"] = strconv.FormatUint(uint64(u.Id), 10)
    sess.Save(c.Request(), c.Response())

    cname := http.Cookie{Name: "username", Value: u.UserName, Expires: time.Now().Add(expire), Path: "/"}
    logo := http.Cookie{Name: "logo", Value: cfg.Logo, Expires: time.Now().Add(expire), Path: "/"}
    timezone := http.Cookie{Name: "timezone", Value: cfg.TimeZone.String(), Expires: time.Now().Add(expire), Path: "/"}
    ctx.SetCookie(&cname)
    ctx.SetCookie(&logo)
    ctx.SetCookie(&timezone)

    return ctx.SendResponse("ok")
}

func UserList(c echo.Context) error {
    ctx := c.(*mid.Context)
    us := model.GetUsers()
    // 隐藏密码
    for _, v := range us {
        v.Password = ""
    }
    return ctx.SendResponse(us)
}

func formatAllowGmCmd(allowGmCmd string) string {
    if len(allowGmCmd)>0{
        tempCmdArr := make([]string,0)
        arr:= strings.Split(allowGmCmd,"\n")
        reg := regexp.MustCompile(`^[a-zA-z_]+`)
        for _,v := range arr {
            tempCmd := reg.FindString(v)
            if len(tempCmd)>0 {
                tempCmdArr = append(tempCmdArr,tempCmd)
            }
        }
        allowGmCmd = strings.Join(tempCmdArr,"\n")
    }
    return allowGmCmd
}

func UserAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    role, _ := strconv.ParseUint(ctx.FormValue("role"), 10, 32)
    allowGmCmd := ctx.FormValue("allowGmCmd")
    allowGmCmd = formatAllowGmCmd(allowGmCmd)
    u := model.AddUser(username, password, uint32(role),allowGmCmd)
    if u == nil {
        return ctx.SendError(-1, "用户名已存在")
    }

    return ctx.SendResponse("添加用户成功")
}

func UserDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "用户不存在")
    }

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        u := model.GetUser(uint32(id)) 
        if u == nil {
            return ctx.SendError(-1, "用户不存在")
        }
        // 不能删除超级管理员
        if u.GetId() == 1001 {
            return ctx.SendError(-1, "不能删除超级管理员")
        }
        if model.DelUser(u) == false {
            return ctx.SendError(-1, "删除用户失败")
        }
    }
    return ctx.SendResponse("删除用户成功")
}

func UserUpdateRole(c echo.Context) error {
    ctx := c.(*mid.Context)
    userid, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 32)
    role, _ := strconv.ParseUint(ctx.FormValue("role"), 10, 32)

    u := model.GetUser(uint32(userid))
    if u == nil {
        return ctx.SendError(-1, "用户不存在")
    }

    r := model.GetRole(uint32(role))
    if r == nil {
        return ctx.SendError(-1, "角色不存在")
    }

    u.Role = uint32(role)
    if !model.UpdateUser(u) {
        return ctx.SendError(-1, "更新用户角色失败")
    }
    return ctx.SendResponse("更新用户角色成功")
}

func UserUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    userid, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 32)
    role, _ := strconv.ParseUint(ctx.FormValue("role"), 10, 32)
    password := ctx.FormValue("password")
    allowGmCmd := ctx.FormValue("allowGmCmd")
    allowGmCmd = formatAllowGmCmd(allowGmCmd)

    u := model.GetUser(uint32(userid))
    if u == nil {
        return ctx.SendError(-1, "用户不存在")
    }

    r := model.GetRole(uint32(role))
    if r == nil {
        return ctx.SendError(-1, "角色不存在")
    }

    if password != "" {
        if !u.EncodePwd(password) {
            return ctx.SendError(-1, "更新密码失败")
        }
    }

    u.Role = uint32(role)
    u.AllowGmCmd = allowGmCmd
    if !model.UpdateUser(u) {
        return ctx.SendError(-1, "更新用户失败")
    }
    return ctx.SendResponse("更新用户成功")
}

func UserPwd(c echo.Context) error {
    ctx := c.(*mid.Context)
    username := ctx.FormValue("username")
    oldpassword := ctx.FormValue("oldpassword")
    password := ctx.FormValue("password")

    if username == "" || oldpassword == "" || password == "" {
        return ctx.SendError(-1, "参数非法")
    }

    u := model.GetUserByUserName(username)
    if u == nil {
        return ctx.SendError(-1, "用户不存在")
    }

    if !u.ComparePwd(oldpassword) {
        return ctx.SendError(-1, "原始密码不对")
    }

    if !u.EncodePwd(password) || !model.UpdateUser(u) {
        return ctx.SendError(-1, "修改密码失败")
    }

    return ctx.SendResponse("修改密码成功")
}
