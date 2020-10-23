package sync

import (
    "fmt"
    "time"
    dsql "database/sql"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

type tabler interface {
    sync(from *dsql.Conn, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error
}

type zone struct {
    id uint32
    zoneid uint32
    dbhost string   // 日志数据库地址
    quit chan bool // 结束标识
    dur time.Duration // 日志同步间隔
    tables []tabler // 需要同步的表
    fromconn *dsql.Conn
    toconn *dsql.Conn
}

func (z *zone) init() error {
    var err error
    var fromdb *dsql.DB

    if z.zoneid != 0 {
        fromdb, err = dsql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/log_zone_%d?multiStatements=true", cfg.LogDbUser, cfg.LogDbPwd, z.dbhost, z.zoneid))
        if err != nil {
            return err
        }
    } else {
        fromdb = cfg.LogDb
    }
    z.fromconn, err = fromdb.Conn(ctx)
    if err != nil {
        return err
    }

    todb := cfg.StatDb
    z.toconn, err = todb.Conn(ctx)
    if err != nil {
        // 释放conn连接
        z.fromconn.Close()
        return err
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

func (z *zone) run() {
    defer func() {
        z.fromconn.Close()
        z.toconn.Close()
        log.Debugf("zone sync stop, zoneid: %d", z.zoneid)
    }()

    log.Debugf("zone sync start, zoneid: %d", z.zoneid)

    for {
        ticker := time.NewTicker(z.dur)
        defer ticker.Stop()

        select {
        case <- z.quit: {
            return
        }
        case <- ticker.C: {
            for _, t := range z.tables {
                if err := t.sync(z.fromconn, z.toconn, z.zoneid, z.id); err != nil {
                    log.Errorf("zone sync err: %s, zoneid: %d", err.Error(), z.zoneid)
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
