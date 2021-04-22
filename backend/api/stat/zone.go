package stat

import (
    "fmt"
    "time"
    "sync"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type zone struct {
    ID  uint32      `json:"zoneid_fk"`
    Zoneid uint32   `json:"zoneid"`     // 分区id
    Name string     `json:"zonename"`   // 分区名字
    OpenDay uint16  `json:"openday"`    // 开服累计天数
    openTime time.Time                  // 开服日期
    openTimeID uint32                   // 开服日期id
}

var mzones map[uint32]*zone
var vzones []*zone
var m sync.Mutex
var nextUpdateTime time.Time

func refreshZone() {
    now := time.Now()
    if now.Before(nextUpdateTime) {
        return
    }

    nowDate := getDateByTime(now)
    if nowDate == nil {
        return
    }

    db := cfg.StatDb
    sql := "SELECT a.id,a.zoneid,a.zonename,date.yyyymmdd FROM zone as a,date WHERE a.openday_fk = date.id ORDER BY zoneid desc"
    rows, _ := db.Query(sql)
    defer rows.Close()

    mtmp := make(map[uint32]*zone)
    vtmp := make([]*zone, 0)

    var openDay string
    for rows.Next() {
        var z zone
        if err := rows.Scan(&z.ID, &z.Zoneid, &z.Name, &openDay); err != nil {
            return
        }
        z.openTime = common.ParseTimeInLocal("2006-01-02", openDay)
        z.openTimeID = getDateByTime(z.openTime).ID
        
        if nowDate.ID >= z.openTimeID {
            z.OpenDay = uint16(nowDate.ID - z.openTimeID + 1)
        }
        z.Name = fmt.Sprintf("%s(%d)", z.Name, z.Zoneid)
        mtmp[z.ID] = &z
        vtmp = append(vtmp, &z)
    }

    mzones = mtmp
    vzones = vtmp
    nextUpdateTime = now.Add(time.Minute * 5)
}

// 根据zoneidFk获取分区信息
func getZone(zoneidFk uint32) *zone {
    m.Lock()
    defer m.Unlock()

    refreshZone()

    z, _ := mzones[zoneidFk]
    return z
}

// 获取分区列表
func getAllZone() []*zone {
    m.Lock()
    defer m.Unlock()

    refreshZone()

    zones := make([]*zone, len(vzones))
    copy(zones, vzones)

    return zones
}
