package middleware

import (
    "strings"
    "github.com/labstack/echo"
)

func pathIgnore(c echo.Context) bool {
    path := c.Request().URL.Path
    suffix := strings.Split(path, ".")
    switch suffix[len(suffix)-1] {
    case "html":
        // 特定页面不需要验证
        switch path {
        case "/login.html","/403.html","/404.html","/500.html":
            return true
        default:
            return false
        }
    case "css","js","jpg","png","gif","ico","woff2","woff","ttf":
        // 资源不需要验证
        return true
    default:
        if strings.HasPrefix(path, "/api") {
            // api需要特殊处理
            switch path {
            case "/api/login":
                // 上传了cookie的话需要处理
                if _, err := c.Cookie("textoken"); err != nil {
                    return true
                }
                return false
            default:
                return false
            }
        }
        return false
    }
}
