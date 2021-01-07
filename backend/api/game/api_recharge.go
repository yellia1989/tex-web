package game

import (
    "fmt"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

type steptrace struct {
    Time string `json:"time"`
    Status string `json:"status"`
}

type flowtrace struct {
    FlowId uint64 `json:"flow_id"`
    ExternOrderId string `json:"extern_order_id"`
    ProductId uint32 `json:"product_id"`
    Steps []steptrace `json:"steps"`
}

func RechargeTrace(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    roleid := ctx.QueryParam("roleid")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    if zoneid == "" || roleid == "" || startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db, err := zoneLogDb(zoneid)

    if err != nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

    sqlcount := "SELECT count(DISTINCT flowid) as total FROM iap_trace_buy"
    sqlcount += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 

    c.Logger().Debug(sqlcount)

    var total int
    db.QueryRow(sqlcount).Scan(&total)

    logs := make([]*flowtrace,0)
    if total == 0 {
        return ctx.SendArray(logs, total)
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT flowid,product_id,min(time) as start_time FROM iap_trace_buy"
    sql += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    sql += " GROUP by flowid"
    sql += " ORDER by min(time) desc"
    sql += " LIMIT "+limitstart+","+limitrow

    c.Logger().Debug(sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    flowids := make([]uint64,0)
    flows := make(map[uint64]*flowtrace, 0)
    for rows.Next() {
        var r flowtrace
        var t string
        if err := rows.Scan(&r.FlowId, &r.ProductId, &t); err != nil {
            return err
        }
        _ = t
        flowids = append(flowids, r.FlowId)
        flows[r.FlowId] = &r
    }
    if err := rows.Err(); err != nil {
        return err
    }

    flowids2 := make([]byte,0)
    for i,v := range flowids {
        if i != 0 {
            flowids2 = append(flowids2, ',')
        }
        flowids2 = strconv.AppendUint(flowids2, v, 10)
    }

    sql = "SELECT flowid,time,status FROM iap_trace_buy WHERE flowid in("+string(flowids2)+") ORDER BY time"

    c.Logger().Debug(sql)

    rows2, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows2.Close()
    for rows2.Next() {
        var f uint64
        var r steptrace
        if err := rows2.Scan(&f, &r.Time, &r.Status); err != nil {
            return err
        }
        l := flows[f]
        l.Steps = append(l.Steps, r)
    }
    if err := rows2.Err(); err != nil {
        return err
    }

    sql = "SELECT flowid,extern_order_id FROM iap_recharge WHERE flowid in("+string(flowids2)+")"

    c.Logger().Debug(sql)

    rows3, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows3.Close()
    for rows3.Next() {
        var f uint64
        var orderid string
        if err := rows3.Scan(&f, &orderid); err != nil {
            return err
        }
        flows[f].ExternOrderId = orderid
    }
    if err := rows3.Err(); err != nil {
        return err
    }

    for _, v := range flowids {
        logs = append(logs, flows[v])
    }

    return ctx.SendArray(logs, total)
}
