package api

import (
    "strconv"
    "strings"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func RoleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    rs := model.GetRoles()
    return ctx.SendResponse(rs)
}

func RoleAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    name := ctx.FormValue("name")
    perm := ctx.FormValue("perms")

    if name == "" {
        return ctx.SendError(-1, "参数非法")
    }

    perms := strings.Split(perm, ",")
    if len(perms) == 0 {
        return ctx.SendError(-1, "请给角色选择至少一个权限")
    }

    ids := make([]uint32, 0)
    for _,p := range perms {
        id, _ := strconv.ParseUint(p, 10, 32)
        if model.GetPerm(uint32(id)) == nil {
            return ctx.SendError(-1, "权限不存在")
        }
        ids = append(ids, uint32(id))
    }

    if model.AddRole(name, ids) == nil {
        return ctx.SendError(-1, "添加角色失败");
    }

    return ctx.SendResponse("添加角色成功")
}

func RoleDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "参数非法")
    }

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        r := model.GetRole(uint32(id)) 
        if r == nil {
            return ctx.SendError(-1, "角色不存在")
        }
        if model.DelRole(r) == false {
            return ctx.SendError(-1, "删除角色失败")
        }
    }
    return ctx.SendResponse("删除角色成功")
}

func RoleUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    rid, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 32)
    name := ctx.FormValue("name")
    perm := ctx.FormValue("perms")

    if name == "" {
        return ctx.SendError(-1, "参数非法")
    }

    perms := strings.Split(perm, ",")
    if len(perms) == 0 {
        return ctx.SendError(-1, "请给角色选择至少一个权限")
    }

    ids := make([]uint32, 0)
    for _,p := range perms {
        id, _ := strconv.ParseUint(p, 10, 32)
        if model.GetPerm(uint32(id)) == nil {
            return ctx.SendError(-1, "权限不存在")
        }
        ids = append(ids, uint32(id))
    }

    r := model.GetRole(uint32(rid))
    if r == nil {
        return ctx.SendError(-1, "参数非法");
    }

    r.Name = name
    r.Perms = ids[:]
    if !model.UpdateRole(r) {
        return ctx.SendError(-1, "更新角色失败")
    }

    return ctx.SendResponse("更新角色成功")
}
