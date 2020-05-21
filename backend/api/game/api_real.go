package game

import (
    "fmt"
    "time"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
)

func online(day string) ([]uint32,error) {
    db := common.GetLogDb()
    if db == nil {
        return nil,fmt.Errorf("连接数据库失败")
    }

    tx, err := db.Begin()
    if err != nil {
        return nil,err
    }
    defer tx.Rollback()

    _, err = tx.Exec("USE log_global")
    if err != nil {
        return nil,err
    }

    sql := "SELECT time, sum(num) as num FROM realtime_online WHERE left(time,10)="
    sql += "'"+day+"'"
    sql += " GROUP BY time ORDER BY time"

    rows, err := tx.Query(sql)
    if err != nil {
        return nil,err
    }
    defer rows.Close()

    points := make(map[string]uint32, 0)
    for rows.Next() {
        var time string
        var num uint32
        if err := rows.Scan(&time, &num); err != nil {
            return nil,err
        }
        points[time] = num
    }
    if err := rows.Err(); err != nil {
        return nil,err
    }

    if err := tx.Commit(); err != nil {
        return nil,err
    }

    daytime, err := time.Parse("2006-01-02 15:04:05", day + " 00:00:00")
    if err != nil {
        return nil,err
    }

    ret := make([]uint32,288)
    for i,_ := range ret {
        t := daytime.Add(time.Duration(i*300)*time.Second).Format("2006-01-02 15:04:05")
        if n,ok := points[t]; ok {
            ret[i] = n
        }
    }

    return ret,nil
}

func RealOnline(c echo.Context) error {
    ctx := c.(*mid.Context)

    now := time.Now()

    // 今日在线
    data := make(map[string][]uint32,0)
    today, err := online(now.Format("2006-01-02"))
    if err != nil {
        return err
    }
    data["today"] = today

    // 昨日在线
    yesterday, err := online(now.Add(-24*time.Hour).Format("2006-01-02"))
    if err != nil {
        return err
    }
    data["yesterday"] = yesterday

    // 上周同期
    lastweek, err := online(now.Add(-24*7*time.Hour).Format("2006-01-02"))
    if err != nil {
        return err
    }
    data["lastweek"] = lastweek

    // 上月同期
    lastmonth, err := online(now.Add(-24*30*time.Hour).Format("2006-01-02"))
    if err != nil {
        return err
    }
    data["lastmonth"] = lastmonth
    
    return ctx.SendResponse(data)
}
