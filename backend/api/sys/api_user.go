package sys

import (
    "time"
    "net/http"
    "strconv"
    "strings"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func UserLogin(c echo.Context) error {
    ctx := c.(*mid.Context)
    if ctx.GetUserId() != 0 {
        return ctx.SendError(-2, "已经登陆不用重新登录")
    }

    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    u := model.GetUserByUserName(username)
    if u == nil || !u.ComparePwd(password) {
        return ctx.SendError(-1, "用户名或密码输入错误")
    }

    // Create token
    token := jwt.New(jwt.SigningMethodHS256)

    expire := time.Now().Add(time.Hour * 24)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = strconv.FormatUint(uint64(u.Id), 10)
    claims["exp"] = strconv.FormatInt(expire.Unix(), 10)

    // Generate encoded token and send it as response.
    t, err := token.SignedString([]byte(mid.UserKey))
    if err != nil {
        return ctx.SendError(-1, err.Error())
    }

    cjwt := http.Cookie{Name: "textoken", Value: t, Expires: expire, Path: "/"}
    cname := http.Cookie{Name: "username", Value: username, Expires: expire, Path: "/"}
    ctx.SetCookie(&cjwt)
    ctx.SetCookie(&cname)

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

func UserAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    role, _ := strconv.ParseUint(ctx.FormValue("role"), 10, 32)

    u := model.AddUser(username, password, uint32(role))
    if u == nil {
        return ctx.SendError(-1, "invalid param")
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

    u := model.GetUser(uint32(userid))
    if u == nil {
        return ctx.SendError(-1, "用户不存在")
    }

    r := model.GetRole(uint32(role))
    if r == nil {
        return ctx.SendError(-1, "角色不存在")
    }

    if password == "" {
        return ctx.SendError(-1, "密码不能为空");
    }
    if !u.EncodePwd(password) {
        return ctx.SendError(-1, "更新密码失败")
    }

    u.Role = uint32(role)
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
