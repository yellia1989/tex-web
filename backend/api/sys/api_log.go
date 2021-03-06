package sys

import (
    "time"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
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

    sqlcount := "SELECT count(*) as total FROM t_log"
    sqlcount += " WHERE time between '"+startTime+"' AND '"+endTime+"'" 
    if username != "" {
        sqlcount += " AND username='"+username+"'"
    }
    if action != "" {
        sqlcount += " AND action='"+action+"'"
    }
    var total int
    err = tx.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql := "SELECT time,username,action,`desc` FROM t_log"
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

    rows, err := tx.Query(sql)
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

    if err := tx.Commit(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}

func LogAdd(c echo.Context, action string, desc string) {
    ctx := c.(*mid.Context)

    var err error
    defer func() {
        if err != nil {
            c.Logger().Errorf("add sys log failed, %v", err)
        }
    }()

    db := common.GetStatDb()
    if db == nil {
        ctx.Logger().Error("connect stat db failed")
        return
    }

    tx, err := db.Begin()
    if err != nil {
        return
    }
    defer tx.Rollback()

    _, err = tx.Exec("USE db_stat")
    if err != nil {
        return
    }

    u := ctx.GetUser()
    if u == nil {
        return
    }

    _, err = tx.Exec("insert into t_log(time, username, action, `desc`) values(?,?,?,?)", time.Now().Format("2006-01-02 15:04:05"), u.UserName, action, desc)
    if err != nil {
        return
    }

    err = tx.Commit()
    if err != nil {
        return
    }
}
