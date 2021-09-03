package sys

import (
    "errors"
    "strconv"
    "strings"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/util"
)

func MenuList(c echo.Context) error {
    ctx := c.(*mid.Context)
    ms := model.GetMenus()

    // 权限控制
    user := ctx.GetUser()
    ms = filterMenu(ms, user.Role, user.IsAdmin())

    return ctx.SendResponse(ms)
}

func filterMenu(ms []*model.Menu, role uint32, bAdmin bool) []*model.Menu {
    if len(ms) == 0 {
        return nil
    }
    ret := make([]*model.Menu, 0)
    for _, m := range ms {
        if !bAdmin {
            if !util.Contain(m.Role, role) {
                continue
            }
        }

        // 导入数据菜单只有内网才看得见
        if m.Id == 216 && cfg.ServerID != "1" {
            continue
        }

        // 处理子节点
        m.Children = filterMenu(m.Children, role, bAdmin)
        ret = append(ret, m)
    }
    return ret
}

func setMenuRole(m *model.Menu, role []uint32) {
    m.Role = role[:]
    for _,c := range m.Children {
        setMenuRole(c, role)
    }
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

    setMenuRole(m, ids)

    if !model.UpdateMenu(m) {
        return errors.New("更新菜单失败")
    }

    return ctx.SendResponse("");
}
