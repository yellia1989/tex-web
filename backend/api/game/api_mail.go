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
    Title string `json:"title"`
    Content string `json:"content"`
    Items string `json:"items"`
}

type _mailOptLog struct {
    Id uint32 `json:"id"`
    MailId uint32 `json:"mailid"`
    Time string `json:"time"`
    Opt string `json:"opt"`
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

    sqlcount := "SELECT count(*) as total FROM recv_mail"
    sqlcount += " WHERE roleid="+roleid+" AND zoneid="+zoneid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    var total int
    err = db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT _rid as id,mail_uid,time,mail_title,mail_content,mail_items FROM recv_mail"
    sql += " WHERE roleid="+roleid+" AND zoneid="+zoneid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
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
        if err := rows.Scan(&r.Id, &r.MailId, &r.Time, &r.Title, &r.Content, &r.Items); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}

func MailOptLog(c echo.Context) error {
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

    sqlcount := "SELECT count(*) as total FROM opt_mail"
    sqlcount += " WHERE roleid="+roleid+" AND zoneid="+zoneid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    var total int
    err = db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT _rid as id,mail_uid,time,action FROM opt_mail"
    sql += " WHERE roleid="+roleid+" AND zoneid="+zoneid+" AND time between '"+startTime+"' AND '"+endTime+"'" 
    sql += " ORDER BY time desc, _rid desc"
    sql += " LIMIT "+limitstart+","+limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]_mailOptLog, 0)
    for rows.Next() {
        var r _mailOptLog
        if err := rows.Scan(&r.Id, &r.MailId, &r.Time, &r.Opt); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}
