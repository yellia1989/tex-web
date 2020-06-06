package stat

import (
    "time"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/api/gm"
)

type _incomelog struct {
    Statymd string `json:"statymd"`
    zoneid uint32
    Zonename string `json:"zone_name"`
    Zoneopenday uint32 `json:"zone_openday"`
    Active uint32 `json:"active"`
    RechargeRolenum float32 `json:"recharge_rolenum"`
    RechargeMoney float32 `json:"recharge_money"`
    RechargeNewRolenum float32 `json:"recharge_newrolenum"`
}

func IncomeList(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    now := time.Now()
    if startTime == "" {
        startTime = now.Add(-7*24*time.Hour).Format("2006-01-02")
    }
    if endTime == "" {
        endTime = now.Format("2006-01-02")
    }

    db := common.GetStatDb()
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("USE db_stat")
    if err != nil {
        return err
    }

    sql := "SELECT statymd,zoneid,sum(active) as active,sum(recharge_num) as recharge_num,sum(recharge_money) as recharge_money,sum(recharge_new) as recharge_new FROM t_income_all"
    sql += " WHERE statymd between '"+startTime+"' AND '"+endTime+"'" 
    if zoneid != "" {
        sql += " AND zoneid in ("+zoneid+")"
    }
    sql += " GROUP BY statymd,zoneid ORDER BY statymd desc,zoneid desc"
    var total int
    err = tx.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

    c.Logger().Error(sql)

    rows, err := tx.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    mzone := gm.ZoneMap()
    logs := make([]_incomelog, 0)
    for rows.Next() {
        var r _incomelog
        if err := rows.Scan(&r.Statymd, &r.zoneid, &r.Active, &r.RechargeRolenum, &r.RechargeMoney, &r.RechargeNewRolenum); err != nil {
            return err
        }
        if v,ok := mzone[r.zoneid]; ok {
            r.Zonename = v.SZoneName+"("+strconv.Itoa(int(v.IZoneId))+")"
            r.Zoneopenday = uint32(now.Sub(time.Unix(int64(v.IPublishTime),0)).Hours()/24)
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}
