package game

import (
    "fmt"
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/log"
)

type _mailSendLog struct {
    Id uint32 `json:"id"`
    MailId uint32 `json:"mailid"`
    Time string `json:"sendtime"`
    TemplateId uint32 `json:"templateid"`
    ItemId string `json:"itemid"`
    ItemNum string `json:"itemnum"`
}

type _mailRevLog struct {
    Id uint32 `json:"id"`
    MailId uint32 `json:"mailid"`
    Time string `json:"time"`
    ItemId string `json:"itemid"`
    ItemNum string `json:"itemnum"`   
}

type _mailDelLog struct {
    Id uint32 `json:"id"`
    MailId uint32 `json:"mailid"`
    Time string `json:"time"`
}

func MailSendLog(c echo.Context) error {
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

    if db == nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

    sqlcount := "SELECT count(*) as total FROM send_mail"
    sqlcount += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    var total int
    err = db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT _rid as id,mailid,sendtime,templateid,itemid,itemnum FROM send_mail"
    sql += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    sql += " ORDER BY time desc, _rid desc"
    sql += " LIMIT "+limitstart+","+limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]_mailSendLog, 0)
    for rows.Next() {
        var r _mailSendLog
        if err := rows.Scan(&r.Id, &r.MailId, &r.Time, &r.TemplateId, &r.ItemId, &r.ItemNum); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}

func MailRevLog(c echo.Context) error {
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

    if db == nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

    sqlcount := "SELECT count(*) as total FROM rcv_mail"
    sqlcount += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    var total int
    err = db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT _rid as id,mailid,time,itemid,itemnum FROM rcv_mail"
    sql += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    sql += " ORDER BY time desc, _rid desc"
    sql += " LIMIT "+limitstart+","+limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]_mailRevLog, 0)
    for rows.Next() {
        var r _mailRevLog
        if err := rows.Scan(&r.Id, &r.MailId, &r.Time, &r.ItemId, &r.ItemNum); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}

func MailDelLog(c echo.Context) error {
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

    if db == nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

    sqlcount := "SELECT count(*) as total FROM del_mail"
    sqlcount += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    var total int
    err = db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT _rid as id,mailid,time FROM del_mail"
    sql += " WHERE roleid="+roleid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    sql += " ORDER BY time desc, _rid desc"
    sql += " LIMIT "+limitstart+","+limitrow

    c.Logger().Info(sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]_mailDelLog, 0)
    for rows.Next() {
        var r _mailDelLog
        if err := rows.Scan(&r.Id, &r.MailId, &r.Time); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}
