package middleware

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "github.com/gorilla/sessions"
    "github.com/labstack/echo-contrib/session"
    "github.com/yellia1989/tex-web/backend/model"
)

const (
    SessionKey = "github.com/yellia1989/tex-web/backend/context"
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
    cfg := session.Config{
        Skipper: func(c echo.Context) bool {
            return pathIgnore(c)
        },
        Store: sessions.NewCookieStore([]byte(SessionKey)),
    }

    return session.MiddlewareWithConfig(cfg)
}

func GetUserId(c echo.Context) uint32 {
    sess, err := session.Get("texweb_session", c)
    if err!= nil{
        return 0
    }
    userid, ok := sess.Values["userid"]
    if !ok {
        return 0
    }

    id, err := strconv.ParseUint(userid.(string), 10, 32)
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

func (ctx *Context) SendArray(data interface{}, total int) error {
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "code": 0,
        "data": data,
        "count": total,
    })
}

func (ctx *Context) SendError(code int, msg interface{}) error {
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "code": code,
        "msg": msg,
    })
}
