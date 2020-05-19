package game

import (
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
)

type _role struct {
    Uid uint64  `json:"id"`
    Name string `json:"name"`
    VipLevel uint32 `json:"vip_level"`
    LastLoginTime string `json:"last_login_time"`
    RegTime string `json:"reg_time"`
}

func RoleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    name := ctx.QueryParam("name")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    field := ctx.QueryParam("field")
    order := ctx.QueryParam("order")

    db := common.GetDb()
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败");
    }

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("USE db_zone_" + zoneid)
    if err != nil {
        return err
    }

    sql := "SELECT uid,name,vip_level,last_login_time,reg_time FROM t_role"
    if name != "" {
        sql += " WHERE name like '%" + name + "%'"
    }
    if field != "" {
        sql += " ORDER BY " + field + " " + order
    }
    rows, err := tx.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    roles := make([]_role, 0)
    for rows.Next() {
        var r _role
        if err := rows.Scan(&r.Uid, &r.Name, &r.VipLevel, &r.LastLoginTime, &r.RegTime); err != nil {
            return err
        }
        roles = append(roles, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }

    vPage := common.GetPage(roles, page, limit)
    return ctx.SendArray(vPage, len(roles))
}
