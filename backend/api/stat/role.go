package stat

import (
    "fmt"
    "time"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

type role struct {
    zoneidFk uint32             // 分区id
    accountidFk uint32          // 账号id
    regDateFk uint32            // 角色创建时间
    login1 uint64               // 90日登陆1
    login2 uint64               // 90日登陆2
    lastLoginDateFk uint32      // 最后一次登陆日期
    rgeTotal3 uint32            // 3日累计充值分
    rgeTotal7 uint32            // 7日累计充值分
    rgeTotal14 uint32           // 14日累计充值分
    rgeTotal30 uint32           // 30日累计充值分
    rgeTotal uint32             // 累计充值分
    rgeDay []uint32             // 90天充值分
    rgeMDay map[uint8]uint8     // 90天充值映射day->index
}

func (r *role) isActive(d uint32) bool {
    d -= 1
    if d <= 63 {
        return (r.login1 & (1 << d)) > 0
    }
    return (r.login2 & (1 << (d-63))) > 0
}

func (r *role) isWeekActive() bool {
    for i := uint32(1); i < 7; i++ {
        if r.isActive(i+1) {
            return true
        }
    }
    return false
}

func (r *role) isDWeekActive() bool {
    for i := uint32(1); i < 14; i++ {
        if r.isActive(i+1) {
            return true
        }
    }
    return false
}

func (r *role) isMonthActive() bool {
    for i := uint32(1); i < 31; i++ {
        if r.isActive(i+1) {
            return true
        }
    }
    return false
}

func (r *role) isR() bool {
    return r.rgeTotal >= cfg.StatRmoney
}

func (r *role) getRge(d uint32) uint32 {
    i,ok := r.rgeMDay[uint8(d)]
    if !ok {
        return 0
    }
    return r.rgeDay[i]
}

func (r *role) String() string {
    actives := make(map[uint8]bool)
    for i := uint32(0); i < 90; i++ {
        if r.isActive(i+1) {
            actives[uint8(i+1)] = true
        }
    }
    return fmt.Sprintf("zoneidfk:%d,accountidfk:%d,regdatefk:%d,actives:%v,lastlogindatefk:%d,rgetotal7:%d,rgetotal30:%d,rgetotal:%d,rgeday:%v,rgemday:%v", r.zoneidFk, r.accountidFk, r.regDateFk, actives, r.lastLoginDateFk, r.rgeTotal7, r.rgeTotal30, r.rgeTotal, r.rgeDay, r.rgeMDay)
}

// 统计到目前为止总创角
func getRoleNumUtilNow(cond *condition) (uint32,error) {
    now := getDateByTime(time.Now())

    m, err := getRoleNumUtilDate(nil, now.ID, cond)
    if err != nil {
        return 0, err
    }

    var total uint32
    for _, v := range m {
        total += v
    }
    return total,nil
}

// 统计到某一天为止总创角
func getRoleNumUtilDate(vzoneid []uint32, dateEnd uint32, cond *condition) (map[uint32]uint32, error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    t1 := time.Now()

    sql := "SELECT zoneid_fk,count(*) FROM role WHERE reg_date_fk <= ?"
    if len(vzoneid) != 0 {
        sql += " AND zoneid_fk IN(" + common.U32vtoa(vzoneid, ",") + ")"
    }
    if cond != nil {
        sql += " AND " + cond.String()
    }
    sql += " GROUP BY zoneid_fk"
    rows, err := db.Query(sql, dateEnd)
    if err != nil {
        return nil,err
    }
    defer rows.Close()

    m := make(map[uint32]uint32)
    var zoneidFk uint32
    var total uint32
    for rows.Next() {
        if err := rows.Scan(&zoneidFk, &total); err != nil {
            return nil, err
        }
        m[zoneidFk] = total
    }

    t2 := time.Now()

    log.Debugf("sql role total by date query: %v", t2.Sub(t1))

    return m, nil
}

// 统计每一天创建的角色
func getRoleNumByDate(vzoneid []uint32, dateBegin uint32, dateEnd uint32, cond *condition) (map[dateZoneKey]uint32,error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    t1 := time.Now()

    sql := "SELECT reg_date_fk,zoneid_fk,count(*) FROM role WHERE reg_date_fk BETWEEN ? AND ?"
    if len(vzoneid) != 0 {
        sql += " AND zoneid_fk IN(" + common.U32vtoa(vzoneid, ",") + ")"
    }
    if cond != nil {
        sql += " AND " + cond.String()
    }
    sql += " GROUP BY reg_date_fk,zoneid_fk ORDER BY NULL"
    rows, err := db.Query(sql, dateBegin, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()

    var dateFk uint32
    var roleNum uint32
    var zoneidFk uint32
    roles := make(map[dateZoneKey]uint32)
    for rows.Next() {
        if err := rows.Scan(&dateFk, &zoneidFk, &roleNum); err != nil {
            return nil,err
        }
        roles[dateZoneKey{dateFk,zoneidFk}] = roleNum
    }

    t3 := time.Now()

    l := int64(len(roles))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql role by date query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return roles,nil
}

// 根据分区和创角时间来搜索玩家
func getRoles(vzoneid []uint32, dateBegin uint32, dateEnd uint32, global bool, cond *condition) ([]*role,error) {
    db := cfg.StatDb
    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT zoneid_fk,accountid_fk,reg_date_fk,login_1,login_2,last_login_date_fk,rge_total_3,rge_total_7,rge_total_14,rge_total_30,rge_total,rge_day1,rge_day2,rge_day3,rge_day4,rge_day5,rge_day6,rge_day7,rge_day8,rge_day9,rge_day10,rge_day11,rge_day12,rge_day13,rge_day14,rge_day15,rge_day16,rge_day17,rge_day18,rge_day19,rge_day20,rge_day21,rge_day22,rge_day23,rge_day24,rge_day25,rge_day26,rge_day27,rge_day28,rge_day29,rge_day30,rge_day31,rge_day32,rge_day33,rge_day34,rge_day35,rge_day36,rge_day37,rge_day38,rge_day39,rge_day40,rge_day41,rge_day42,rge_day43,rge_day44,rge_day45,rge_day46,rge_day47,rge_day48,rge_day49,rge_day50,rge_day51,rge_day52,rge_day53,rge_day54,rge_day55,rge_day56,rge_day57,rge_day58,rge_day59,rge_day60,rge_day61,rge_day62,rge_day63,rge_day64,rge_day65,rge_day66,rge_day67,rge_day68,rge_day69,rge_day70,rge_day71,rge_day72,rge_day73,rge_day74,rge_day75,rge_day76,rge_day77,rge_day78,rge_day79,rge_day80,rge_day81,rge_day82,rge_day83,rge_day84,rge_day85,rge_day86,rge_day87,rge_day88,rge_day89,rge_day90 FROM role WHERE"

    if len(vzoneid) != 0 {
        sql += " zoneid_fk in ("+ common.U32vtoa(vzoneid, ",") + ") AND"
    }
    sql += " reg_date_fk BETWEEN ? AND ?"

    if global {
        // 全服去重
        sql += " AND `first`=1"
    }

    if cond != nil {
        sql += " AND " + cond.String()
    }

    t1 := time.Now()

    rows, err := db.Query(sql, dateBegin, dateEnd)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    t2 := time.Now()

    var roles []*role
    rgeDay := make([]uint32,90)
    for rows.Next() {
        var r role
        if err := rows.Scan(&r.zoneidFk, &r.accountidFk, &r.regDateFk, &r.login1, &r.login2, &r.lastLoginDateFk, &r.rgeTotal3, &r.rgeTotal7, &r.rgeTotal14, &r.rgeTotal30, &r.rgeTotal, &rgeDay[0], &rgeDay[1], &rgeDay[2], &rgeDay[3], &rgeDay[4], &rgeDay[5], &rgeDay[6], &rgeDay[7], &rgeDay[8], &rgeDay[9], &rgeDay[10], &rgeDay[11], &rgeDay[12], &rgeDay[13], &rgeDay[14], &rgeDay[15], &rgeDay[16], &rgeDay[17], &rgeDay[18], &rgeDay[19], &rgeDay[20], &rgeDay[21], &rgeDay[22], &rgeDay[23], &rgeDay[24], &rgeDay[25], &rgeDay[26], &rgeDay[27], &rgeDay[28], &rgeDay[29], &rgeDay[30], &rgeDay[31], &rgeDay[32], &rgeDay[33], &rgeDay[34], &rgeDay[35], &rgeDay[36], &rgeDay[37], &rgeDay[38], &rgeDay[39], &rgeDay[40], &rgeDay[41], &rgeDay[42], &rgeDay[43], &rgeDay[44], &rgeDay[45], &rgeDay[46], &rgeDay[47], &rgeDay[48], &rgeDay[49], &rgeDay[50], &rgeDay[51], &rgeDay[52], &rgeDay[53], &rgeDay[54], &rgeDay[55], &rgeDay[56], &rgeDay[57], &rgeDay[58], &rgeDay[59], &rgeDay[60], &rgeDay[61], &rgeDay[62], &rgeDay[63], &rgeDay[64], &rgeDay[65], &rgeDay[66], &rgeDay[67], &rgeDay[68], &rgeDay[69], &rgeDay[70], &rgeDay[71], &rgeDay[72], &rgeDay[73], &rgeDay[74], &rgeDay[75], &rgeDay[76], &rgeDay[77], &rgeDay[78], &rgeDay[79], &rgeDay[80], &rgeDay[81], &rgeDay[82], &rgeDay[83], &rgeDay[84], &rgeDay[85], &rgeDay[86], &rgeDay[87], &rgeDay[88], &rgeDay[89]); err != nil {
            return nil, err
        }
        i := uint8(0)
        for d,v := range rgeDay {
            if v == 0 {
                continue
            }
            if len(r.rgeDay) == 0 {
                r.rgeDay = make([]uint32,0)
                r.rgeMDay = make(map[uint8]uint8,0)
            }
            r.rgeDay = append(r.rgeDay, v)
            r.rgeMDay[uint8(d+1)] = i
            i++
        }
        roles = append(roles, &r)
    }

    t3 := time.Now()

    l := int64(len(roles))
    perrow := int64(0)
    if l > 0 {
        perrow = t3.Sub(t2).Microseconds()/l
    }

    log.Debugf("sql query: %v, scan: %v, rows: %d, perrow: %d micr", t2.Sub(t1), t3.Sub(t2), l, perrow)

    return roles, nil
}
