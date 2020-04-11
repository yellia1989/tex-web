package api

import (
    "time"
    "strconv"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func UserLogin(c echo.Context) error {
    ctx := c.(*mid.Context)
    if ctx.GetUserId() != 0 {
        return ctx.SendError(-1, "已经登陆不用重新登录")
    }

    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    u := model.GetUserByUserName(username)
    if u == nil || u.Password != password {
        return ctx.SendError(-1, "用户名或密码输入错误")
    }

    // Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = strconv.FormatUint(uint64(u.Id), 10)
    claims["exp"] = strconv.FormatInt(time.Now().Add(time.Hour * 24).Unix(), 10)

    // Generate encoded token and send it as response.
    t, err := token.SignedString([]byte(mid.UserKey))
    if err != nil {
        return ctx.SendError(-1, err.Error())
    }

    return ctx.SendResponse(map[string]interface{}{"token":t, "day":1, "username": u.UserName})
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
