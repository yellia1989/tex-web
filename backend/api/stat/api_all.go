package stat

import (
    "time"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/api/gm"
)

type _alllog struct {
    Statymd string `json:"statymd"`
    Zonename string `json:"zone_name"`
    Zoneopenday uint32 `json:"zone_openday"`
    Accountnum uint32 `json:"accountnum"`
    Rolenum uint32 `json:"rolenum"`
    RolenumIncrease float32 `json:"rolenum_increase"`
    Newadd uint32 `json:"newadd"`
    Active uint32 `json:"active"`
    LoginTimes uint32 `json:"login_times"`
    LoginTimesPer float32 `json:"login_times_per"`
    RechargeNum uint32 `json:"recharge_num"`
    RechargeMoney uint32 `json:"recharge_money"`
    RechargePer float32 `json:"recharge_per"`
    WeekActive uint32 `json:"week_active"`
    DoubleWeekActive uint32 `json:"double_week_active"`
    MonthActive uint32 `json:"month_active"`
    ActivePer float32 `json:"active_per"`
}

func AllList(c echo.Context) error {
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

    mzone := gm.ZoneMap()

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

    sql := "SELECT statymd,zoneid,sum(accountnum) as accountnum,sum(rolenum) as rolenum,sum(newadd) as newadd,sum(active) as active,sum(login_times) as login_times,sum(recharge_num) as recharge_num,sum(recharge_money) as recharge_money,sum(week_active) as week_active,sum(double_week_active) as double_week_active, sum(month_active) as month_active FROM t_all"
    sql += " WHERE statymd between '"+startTime+"' AND '"+endTime+"'" 
    if zoneid != "" {
        sql += " AND zoneid in ("+zoneid+")"
    }
    sql += " GROUP BY statymd, zoneid"
    sql += " ORDER BY statymd desc, zoneid desc"
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

    logs := make([]_alllog, 0)
    for rows.Next() {
        var r _alllog
        var zoneid uint32
        if err := rows.Scan(&r.Statymd, &zoneid, &r.Accountnum, &r.Rolenum, &r.Newadd, &r.Active, &r.LoginTimes, &r.RechargeNum, &r.RechargeMoney, &r.WeekActive, &r.DoubleWeekActive, &r.MonthActive); err != nil {
            return err
        }
        if v,ok := mzone[zoneid]; ok {
            r.Zonename = v.SZoneName
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
