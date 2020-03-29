package middleware

import (
    "time"
    "strings"
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/middleware"
)

const (
    user_key = "github.com/yellia1989/tex-web/backend/context"
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
            path := c.Request().URL.Path
            suffix := strings.Split(path, ".")
            switch suffix[len(suffix)-1] {
            case "html":
                // 特定页面不需要验证
                switch path {
                case "/login.html","/403.html","/404.html","505.html":
                    return true
                default:
                    return false
                }
            case "css","js","jpg","png","gif","ico":
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
        },
        SigningKey: []byte(user_key),
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

func Login(c echo.Context) error {
    ctx := c.(*Context)
    if ctx.GetUserId() != 0 {
        return ctx.SendError(-1, "已经登陆不用重新登录")
    }

    //username := ctx.FormValue("username")
    //password := ctx.FormValue("password")

    // Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = "1001"
    claims["exp"] = strconv.FormatInt(time.Now().Add(time.Hour * 72).Unix(), 10)

    // Generate encoded token and send it as response.
    t, err := token.SignedString([]byte(user_key))
    if err != nil {
        return ctx.SendError(-1, err.Error())
    }

    return ctx.SendResponse(map[string]interface{}{"token":t, "day":1})
}
