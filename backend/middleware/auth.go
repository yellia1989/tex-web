package middleware

import (
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-web/backend/model"
)

func RequireAuth() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := c.(*Context)
            method := ctx.Request().Method
            path := ctx.Request().URL.Path
            if path == "/api/login" {
                return next(c)
            }

            if pathIgnore(c) {
                return next(c)
            }

            // 验证权限
            userid := ctx.GetUserId()
            user := model.GetUser(userid)
            if user == nil || user.IsNeedLogin()  {
                return &echo.HTTPError{
                    Code:    9999,
                    Message: "登陆已过期，请重新登录",
                }
            }

            err := user.CheckPermission(path, method)
            pass := "pass"
            if err != nil {
                pass = "failed"
            }
            ctx.Logger().Debug("username:"+user.UserName+",path:"+path+",method:"+method+",pass:"+pass)
            if err != nil {
                return err
            }

            return next(c)
        }
    }
}
