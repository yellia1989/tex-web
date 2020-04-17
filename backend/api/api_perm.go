package api

import (
    "strings"
    "strconv"
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
    ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "权限不存在")
    }

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        p := model.GetPerm(uint32(id)) 
        if p == nil {
            return ctx.SendError(-1, "权限不存在")
        }
        if model.DelPerm(p) == false {
            return ctx.SendError(-1, "删除权限失败")
        }
    }
    return ctx.SendResponse("删除权限成功")
}

func PermUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    id, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 32)
    name := ctx.FormValue("name")
    path := ctx.FormValue("paths")

    if name == "" || path == "" {
        return ctx.SendError(-1, "参数非法")
    }

    paths := strings.Split(path, "\n")
    p := model.GetPerm(uint32(id))
    if p == nil {
        return ctx.SendError(-1, "权限不存在")
    }

    p.Name = name
    p.Paths = paths[:]
    if !model.UpdatePerm(p) {
        return ctx.SendError(-1, "更新权限失败")
    }

    return ctx.SendResponse("更新权限成功")
}
