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

var ctx context.Context
var mu sync.Mutex
var conn *dsql.Conn

var dates gcache.Cache
var dateMinError = errors.New("日期小于最小日期")
var dateMax time.Time

type Date struct {
    Id uint32
    Yyyymmdd string
    Ymd uint32
}

func checkConn() (err error) {
    if conn != nil {
        err = conn.PingContext(ctx)
        if err != nil {
            conn = nil
        } else {
            return
        }
    }

    if conn == nil {
        conn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            return
        }
    }
    return
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

            if err := checkConn(); err != nil {
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
    mu.Lock()

    if err := checkConn(); err != nil {
        mu.Unlock()
        log.Errorf("cron [date] check conn err: %s", err.Error())
        return
    }

    if dateMax.IsZero() {
        var t dsql.NullInt32
        if err := conn.QueryRowContext(ctx, "SELECT max(ymd) FROM date").Scan(&t); err != nil {
            mu.Unlock()
            log.Errorf("cron [date] query err: %s", err.Error())
            return
        }
        if t.Valid {
            dateMax = common.ParseTimeInLocal("20060102", common.U32toa(uint32(t.Int32)))
        } else {
            dateMax = cfg.LogDateMin
        }
    }
    mu.Unlock()

    dateFrom := dateMax
    dateTo := now

    for i := dateFrom; i.After(dateTo) == false; i = i.Add(time.Hour*24) {
        _, err := dates.Get(common.Atou32(i.Format("20060102")))
        if err == nil {
            continue
        }

        if err != dsql.ErrNoRows {
            log.Errorf("cron [date] get cache date err: %s", err.Error())
            return
        }

        // 需要创建新的date
        yyyymmdd := i.Format("2006-01-02")
        ymd := common.Atou32(i.Format("20060102"))
        year,month,day := i.Date()
        week := int(i.Weekday())
        if week == 0 {
            week = 7
        }

        mu.Lock()
        if err := checkConn(); err != nil {
            mu.Unlock()
            log.Errorf("cron [date] check conn, err: %s", err.Error())
            return
        }
        sql := "INSERT INTO `date`(yyyymmdd,`desc`,`year`,`month`,`day`,`week`,ymd) VALUES(?,'',?,?,?,?,?)"
        if _, err := conn.ExecContext(ctx, sql, yyyymmdd, year, month, day, week, ymd); err != nil {
            mu.Unlock()
            log.Errorf("cron [date] create date err: %s", err.Error())
            return
        }
        mu.Unlock()
        log.Debugf("cron [date] create date: %s", yyyymmdd)
    }
}

func Get(t time.Time) *Date {
    d, err := dates.Get(common.Atou32(t.Format("20060102")))
    if d == nil {
        if err != dsql.ErrNoRows {
            log.Errorf("cron [date] get cache date err: %s", err.Error())
        }
        return nil
    }
    return d.(*Date)
}
