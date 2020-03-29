package middleware

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/casbin/casbin"
    "github.com/yellia1989/tex-web/backend/model"
)

var ce *casbin.Enforcer

func init() {
    var err error
    ce, err = casbin.NewEnforcer("data/auth_model.conf", "data/auth_policy.csv")
    _ = err
}

func RequireAuth() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := c.(*Context)
            // 验证权限
            userid := ctx.GetUserId()
            if userid == 0 {
                // 只处理登陆的情况
                return next(c)
            }
            return checkAuth(userid, ctx)
        }
    }
}

func checkAuth(userid uint32, ctx *Context) error {
    if ce == nil {
        return nil
    }

    user := model.GetUser(userid)
    if user == nil {
        return &echo.HTTPError{
            Code:    http.StatusInternalServerError,
            Message: "invalid userid",
        }
    }

    method := ctx.Request().Method
    path := ctx.Request().URL.Path
    pass, err := ce.Enforce(user.Name, path, method)
    if err != nil {
        return err
    }
    if !pass {
        return echo.ErrForbidden
    }
    return nil
}
