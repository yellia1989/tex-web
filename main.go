package main

import (
    "os"
    "fmt"
    "strings"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api"
    "github.com/yellia1989/tex-web/backend/cron"
    "github.com/yellia1989/tex-web/backend/common"
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
                if mid.GetUserId(c) == 0 {
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
    if err := common.ParseCfg("conf.cfg"); err != nil {
        fmt.Printf("%s", err)
        os.Exit(-1)
    }
    
    debug := common.Cfg.GetBool("debug", false)
    framework_debug := common.Cfg.GetBool("framework-debug", false)

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

    // Start Cron
    cron.Start()

    // Start server
    e.Logger.Fatal(e.Start(common.Cfg.GetCfg("listen", "")))

    // Stop Cron
    cron.Stop()

    log.FlushLogger()
}
