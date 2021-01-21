package sync

import (
    "sync"
    "time"
    "context"
    dsql "database/sql"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
)

var ctx context.Context
var mu sync.Mutex
var conn *dsql.Conn
var zones map[uint32]*zone
var stoping bool
var wg sync.WaitGroup

// 最后一次账号创建时间
var accountMaxTime time.Time
var accountMu sync.Mutex

// 最后一次角色创建时间
var roleMaxTime time.Time
var roleMu sync.Mutex

func isAccountMissed(t time.Time) bool {
    accountMu.Lock()
    defer accountMu.Unlock()
    return t.Sub(accountMaxTime) > time.Minute * 10
}

func UpdateAccountMaxTime(t time.Time) {
    accountMu.Lock()
    defer accountMu.Unlock()
    if t.After(accountMaxTime) {
        accountMaxTime = t
    }
}

func isRoleMissed(t time.Time) bool {
    roleMu.Lock()
    defer roleMu.Unlock()
    return t.Sub(roleMaxTime) > time.Minute * 10
}

func UpdateRoleMaxTime(t time.Time) {
    roleMu.Lock()
    defer roleMu.Unlock()
    if t.After(roleMaxTime) {
        roleMaxTime = t
    }
}

func init() {
    ctx = context.Background()
    zones = make(map[uint32]*zone)
}

func checkConn() (err error) {
    if conn != nil {
        err = conn.PingContext(ctx)
        if err != nil {
            conn.Close()
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

func Cron(now time.Time) {
    if err := checkConn(); err != nil {
        log.Errorf("cron [sync] check conn err: %s", err.Error())
        return
    }

    mu.Lock()
    if stoping {
        mu.Unlock()
        return
    }
    mu.Unlock()

    sql := "SELECT id,zoneid,logdbhost FROM zone"
    rows, err := conn.QueryContext(ctx, sql)
    if err != nil {
        log.Errorf("cron [sync] query err: %s", err.Error())
        return
    }
    defer rows.Close()

    tmp := make(map[uint32]*zone)
    for rows.Next() {
        z := zone{}
        if err := rows.Scan(&z.id, &z.zoneid, &z.dbhost); err != nil {
            return
        }
        z.dur = cfg.LogSyncInterval
        tmp[z.zoneid]= &z
    }
    rows.Close()

    // 增加全局日志同步
    global := zone{}
    global.dur = cfg.LogSyncInterval
    tmp[global.id]= &global

    for zoneid, new := range tmp {
        mu.Lock()
        _, ok := zones[zoneid]
        if ok {
            mu.Unlock()
            continue
        }

        zones[zoneid] = new
        mu.Unlock()

        wg.Add(1)
        go func(z *zone) {
            defer func() {
                mu.Lock()
                delete(zones, z.zoneid)
                mu.Unlock()

                wg.Done()
            }()

            z.run()
        }(new)
    }
}

func Stop() {
    mu.Lock()
    vzone := make([]*zone,0)
    for _, z := range zones {
        vzone = append(vzone, z)
    }
    stoping = true
    mu.Unlock()

    for _, z := range vzone {
        z.stop()
    }

    wg.Wait()

    log.Debug("log sync stop")
}
