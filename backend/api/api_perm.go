package api

import (
    "strings"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func PermList(c echo.Context) error {
    ctx := c.(*mid.Context)
    ps := model.GetPerms()
    return ctx.SendResponse(ps)
}

func PermAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    name := ctx.FormValue("name")
    path := ctx.FormValue("paths")

    if name == "" || path == "" {
        return ctx.SendError(-1, "参数非法")
    }

    paths := strings.Split(path, "\n")
    if model.AddPerm(name, paths) == nil {
        return ctx.SendError(-1, "参数非法")
    }

    return ctx.SendResponse("添加权限成功")
}

func PermDel(c echo.Context) error {
    return nil
}

func PermUpdate(c echo.Context) error {
    return nil
}
