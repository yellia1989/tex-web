package stat

import (
    "time"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
)

// 根据分区和日期来获取活跃玩家
func getActiveByDate(vzoneid []uint32, dateBegin uint32, dateEnd uint32, cond *condition) (map[dateZoneKey]uint32,error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT zoneid_fk,date_fk,count(DISTINCT accountid_fk) as active FROM login WHERE"
    if len(vzoneid) != 0 {
        sql += " zoneid_fk in ("+ common.U32vtoa(vzoneid,",") +") AND"
    }
    if cond != nil {
        sql += cond.String() + " AND "
    }
    sql += " date_fk BETWEEN ? AND ? GROUP BY date_fk,zoneid_fk ORDER BY NULL"

    t1 := time.Now()

    rows, err := db.Query(sql, dateBegin, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()

    dateZoneActive := make(map[dateZoneKey]uint32)
    var zoneidFk uint32
    var dateFk uint32
    var active uint32
    for rows.Next() {
        if err := rows.Scan(&zoneidFk, &dateFk, &active); err != nil {
            return nil, err
        }
        dateZoneActive[dateZoneKey{dateFk,zoneidFk}] = active
    }

    t3 := time.Now()

    l := int64(len(dateZoneActive))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return dateZoneActive,nil
}

// 根据分区和日期来获取登陆次数
func getLoginTimesByDate(vzoneid []uint32, dateBegin uint32, dateEnd uint32, cond *condition) (map[dateZoneKey]uint32,error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT zoneid_fk,date_fk,count(accountid_fk) as login_times FROM login WHERE"
    if len(vzoneid) != 0 {
        sql += " zoneid_fk in ("+ common.U32vtoa(vzoneid,",") +") AND"
    }
    if cond != nil {
        sql += cond.String() + " AND "
    }
    sql += " date_fk BETWEEN ? AND ? GROUP BY date_fk,zoneid_fk ORDER BY NULL"

    t1 := time.Now()

    rows, err := db.Query(sql, dateBegin, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()

    dateZoneLoginTimes := make(map[dateZoneKey]uint32)
    var zoneidFk uint32
    var dateFk uint32
    var loginTimes uint32
    for rows.Next() {
        if err := rows.Scan(&zoneidFk, &dateFk, &loginTimes); err != nil {
            return nil, err
        }
        dateZoneLoginTimes[dateZoneKey{dateFk,zoneidFk}] = loginTimes
    }

    t3 := time.Now()

    l := int64(len(dateZoneLoginTimes))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return dateZoneLoginTimes,nil
}
