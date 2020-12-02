package stat

import (
    "fmt"
    "time"
    "sync"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type date struct {
    ID  uint32      `json:"id"`
    Time    string  `json:"time"`
    Desc    string  `json:"mark"`
    time    time.Time
}

var mdates map[uint32]*date
var mdates2 map[uint32]*date
var vdates []*date
var vmarkdates []*date
var mutex sync.Mutex

func refreshDate() {
    db := cfg.StatDb

    if db == nil {
        panic("连接数据库失败")
    }

    sql := "SELECT id,yyyymmdd,`desc`,ymd FROM `date` order by ymd desc"
    rows, err := db.Query(sql)
    if err != nil {
        panic(fmt.Sprintf("sql: %s, err: %s", sql, err.Error()))
        return
    }
    defer rows.Close()

    mtmp := make(map[uint32]*date)
    mtmp2 := make(map[uint32]*date)
    vtmp := make([]*date, 0)
    vtmp2 := make([]*date, 0)

    var ymd uint32
    for rows.Next() {
        var d date
        if err := rows.Scan(&d.ID, &d.Time, &d.Desc, &ymd); err != nil {
            return
        }
        d.time = common.ParseTimeInLocal("2006-01-02", d.Time)
        mtmp[d.ID] = &d
        mtmp2[ymd] = &d
        vtmp = append(vtmp, &d)
        if len(d.Desc) != 0 {
            vtmp2 = append(vtmp2, &d)
        }
    }

    mdates = mtmp
    mdates2 = mtmp2
    vdates = vtmp
    vmarkdates = vtmp2
}

// 根据dateFk获取日期信息
func getDate(id uint32) *date {
    mutex.Lock()
    defer mutex.Unlock()

    refreshDate()

    d, _ := mdates[id]
    return d
}

// 根据当前时间获取日期信息
func getDateByTime(t time.Time) *date {
    mutex.Lock()
    defer mutex.Unlock()

    refreshDate()

    d, _ := mdates2[common.Atou32(t.Format("20060102"))]
    return d
}

// 根据字符串20200915获取日期信息
func getDateByString(t string) *date {
    mutex.Lock()
    defer mutex.Unlock()

    refreshDate()

    d, _ := mdates2[common.Atou32(t)]
    return d
}

// 获取统计开始时间
func getMinDate() *date {
    mutex.Lock()
    defer mutex.Unlock()

    refreshDate()

    return vdates[len(vdates)-1]
}

// 获取最新统计时间
func getMaxDate() *date {
    mutex.Lock()
    defer mutex.Unlock()

    refreshDate()

    return vdates[0]
}

// 获取标注过的日期
func getMarkDates() []*date {
    mutex.Lock()
    defer mutex.Unlock()

    refreshDate()

    marks := make([]*date,len(vmarkdates),len(vmarkdates))
    copy(marks, vmarkdates)

    return marks
}
