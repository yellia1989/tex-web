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
var conn *dsql.Conn
var mu sync.Mutex
var wg sync.WaitGroup
var zones map[uint32]*zone

// 最后一次账号创建时间
var accountMaxTime time.Time
var accountMu sync.Mutex

// 最后一次角色创建时间
var roleMaxTime time.Time
var roleMu sync.Mutex

func isAccountMissed(t time.Time) bool {
    accountMu.Lock()
    defer accountMu.Unlock()
    return t.Before(accountMaxTime)
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
    return t.Before(roleMaxTime)
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

func Cron(now time.Time) {
    if conn == nil {
        var err error
        conn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            log.Errorf("create sync conn err: %s", err.Error())
            return
        }
    }

    log.Debugf("db stat: %v", cfg.StatDb.Stats())

    if err := conn.PingContext(ctx); err != nil {
        conn.Close()
        log.Debugf("sync ping err: %s, try to create conn again", err.Error())

        conn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            log.Errorf("create sync conn err: %s", err.Error())
            return
        }
    }

    sql := "SELECT id,zoneid,logdbhost FROM zone"
    rows, err := conn.QueryContext(ctx, sql)
    if err != nil {
        return
    }
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
        mu.Unlock()

        // 检测到新分区，开启同步
        if err := new.init(); err != nil {
            log.Errorf("zone init err: %s, zoneid: %d", err.Error(), zoneid)
            continue
        }

        mu.Lock()
        zones[zoneid] = new
        mu.Unlock()

        wg.Add(1)
        go func(z *zone, zones *map[uint32]*zone) {
            defer wg.Done()

            z.run()

            mu.Lock()
            delete(*zones, z.zoneid)
            mu.Unlock()
        }(new, &zones)
    }
}

func Stop() {
    mu.Lock()
    vzone := make([]*zone,0)
    for _, z := range zones {
        vzone = append(vzone, z)
    }
    mu.Unlock()

    for _, z := range vzone {
        z.stop()
    }

    wg.Wait()

    log.Debug("log sync stop")
}
