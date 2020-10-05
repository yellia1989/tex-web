package date

import (
    "time"
    "sync"
    "errors"
    "context"
    "github.com/bluele/gcache"
    dsql "database/sql"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

var mu sync.Mutex
var ctx context.Context
var conn *dsql.Conn
var dates gcache.Cache
var dateMinError = errors.New("日期小于最小日期")
var connError = errors.New("conn数据库没有准备好")
var dateMax time.Time

type Date struct {
    Id uint32
    Yyyymmdd string
    Ymd uint32
}

func init() {
    ctx = context.Background()
    dates = gcache.New(30).
        LFU().
        LoaderFunc(func(key interface{}) (interface{}, error) {
            keyInt := key.(uint32)
            t := common.ParseTimeInLocal("20060102", common.U32toa(keyInt))
            if t.Before(cfg.LogDateMin) {
                return nil, dateMinError
            }
            mu.Lock()
            defer mu.Unlock()

            if conn == nil {
                return nil, connError
            }
            if err := conn.PingContext(ctx); err != nil {
                return nil, err
            }
            d := Date{}
            if err := conn.QueryRowContext(ctx, "SELECT id,yyyymmdd,ymd FROM date WHERE ymd=?", key).Scan(&d.Id, &d.Yyyymmdd, &d.Ymd); err != nil {
                return nil, err
            }
            return &d,nil
        }).
        Build()
}

func Cron(now time.Time) {
    if conn == nil {
        var err error
        conn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            log.Errorf("create date conn err: %s", err.Error())
            return
        }
    }
    if dateMax.IsZero() {
        conn.PingContext(ctx)
        var t dsql.NullInt32
        if err := conn.QueryRowContext(ctx, "SELECT max(ymd) FROM date").Scan(&t); err != nil {
            log.Errorf("date cron: %s", err.Error())
            return
        }
        if t.Valid {
            dateMax = common.ParseTimeInLocal("20060102", common.U32toa(uint32(t.Int32)))
        } else {
            dateMax = cfg.LogDateMin
        }
    }

    dateFrom := dateMax
    dateTo := now

    for i := dateFrom; i.After(dateTo) == false; i = i.Add(time.Hour*24) {
        _, err := dates.Get(common.Atou32(i.Format("20060102")))
        if err == nil {
            continue
        }

        if err != dsql.ErrNoRows {
            log.Errorf("get date err: %s", err.Error())
            return
        }

        // 需要创建新的date
        if err := conn.PingContext(ctx); err != nil {
            log.Errorf("date conn err: %s", err.Error())
            return
        }

        yyyymmdd := i.Format("2006-01-02")
        ymd := common.Atou32(i.Format("20060102"))
        year,month,day := i.Date()
        week := int(i.Weekday())
        if week == 0 {
            week = 7
        }

        sql := "INSERT INTO `date`(yyyymmdd,`desc`,`year`,`month`,`day`,`week`,ymd) VALUES(?,'',?,?,?,?,?)"
        if _, err := conn.ExecContext(ctx, sql, yyyymmdd, year, month, day, week, ymd); err != nil {
            log.Errorf("create date err: %s", err.Error())
            return
        }
        log.Debugf("create date: %s", yyyymmdd)
    }
}

func Get(t time.Time) *Date {
    d, err := dates.Get(common.Atou32(t.Format("20060102")))
    if d == nil {
        if err != dsql.ErrNoRows {
            log.Errorf("get date err: %s", err.Error())
        }
        return nil
    }
    return d.(*Date)
}
