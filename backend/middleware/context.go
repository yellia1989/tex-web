package middleware

import (
    "time"
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

const (
    UserKey = "github.com/yellia1989/tex-web/backend/context"
)

type Context struct {
    echo.Context
}

func NewContext() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := &Context{c}
            return next(ctx)
        }
    }
}

func RequireLogin() echo.MiddlewareFunc {
    return middleware.JWTWithConfig(middleware.JWTConfig{
        Skipper: func(c echo.Context) bool {
            return pathIgnore(c)
        },
        SigningKey: []byte(UserKey),
        TokenLookup: "cookie:textoken",
    })
}

func GetUserId(c echo.Context) uint32 {
    v := c.Get("user")
    if v == nil {
        return 0
    }
    user := v.(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    exp, err := strconv.ParseInt(claims["exp"].(string), 10, 64)
    if err != nil || time.Now().Unix() > exp {
        return 0
    }
    id, err := strconv.ParseUint(claims["id"].(string), 10, 32)
    if err != nil {
        return 0
    }
    return uint32(id)
}

func (ctx *Context) GetUserId() uint32 {
    return GetUserId(ctx)
}

func (ctx *Context) GetUser() *model.User {
    uid := ctx.GetUserId()
    if uid == 0 {
        return nil
    }
    return model.GetUser(uid)
}

func (ctx *Context) SendResponse(data interface{}) error {
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "code": 0,
        "data": data,
    })
}

func (ctx *Context) SendError(code int, msg interface{}) error {
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "code": code,
        "msg": msg,
    })
}
