package main

import (
    "github.com/yellia1989/tex-web/backend/model"
    "os"
    "fmt"
    "strings"
    "net/http"
    _ "net/http/pprof"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api"
    "github.com/yellia1989/tex-web/backend/api/stat"
    "github.com/yellia1989/tex-web/backend/cron"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

func httpErrorHandler(err error, c echo.Context) {
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
            Message: err.Error(),
        }
    }

    code := he.Code
    message := he.Message

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
                bReLogin := false
                userId := mid.GetUserId(c)
                if userId==0 {
                    bReLogin = true
                }else {
                    pUser := model.GetUser(userId)
                    if pUser == nil {
                        bReLogin = true
                    }else if pUser.NeedReLogin {
                        bReLogin = true
                    }
                }
                if bReLogin {
                    err = c.Redirect(http.StatusFound, "/login.html")
                } else {
                    redirect := "/500.html"
                    switch code {
                    case http.StatusForbidden:redirect = "/403.html"
                    case http.StatusNotFound:redirect = "/404.html"
                    }
                    err = c.Redirect(http.StatusFound, redirect)
                }
            }
        }
        if err != nil {
            c.Logger().Error(err)
        }
    }
}

func main() {
    if err := cfg.ParseCfg("conf.cfg"); err != nil {
        fmt.Printf("%s", err)
        os.Exit(-1)
    }

    debug := cfg.Debug
    framework_debug := cfg.FrameworkDebug

    // Echo instance
    e := echo.New()
    e.Debug = debug
    e.HTTPErrorHandler = httpErrorHandler
    e.Logger.SetHeader("${time_custom}|${short_file}:${line}|${level}|")

    // Middleware
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "${time_custom}|${remote_ip}|${method}|${path}|${status}|${latency_human}|${error}\n",
        CustomTimeFormat: "2006-01-02 15:04:05", 
    }))
    e.Use(middleware.Recover())

    e.Use(mid.NewContext(), mid.RequireLogin(), mid.RequireAuth())

    e.Static("/", "front/pages")
    e.Static("/lib", "front/lib")
    e.Static("/css", "front/css")
    e.Static("/js", "front/js")
    e.Static("/images", "front/images")

    api.RegisterHandler(e.Group("/api"))

    if debug {
        log.SetLevel(log.DEBUG)
    }

    if framework_debug {
        log.SetFrameworkLevel(log.DEBUG)
    }

    go func() {                                                                                                                                                        
        log.Debug(http.ListenAndServe(":16060", nil))
    }()

    stat.InitCondition()

    // Start Cron
    cron.Start()

    // Start server
    e.Logger.Fatal(e.Start(cfg.Listen))

    // Stop Cron
    cron.Stop()

    log.FlushLogger()
}
