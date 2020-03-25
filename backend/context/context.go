package context

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
                        return true
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

func HTTPErrorHandler(err error, c echo.Context) {
    he, ok := err.(*echo.HTTPError)
    if ok {
        if he.Internal != nil {
            if herr, ok := he.Internal.(*echo.HTTPError); ok {
                he = herr
            }
        }
    } else {
        he = &echo.HTTPError{
            Code:    http.StatusInternalServerError,
            Message: http.StatusText(http.StatusInternalServerError),
        }
    }

    code := he.Code
    message := he.Message

    c.Logger().Error("error:"+he.Error())

    // Send response
    if !c.Response().Committed {
        if c.Request().Method == http.MethodHead { // Issue #608
            err = c.NoContent(he.Code)
        } else {
            // 格式化错误消息
            path := c.Request().URL.Path
            if strings.HasPrefix(path, "/api") {
                err = c.JSON(http.StatusOK, map[string]interface{}{
                    "code": code,
                    "msg": message,
                })
            } else {
                // 没有登陆的话重定位到登陆
                c.Logger().Error(c.Request())
                if getUserId(c) == 0 {
                    err = c.Redirect(http.StatusMovedPermanently, "/login.html")
                } else {
                    redirect := "/500.html"
                    switch code {
                    case http.StatusForbidden:redirect = "/403.html"
                    case http.StatusNotFound:redirect = "/404.html"
                    }
                    err = c.Redirect(http.StatusMovedPermanently, redirect)
                }
            }
        }
        if err != nil {
            c.Logger().Error(err)
        }
    }
}

func getUserId(c echo.Context) int {
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
    id, err := strconv.Atoi(claims["id"].(string))
    if err != nil {
        return 0
    }
    return id
}

func (ctx *Context) GetUserId() int {
    return getUserId(ctx)
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
        return ctx.Redirect(http.StatusMovedPermanently, "/index.html")
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
