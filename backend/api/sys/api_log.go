package sys

import (
    "time"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type _syslog struct {
    Time string `json:"time"`
    Action string `json:"action"`
    Username string `json:"username"`
    Desc string `json:"desc"`
}

func LogList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    username := ctx.QueryParam("username")
    action := ctx.QueryParam("action")

    if startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    sqlcount := "SELECT count(*) as total FROM log"
    sqlcount += " WHERE time between '"+startTime+"' AND '"+endTime+"'" 
    if username != "" {
        sqlcount += " AND username='"+username+"'"
    }
    if action != "" {
        sqlcount += " AND action='"+action+"'"
    }
    var total int
    err := db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT time,username,action,`desc` FROM log"
    sql += " WHERE time between '"+startTime+"' AND '"+endTime+"'" 
    if username != "" {
        sql += " AND username='"+username+"'"
    }
    if action != "" {
        sql += " AND action='"+action+"'"
    }
    sql += " ORDER BY time desc"
    sql += " LIMIT "+limitstart+","+limitrow

    c.Logger().Debug(sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]_syslog, 0)
    for rows.Next() {
        var r _syslog
        if err := rows.Scan(&r.Time, &r.Username, &r.Action, &r.Desc); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}

func LogAdd(userName string, action string, desc string) {
    var err error
    db := cfg.StatDb
    if db == nil {
        return
    }

    _, err = db.Exec("insert into log(time, username, action, `desc`) values(?,?,?,?)", time.Now().Format("2006-01-02 15:04:05"), userName, action, desc)
    if err != nil {
        return
    }
}
