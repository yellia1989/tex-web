package stat

import (
    "time"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
)

// 获取总注册账号数量
func getAccountUtilNow(cond *condition) (uint32,error) {
    now := getDateByTime(time.Now())
    return getAccountUtilDate(now.ID, cond)
}

// 获取到某一天为止的总注册账号
func getAccountUtilDate(dateEnd uint32, cond *condition) (uint32, error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    t1 := time.Now()
    sql := "SELECT count(*) FROM account WHERE date_fk <= ?"
    if cond != nil {
        sql += " AND " + cond.String()
    }

    var total uint32
    err := db.QueryRow(sql, dateEnd).Scan(&total)
    if err != nil {
        return 0,err
    }

    t2 := time.Now()

    log.Debugf("sql account total by date query: %v", t2.Sub(t1))

    return total, nil
}

// 获取特定日期的注册账号
func getAccountByDate(dateBegin uint32, dateEnd uint32, cond *condition) (map[uint32]uint32,error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT date_fk,count(*) FROM account WHERE"
    if cond != nil {
        sql += cond.String() + " AND "
    }
    sql += " date_fk BETWEEN ? AND ? GROUP BY date_fk ORDER BY NULL"

    t1 := time.Now()

    rows, err := db.Query(sql, dateBegin, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()

    accounts := make(map[uint32]uint32)
    var dateFk uint32
    var accountNum uint32
    for rows.Next() {
        if err := rows.Scan(&dateFk, &accountNum); err != nil {
            return nil, err
        }
        accounts[dateFk] = accountNum
    }

    t3 := time.Now()

    l := int64(len(accounts))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql account by date query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return accounts,nil
}
