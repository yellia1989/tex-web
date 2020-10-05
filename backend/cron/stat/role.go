package stat

import (
    "fmt"
    "sync"
    "bytes"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

var roles gcache.Cache
var mu sync.Mutex
var rge90 []uint32
var conn *dsql.Conn

type zoneAccountKey struct {
    zoneidFk uint32
    accountidFk uint32
}

type role struct {
    zoneidFk uint32
    accountidFk uint32
    regDateFk uint32
    login1 uint64
    login2 uint64
    lastLoginDateFk uint32
    rge90 map[int]uint32
    rgeTotal uint32
}

func (r *role) init(rge90 []uint32) {
    r.rge90 = make(map[int]uint32)
    for i, v := range rge90 {
        if v == 0 {
            continue
        }
        r.rge90[i+1] = v
    }
}

func (r *role) rgeDay(d int) uint32 {
    total := uint32(0)
    for i, v := range r.rge90 {
        if i <= d {
            total += v
        }
    }
    return total
}

func (r *role) login(dateFk uint32) {
    if dateFk < r.regDateFk {
        return
    }
    d := dateFk - r.regDateFk
    if d <= 63 {
        r.login1 = r.login1 | (0x01 << d)
    } else {
        r.login2 = r.login2 | (0x01 << (d-63))
    }
    if r.lastLoginDateFk < dateFk {
        r.lastLoginDateFk = dateFk
    }
}

func (r *role) rge(dateFk uint32, money uint32) {
    if dateFk < r.regDateFk {
        return
    }
    d := int(dateFk - r.regDateFk + 1)
    v, ok := r.rge90[d]
    if !ok {
        r.rge90[d] = money
    } else {
        r.rge90[d] = v + money
    }
    r.rgeTotal += money
}

// replace into role(zoneid_fk,accountid_fk,login_1,login_2,last_login_date_fk,rge_total_3,rge_total_7,rge_total_14,rge_total_30,rge_total,rge_day1,rge_day2,...)
func (r *role) str(buff *bytes.Buffer) {
    buff.WriteString("UPDATE role SET ")
    first := true
    if r.login1 != 0 {
        buff.WriteString(fmt.Sprintf("login_1=%d",r.login1))
        first = false
    }
    if r.login2 != 0 {
        if !first {
            buff.WriteString(",")
        }
        buff.WriteString(fmt.Sprintf("login_2=%d",r.login2))
        first = false
    }
    if !first {
        buff.WriteString(",")
    }
    buff.WriteString(fmt.Sprintf("last_login_date_fk=%d", r.lastLoginDateFk))
    first = false

    buff.WriteString(fmt.Sprintf(",rge_total_3=%d", r.rgeDay(3)))
    buff.WriteString(fmt.Sprintf(",rge_total_7=%d", r.rgeDay(7)))
    buff.WriteString(fmt.Sprintf(",rge_total_14=%d", r.rgeDay(14)))
    buff.WriteString(fmt.Sprintf(",rge_total_30=%d", r.rgeDay(30)))
    buff.WriteString(fmt.Sprintf(",rge_total=%d", r.rgeTotal))

    for d, v := range r.rge90 {
        buff.WriteString(fmt.Sprintf(",rge_day%d=%d", d, v))
    }
    buff.WriteString(fmt.Sprintf(" WHERE zoneid_fk=%d AND accountid_fk=%d;", r.zoneidFk, r.accountidFk))
}

func init() {
    rge90 = make([]uint32,90,90)
    roles = gcache.New(20000).
        LRU().
        LoaderFunc(func(key interface{}) (interface{}, error) {
            mu.Lock()
            defer mu.Unlock()

            if conn == nil {
                var err error
                conn, err = cfg.StatDb.Conn(ctx)
                if err != nil {
                    return nil, err
                }
            }

            if err := conn.PingContext(ctx); err != nil {
                return nil, err
            }
            zaKey := key.(zoneAccountKey)
            r := role{}
            r.zoneidFk = zaKey.zoneidFk
            r.accountidFk = zaKey.accountidFk
            if err := conn.QueryRowContext(ctx, "SELECT reg_date_fk,login_1,login_2,last_login_date_fk,rge_total,rge_day1,rge_day2,rge_day3,rge_day4,rge_day5,rge_day6,rge_day7,rge_day8,rge_day9,rge_day10,rge_day11,rge_day12,rge_day13,rge_day14,rge_day15,rge_day16,rge_day17,rge_day18,rge_day19,rge_day20,rge_day21,rge_day22,rge_day23,rge_day24,rge_day25,rge_day26,rge_day27,rge_day28,rge_day29,rge_day30,rge_day31,rge_day32,rge_day33,rge_day34,rge_day35,rge_day36,rge_day37,rge_day38,rge_day39,rge_day40,rge_day41,rge_day42,rge_day43,rge_day44,rge_day45,rge_day46,rge_day47,rge_day48,rge_day49,rge_day50,rge_day51,rge_day52,rge_day53,rge_day54,rge_day55,rge_day56,rge_day57,rge_day58,rge_day59,rge_day60,rge_day61,rge_day62,rge_day63,rge_day64,rge_day65,rge_day66,rge_day67,rge_day68,rge_day69,rge_day70,rge_day71,rge_day72,rge_day73,rge_day74,rge_day75,rge_day76,rge_day77,rge_day78,rge_day79,rge_day80,rge_day81,rge_day82,rge_day83,rge_day84,rge_day85,rge_day86,rge_day87,rge_day88,rge_day89,rge_day90 FROM role WHERE zoneid_fk = ? AND accountid_fk = ?", zaKey.zoneidFk, zaKey.accountidFk).Scan(&r.regDateFk, &r.login1, &r.login2, &r.lastLoginDateFk, &r.rgeTotal, &rge90[0], &rge90[1], &rge90[2], &rge90[3], &rge90[4], &rge90[5], &rge90[6], &rge90[7], &rge90[8], &rge90[9], &rge90[10], &rge90[11], &rge90[12], &rge90[13], &rge90[14], &rge90[15], &rge90[16], &rge90[17], &rge90[18], &rge90[19], &rge90[20], &rge90[21], &rge90[22], &rge90[23], &rge90[24], &rge90[25], &rge90[26], &rge90[27], &rge90[28], &rge90[29], &rge90[30], &rge90[31], &rge90[32], &rge90[33], &rge90[34], &rge90[35], &rge90[36], &rge90[37], &rge90[38], &rge90[39], &rge90[40], &rge90[41], &rge90[42], &rge90[43], &rge90[44], &rge90[45], &rge90[46], &rge90[47], &rge90[48], &rge90[49], &rge90[50], &rge90[51], &rge90[52], &rge90[53], &rge90[54], &rge90[55], &rge90[56], &rge90[57], &rge90[58], &rge90[59], &rge90[60], &rge90[61], &rge90[62], &rge90[63], &rge90[64], &rge90[65], &rge90[66], &rge90[67], &rge90[68], &rge90[69], &rge90[70], &rge90[71], &rge90[72], &rge90[73], &rge90[74], &rge90[75], &rge90[76], &rge90[77], &rge90[78], &rge90[79], &rge90[80], &rge90[81], &rge90[82], &rge90[83], &rge90[84], &rge90[85], &rge90[86], &rge90[87], &rge90[88], &rge90[89]); err != nil {
                return nil, err
            }
            r.init(rge90)
            return &r,nil
        }).
        Build()
}

func get(zoneid uint32, accountid uint32) *role {
    r, err := roles.Get(zoneAccountKey{zoneid, accountid})
    if r == nil {
        if err != dsql.ErrNoRows {
            log.Errorf("get stat role err: %s", err.Error())
        }
        return nil
    }
    return r.(*role)
}
