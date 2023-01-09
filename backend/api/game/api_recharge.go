package game

import (
    "fmt"
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/log"
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
    orderno := ctx.QueryParam("orderno")

    if zoneid == "" || startTime == "" || endTime == "" || (roleid == "" && orderno == ""){
        return ctx.SendError(-1, "参数非法")
    }

    db, err := zoneLogDb(zoneid)

    if err != nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

    flowid := ""
    if orderno != "" {
        sqltmp := "SELECT flowid FROM iap_trace_buy WHERE extern_order_id='"+orderno+"' or flowid='" +orderno+ "'"
        if db.QueryRow(sqltmp).Scan(&flowid) != nil {
            return ctx.SendError(-1, "没有对应订单")
        }
    }

    sqlcount := "SELECT count(DISTINCT flowid) as total FROM iap_trace_buy"
    sqlcount += " WHERE time between '"+startTime+"' AND '"+endTime+"'" 
    if roleid != "" {
        sqlcount += " AND roleid=" + roleid 
    }
    if flowid != "" {
        sqlcount += " AND flowid='" + flowid + "'"
    }

    var total int
    db.QueryRow(sqlcount).Scan(&total)

    logs := make([]*flowtrace,0)
    if total == 0 {
        return ctx.SendArray(logs, total)
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT flowid,product_id,min(time) as start_time FROM iap_trace_buy"
    sql += " WHERE time between '"+startTime+"' AND '"+endTime+"'" 
    if roleid != "" {
        sql += " AND roleid=" + roleid 
    }
    if flowid != "" {
        sql += " AND flowid='" + flowid + "'"
    }

    sql += " GROUP by flowid"
    sql += " ORDER by min(time) desc"
    sql += " LIMIT "+limitstart+","+limitrow

    log.Infof("sql: %s", sql)

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

    sql = "SELECT flowid,time,status,extern_order_id FROM iap_trace_buy WHERE flowid in("+string(flowids2)+") ORDER BY time"

    rows2, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows2.Close()
    for rows2.Next() {
        var f uint64
        var r steptrace
        var orderno string
        if err := rows2.Scan(&f, &r.Time, &r.Status, &orderno); err != nil {
            return err
        }
        l := flows[f]
        l.Steps = append(l.Steps, r)
        if orderno != "" {
            l.ExternOrderId = orderno
        }
    }
    if err := rows2.Err(); err != nil {
        return err
    }

    for _, v := range flowids {
        logs = append(logs, flows[v])
    }

    return ctx.SendArray(logs, total)
}
