package stat

import (
    "fmt"
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

type _incomemoneylog struct {
    Statymd string `json:"statymd"`
    zoneid uint32
    Zonename string `json:"zone_name"`
    Zoneopenday uint32 `json:"zone_openday"`
    Newadd uint32 `json:"newadd"`
    Money float32 `json:"money"`
    Money3 float32 `json:"money3"`
    Money7 float32 `json:"money7"`
    Money14 float32 `json:"money14"`
    Money30 float32 `json:"money30"`
    Days []float32 `json:"days"`
}

type statval2 struct {
    days []float32
}

func IncomeTrack(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    now,_ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
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

    sql := "SELECT statymd,zoneid,sum(newadd) as newadd"
    for i := 0;i < 90; i++ {
        sql += fmt.Sprintf(",sum(day_%d) as sum_day_%d",i+1,i+1)
    }
    sql += " FROM t_income_money"
    sql += " WHERE statymd between '%s' AND '%s'" 
    if zoneid != "" {
        sql += " AND zoneid in ("+zoneid+")"
    }
    sql += " GROUP BY statymd,zoneid ORDER BY statymd desc,zoneid desc"
    var total int
    err = tx.QueryRow("SELECT count(*) as total FROM ("+fmt.Sprintf(sql, startTime, endTime)+") a").Scan(&total)
    if err != nil {
        return err
    }

    logs := make([]_incomemoneylog, 0)
    if total == 0 {
        return ctx.SendArray(logs, total)
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    rows, err := tx.Query(fmt.Sprintf(sql+" LIMIT "+limitstart+","+limitrow, startTime, endTime))
    if err != nil {
        return err
    }
    defer rows.Close()

    mdays := make(map[statkey]*statval2,0)
    money := make([]float32,90)

    mzone := gm.ZoneMap()
    for rows.Next() {
        var r _incomemoneylog
        if err := rows.Scan(&r.Statymd, &r.zoneid, &r.Newadd, &money[0], &money[1], &money[2], &money[3], &money[4], &money[5], &money[6], &money[7], &money[8], &money[9], &money[10], &money[11], &money[12], &money[13], &money[14], &money[15], &money[16], &money[17], &money[18], &money[19], &money[20], &money[21], &money[22], &money[23], &money[24], &money[25], &money[26], &money[27], &money[28], &money[29], &money[30], &money[31], &money[32], &money[33], &money[34], &money[35], &money[36], &money[37], &money[38], &money[39], &money[40], &money[41], &money[42], &money[43], &money[44], &money[45], &money[46], &money[47], &money[48], &money[49], &money[50], &money[51], &money[52], &money[53], &money[54], &money[55], &money[56], &money[57], &money[58], &money[59], &money[60], &money[61], &money[62], &money[63], &money[64], &money[65], &money[66], &money[67], &money[68], &money[69], &money[70], &money[71], &money[72], &money[73], &money[74], &money[75], &money[76], &money[77], &money[78], &money[79], &money[80], &money[81], &money[82], &money[83], &money[84], &money[85], &money[86], &money[87], &money[88], &money[89]); err != nil {
            return err
        }

        v := statval2{}
        v.days = make([]float32,90)
        copy(v.days, money)
        k := statkey{statymd:r.Statymd, zoneid:r.zoneid}
        mdays[k] = &v
        r.Days = append(r.Days, money[0])
        r.Money += money[0]
        r.Money3 += money[0]
        r.Money7 += money[0]
        r.Money14 += money[0]
        r.Money30 += money[0]

        if v,ok := mzone[r.zoneid]; ok {
            r.Zonename = v.SZoneName+"("+strconv.Itoa(int(v.IZoneId))+")"
            r.Zoneopenday = uint32(now.Sub(time.Unix(int64(v.IPublishTime),0)).Hours()/24)
        }

        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    t2,_ := time.Parse("2006-01-02", logs[0].Statymd)
    startTime2 := t2.Add(24*time.Hour).Format("2006-01-02")
    endTime2 := t2.Add(90*24*time.Hour).Format("2006-01-02")
    rows2, err := tx.Query(fmt.Sprintf(sql, startTime2, endTime2))
    if err != nil {
        return err
    }
    defer rows2.Close()
    for rows2.Next() {
        var r _incomemoneylog
        if err := rows2.Scan(&r.Statymd, &r.zoneid, &r.Newadd, &money[0], &money[1], &money[2], &money[3], &money[4], &money[5], &money[6], &money[7], &money[8], &money[9], &money[10], &money[11], &money[12], &money[13], &money[14], &money[15], &money[16], &money[17], &money[18], &money[19], &money[20], &money[21], &money[22], &money[23], &money[24], &money[25], &money[26], &money[27], &money[28], &money[29], &money[30], &money[31], &money[32], &money[33], &money[34], &money[35], &money[36], &money[37], &money[38], &money[39], &money[40], &money[41], &money[42], &money[43], &money[44], &money[45], &money[46], &money[47], &money[48], &money[49], &money[50], &money[51], &money[52], &money[53], &money[54], &money[55], &money[56], &money[57], &money[58], &money[59], &money[60], &money[61], &money[62], &money[63], &money[64], &money[65], &money[66], &money[67], &money[68], &money[69], &money[70], &money[71], &money[72], &money[73], &money[74], &money[75], &money[76], &money[77], &money[78], &money[79], &money[80], &money[81], &money[82], &money[83], &money[84], &money[85], &money[86], &money[87], &money[88], &money[89]); err != nil {
            return err
        }

        v := statval2{}
        v.days = make([]float32,90)
        copy(v.days, money)
        k := statkey{statymd:r.Statymd, zoneid:r.zoneid}
        mdays[k] = &v
    }
    if err := rows2.Err(); err != nil {
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }

    for i,_ := range logs {
        rlog := &logs[i]
        t,_ := time.Parse("2006-01-02", rlog.Statymd)
        for j := 1; j < 90; j++ {
            t2 := t.Add(time.Duration(j*24)*time.Hour)
            if !t2.Before(now) {
                break
            }
            day := t2.Format("2006-01-02")
            money := float32(0)

            if v, ok := mdays[statkey{statymd:day,zoneid:rlog.zoneid}]; ok {
                money = v.days[j]
            }

            rlog.Days = append(rlog.Days, money)
            rlog.Money += money

            if j < 30 {
                if j < 3 {
                    rlog.Money3 += money
                    rlog.Money7 += money
                    rlog.Money14 += money
                } else if j < 7 {
                    rlog.Money7 += money
                    rlog.Money14 += money
                } else if j < 14 {
                    rlog.Money14 += money
                }
                rlog.Money30 += money
            }
        }
    }

    return ctx.SendArray(logs, total)
}
