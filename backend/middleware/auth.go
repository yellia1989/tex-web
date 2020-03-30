package middleware

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-web/backend/model"
)

func RequireAuth() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            if pathIgnore(c) {
                return next(c)
            }

            // 验证权限
            ctx := c.(*Context)
            userid := ctx.GetUserId()
            user := model.GetUser(userid)
            if user == nil {
                return &echo.HTTPError{
                    Code:    http.StatusForbidden,
                    Message: "invalid userid",
                }
            }
            method := ctx.Request().Method
            path := ctx.Request().URL.Path
            // 对已经登陆还调用/api/login做特殊处理
            if path == "/api/login" {
                return next(c)
            }

            err := user.CheckPermission(path, method)
            pass := "pass"
            if err != nil {
                pass = "failed"
            }
            ctx.Logger().Error("username:"+user.UserName+",path:"+path+",method:"+method+",pass:"+pass)
            if err != nil {
                return err
            }

            return next(c)
        }
    }
}
