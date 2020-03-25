package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/yellia1989/tex-web/backend/context"
    _ "github.com/yellia1989/tex-web/backend/api"
)

func main() {
    // Echo instance
    e := echo.New()
    e.HTTPErrorHandler = context.HTTPErrorHandler

    // Middleware
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(context.NewContext(), context.RequireLogin())
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Static("/", "front/pages")
    e.Static("/lib", "front/lib")
    e.Static("/css", "front/css")
    e.Static("/js", "front/js")
    e.Static("/images", "front/images")

    api := e.Group("/api")
    api.POST("/login", context.Login)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
