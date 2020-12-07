package cron

import (
    "fmt"
    "time"
    dsql "database/sql"
    ssync "sync"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/cron/welfare"
    "github.com/yellia1989/tex-web/backend/cron/date"
    "github.com/yellia1989/tex-web/backend/cron/zone"
    "github.com/yellia1989/tex-web/backend/cron/sync"
    "github.com/yellia1989/tex-web/backend/cron/stat"
)

var wg ssync.WaitGroup
var stop chan bool

func init() {
    stop = make(chan bool)
}

type croner func(now time.Time)

func startCron(name string, f croner, d time.Duration) {
    wg.Add(1)
    go func(stop <- chan bool, name string) {
        defer wg.Done()

        ticker := time.NewTicker(d)
        defer ticker.Stop()
        for {
            select {
            case <- stop:
                log.Debug("%s cron stop", name)
                return
            case t := <- ticker.C: 
                f(t)
            }
        }
    }(stop, name)
}

func startLogSync(d time.Duration) {
    wg.Add(1)
    go func(stop <- chan bool) {
        defer wg.Done()

        ticker := time.NewTicker(d)
        defer ticker.Stop()
        for {
            select {
            case <- stop:
                sync.Stop()
                return
            case t := <- ticker.C: 
                sync.Cron(t)
            }
        }
    }(stop)
}

func Start() {
    // 初始化账号，角色最大时间
    db := cfg.StatDb
    if db == nil {
        panic("stat db nil")
    }

    var accountDate dsql.NullString
    var accountDaytime dsql.NullInt64
    var accountMaxTime time.Time
    if err := db.QueryRow("SELECT yyyymmdd,daytime from date,(select date_fk,daytime FROM account, (SELECT max(id) as maxid from account) a where id = a.maxid) b where date.id = b.date_fk").Scan(&accountDate, &accountDaytime); err != nil {
        if err != dsql.ErrNoRows {
            panic(fmt.Sprintf("cron start err: %s", err.Error()))
        }
    }
    if len(accountDate.String) != 0 {
        accountMaxTime = common.ParseTimeInLocal("2006-01-02", accountDate.String).Add(time.Second * time.Duration(accountDaytime.Int64))
        sync.UpdateAccountMaxTime(accountMaxTime)
    }
    var roleDate dsql.NullString
    var roleDaytime dsql.NullInt64
    var roleMaxTime time.Time
    if err := db.QueryRow("SELECT yyyymmdd,daytime from date,(select reg_date_fk,daytime FROM role, (SELECT max(id) as maxid from role) a where id = a.maxid) b where date.id = b.reg_date_fk").Scan(&roleDate, &roleDaytime); err != nil {
        if err != dsql.ErrNoRows {
            panic(fmt.Sprintf("cron start err: %s", err.Error()))
        }
    }
    if len(roleDate.String) != 0 {
        roleMaxTime = common.ParseTimeInLocal("2006-01-02", roleDate.String).Add(time.Second * time.Duration(roleDaytime.Int64))
        sync.UpdateRoleMaxTime(roleMaxTime)
    }

    log.Debugf("max account time: %s, max role time: %s", accountMaxTime.Format("2006-01-02 15:04:05"), roleMaxTime.Format("2006-01-02 15:04:05"))

    // 免费福利
    startCron("welfare", welfare.Cron, time.Second * 5)

    // 日期
    startCron("date", date.Cron, time.Second * 5)

    // 分区
    startCron("zone", zone.Cron, time.Second * 5)

    // 开启日志
    startLogSync(time.Second * 5)

    // 玩家登陆和充值
    startCron("stat", stat.Cron, cfg.LogStatInterval)

    log.Debug("all cron start")
}

func Stop() {
    close(stop)
    wg.Wait()

    log.Debug("all cron stop")
}
