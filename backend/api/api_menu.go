package api

import (
    "errors"
    "strconv"
    "strings"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func MenuList(c echo.Context) error {
    ctx := c.(*mid.Context)
    ms := model.GetMenus()
    return ctx.SendResponse(ms)
}

func MenuUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    id, _ := strconv.ParseUint(ctx.FormValue("id"), 10 ,32)
    role := ctx.FormValue("role")

    m := model.GetMenu(uint32(id))
    if m == nil {
        return errors.New("菜单不存在")
    }

    ids := make([]uint32,0)
    if role != "" {
        roles := strings.Split(role, ",")
        for _, rid := range roles {
            id, _ := strconv.ParseUint(rid, 10, 32)
            if model.GetRole(uint32(id)) == nil {
                return errors.New("角色不存在,id:"+rid)
            }
            ids = append(ids, uint32(id))
        }
    }
    m.Role = ids[:]
    if !model.UpdateMenu(m) {
        return errors.New("更新菜单失败")
    }

    return ctx.SendResponse("");
}
