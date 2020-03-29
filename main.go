package main

import (
    "net/http"
    "strings"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    _ "github.com/yellia1989/tex-web/backend/api"
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
                if mid.GetUserId(c) == 0 {
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

func main() {
    // Echo instance
    e := echo.New()
    e.HTTPErrorHandler = httpErrorHandler

    // Middleware
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(mid.NewContext(), mid.RequireLogin(), mid.RequireAuth())

    e.Static("/", "front/pages")
    e.Static("/lib", "front/lib")
    e.Static("/css", "front/css")
    e.Static("/js", "front/js")
    e.Static("/images", "front/images")

    api := e.Group("/api")
    api.POST("/login", mid.Login)           // 登陆
    /*
    api.GET("/user/list", api.UserList)     // 用户列表
    api.POST("/user/add", api.UserAdd)      // 用户增加
    api.POST("/user/edit", api.UserEdit)    // 用户编辑
    api.POST("/user/del", api.UserDel)      // 用户删除
    api.GET("/role/list", api.RoleList)     // 角色列表
    api.POST("/role/add", api.RoleAdd)      // 角色增加
    api.POST("/role/del", api.RoleDel)      // 角色删除
    api.POST("/role/edit", api.RoleEdit)    // 角色编辑
    api.GET("/permission/list", api.PermissionList)     // 权限列表
    api.POST("/permission/add", api.PermissionAdd)      // 权限增加
    api.POST("/permission/del", api.PermissionDel)      // 权限删除
    api.POST("/permission/edit", api.PermissionEdit)    // 权限编辑
    */

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
