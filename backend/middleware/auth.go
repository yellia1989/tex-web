package middleware

import (
    "fmt"
    "sort"
    "strings"
    "crypto/md5"
    "github.com/labstack/echo/v4"
    "github.com/yellia1989/tex-web/backend/model"
    "github.com/yellia1989/tex-go/tools/log"
)

func signstr(path string, params map[string]string) string {
    keys := make([]string,0)
    for k,_ := range params {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    signstr := ""
    for _,k := range keys {
        signstr += k + "&" + params[k]
    }
    return path + signstr + API_KEY
}

func Sign(path string, params map[string]string) string {
    str := signstr(path, params)
    sign := md5.Sum([]byte(str))

    return fmt.Sprintf("%x", sign)
}

func CheckSign(path string, params map[string]string, sign string) bool {
    mysign := Sign(path, params)
    if sign != mysign {
        return false
    }
    return true
}

func RequireAuth() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            ctx := c.(*Context)
            method := ctx.Request().Method
            path := ctx.Request().URL.Path
            if path == "/api/login" {
                return next(c)
            }

            if strings.HasPrefix(path, "/api/public/gm") {
                // 验证签名
                params := ctx.QueryParams()
                sign := ""
                p := make(map[string]string)
                for k,_ := range params {
                    if k == "sign" {
                        sign = ctx.QueryParam(k)
                    } else {
                        p[k] = ctx.QueryParam(k)
                    }
                }
                if !CheckSign(path, p, sign) {
                    log.Debugf("path: %s, params: %v, sign: %s", path, params, sign)
                    return fmt.Errorf("签名验证失败")
                }
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
