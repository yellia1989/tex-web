package sync

import (
    "fmt"
    "time"
    dsql "database/sql"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

type tabler interface {
    sync(from *dsql.DB, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error
}

type zone struct {
    id uint32
    zoneid uint32
    dbhost string   // 日志数据库地址
    quit chan bool // 结束标识
    dur time.Duration // 日志同步间隔
    tables []tabler // 需要同步的表
    fromdb *dsql.DB
    toconn *dsql.Conn
}

func (z *zone) init() (err error) {
    if z.zoneid != 0 {
        z.fromdb, err = dsql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/log_zone_%d?multiStatements=true", cfg.LogDbUser, cfg.LogDbPwd, cfg.LogDbHost, z.zoneid))
        if err != nil {
            return
        }
    } else {
        z.fromdb = cfg.LogDb
    }

    todb := cfg.StatDb
    z.toconn, err = todb.Conn(ctx)
    if err != nil {
        return
    }

    z.quit = make(chan bool)
    z.tables = make([]tabler,0)
    
    if z.zoneid != 0 {
        z.tables = append(z.tables, &login{})
        z.tables = append(z.tables, &logout{})
        z.tables = append(z.tables, &recharge{})
    } else {
        z.tables = append(z.tables, &account{})
        z.tables = append(z.tables, &role{})
    }

    return nil
}

func (z *zone) checkConn() (err error){
    err = z.fromdb.PingContext(ctx)
    if err != nil {
        return
    }

    err = z.toconn.PingContext(ctx)
    return
}

func (z *zone) run() {
    defer func() {
        if z.zoneid != 0 && z.fromdb != nil {
            z.fromdb.Close()
        }
        if z.toconn != nil {
            z.toconn.Close()
        }
        log.Infof("cron [sync][zone] stop, zoneid: %d", z.zoneid)
    }()

    log.Infof("cron [sync][zone] start, zoneid: %d", z.zoneid)

    if err := z.init(); err != nil {
        log.Errorf("cron [sync][zone] init err: %s, zoneid: %d", err.Error(), z.zoneid)
        return
    }

    ticker := time.NewTicker(z.dur)
    defer ticker.Stop()

    for {
        select {
        case <- z.quit: {
            return
        }
        case <- ticker.C: {
            if err := z.checkConn(); err != nil {
                log.Errorf("cron [sync][zone] check conn err: %s", err.Error())
                return
            }
            for _, t := range z.tables {
                if err := t.sync(z.fromdb, z.toconn, z.zoneid, z.id); err != nil {
                    log.Errorf("cron [sync][zone] sync err: %s, zoneid: %d", err.Error(), z.zoneid)
                    return
                }
            }
        }
        }
    }
}

func (z *zone) stop() {
    close(z.quit)
}
