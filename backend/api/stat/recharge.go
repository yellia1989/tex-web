package stat

import (
    "time"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type rgeRecord struct {
    accountidFk uint32
    money uint32
    first bool
}

type rgeRecords []*rgeRecord

// 统计特定日期的每一条充值记录,不去重
func getRgeRecordByDate(vzoneid []uint32, dateBegin uint32, dateEnd uint32, cond *condition) (map[dateZoneKey]*rgeRecords,error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT zoneid_fk,accountid_fk,date_fk,money,first FROM recharge WHERE"
    if len(vzoneid) != 0 {
        sql += " zoneid_fk in ("+ common.U32vtoa(vzoneid,",") +") AND"
    }
    if cond != nil {
        sql += cond.String() + " AND "
    }
    sql += " date_fk BETWEEN ? AND ?"

    t1 := time.Now()

    rows, err := db.Query(sql, dateBegin, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()
    rges := make(map[dateZoneKey]*rgeRecords)
    i := int64(0)
    for rows.Next() {
        var r rgeRecord
        var zoneidFk uint32
        var dateFk uint32
        if err := rows.Scan(&zoneidFk, &r.accountidFk, &dateFk, &r.money, &r.first); err != nil {
            return nil, err
        }

        records, ok := rges[dateZoneKey{dateFk,zoneidFk}]
        if !ok {
            tmp := rgeRecords(make([]*rgeRecord,0))
            records = &tmp
            rges[dateZoneKey{dateFk,zoneidFk}] = records
        }
        *records = append(*records, &r)
        i++
    }

    t3 := time.Now()

    perrow := int64(0)
    if i > 0 {
        perrow = t3.Sub(t2).Microseconds()/i
    }

    log.Debugf("sql query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), i, perrow)

    return rges, nil
}

type rgeDate struct {
    total uint32
    rolenum uint32
}

// 统计到某一天的累计充值和充值人数
func getRgeUtilDate(vzoneid []uint32, dateEnd uint32, cond *condition) (map[uint32]*rgeDate, error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT zoneid_fk,sum(money) as rge_total,count(DISTINCT accountid_fk) as rge_rolenum from recharge WHERE date_fk <= ? "
    if len(vzoneid) != 0 {
        sql += " AND zoneid_fk in ("+ common.U32vtoa(vzoneid,",") +")"
    }
    if cond != nil {
        sql += " AND " + cond.String()
    }
    sql += " GROUP BY zoneid_fk"

    t1 := time.Now()

    rows, err := db.Query(sql, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()
    rges := make(map[uint32]*rgeDate)
    var zoneidFk uint32
    var total uint32
    var rolenum uint32
    for rows.Next() {
        if err := rows.Scan(&zoneidFk, &total, &rolenum); err != nil {
            return nil, err
        }
        rges[zoneidFk] = &rgeDate{total, rolenum}
    }

    t3 := time.Now()

    l := int64(len(rges))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql rge total by date query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return rges, nil
}

// 统计特定日期的累计充值和人数
func getRgeByDate(vzoneid []uint32, dateEnd uint32, cond *condition) (map[dateZoneKey]*rgeDate, error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT zoneid_fk,date_fk,sum(money) as rge_total,count(DISTINCT accountid_fk) as rge_rolenum from recharge WHERE date_fk <= ? "
    if len(vzoneid) != 0 {
        sql += " AND zoneid_fk in ("+ common.U32vtoa(vzoneid,",") +")"
    }
    if cond != nil {
        sql += " AND " + cond.String()
    }
    sql += " GROUP BY zoneid_fk,date_fk"

    t1 := time.Now()

    rows, err := db.Query(sql, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()
    rges := make(map[dateZoneKey]*rgeDate)
    var zoneidFk uint32
    var dateFk uint32
    var total uint32
    var rolenum uint32
    for rows.Next() {
        if err := rows.Scan(&zoneidFk, &dateFk, &total, &rolenum); err != nil {
            return nil, err
        }
        rges[dateZoneKey{dateFk, zoneidFk}] = &rgeDate{total, rolenum}
    }

    t3 := time.Now()

    l := int64(len(rges))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql rge by date query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return rges, nil
}
