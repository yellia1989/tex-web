package zone

import (
    "time"
    "sync"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-web/backend/api/server"
    "github.com/yellia1989/tex-web/backend/api/gm"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cron/date"
)

type Zone struct {
    Id uint32
    ZoneId uint32
    Name string
    OpenDay uint32
    DbHost string
    IsMerge uint32
}

var mu sync.Mutex
var ctx context.Context
var conn *dsql.Conn

var zones gcache.Cache

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

func init() {
    ctx = context.Background()
    zones = gcache.New(0).
        Simple().
        LoaderFunc(func(key interface{}) (interface{}, error) {
            mu.Lock()
            defer mu.Unlock()

            if err := checkConn(); err != nil {
                return nil, err
            }

            z := Zone{}
            if err := conn.QueryRowContext(ctx, "SELECT zone.id,zoneid,zonename,openday_fk,logdbhost,is_merge FROM zone WHERE zoneid=?", key).Scan(&z.Id, &z.ZoneId, &z.Name, &z.OpenDay, &z.DbHost, &z.IsMerge); err != nil {
                return nil, err
            }
            return &z,nil
        }).
        Build()
}

func Cron(now time.Time) {
    mu.Lock()
    if err := checkConn(); err != nil {
        mu.Unlock()
        log.Errorf("cron [zone] check conn err: %s", err.Error())
        return
    }
    mu.Unlock()

    mzoneip, _ := server.RegistryIp()
    mzone := gm.ZoneMap()
    for zoneid, new := range mzone {
        ip, ok := mzoneip[zoneid]
        if !ok {
            continue
        }
        openDayTime := time.Unix(int64(new.IPublishTime),0)
        d := date.Get(openDayTime)
        if d == nil {
            continue
        }

        zone, err := zones.Get(zoneid)
        if err != nil {
            if err == gcache.KeyNotFoundError || err == dsql.ErrNoRows {
                // 新增了分区，需要同步
                // 获取开服时间
                mu.Lock()

                sql := "INSERT INTO zone(zoneid,zonename,openday_fk,logdbhost) VALUES(?,?,?,?)"
                if _,err := conn.ExecContext(ctx, sql, zoneid, new.SZoneName, d.Id, ip); err != nil {
                    mu.Unlock()
                    log.Errorf("cron [zone] add new zone err: %s, zoneid: %d", err.Error(), zoneid)
                    continue
                }

                mu.Unlock()
                log.Debugf("cron [zone] add new zone err: %s, zoneid: %d", err.Error(), zoneid)
            } else {
                log.Errorf("cron [zone] get cache zone err: %s, zoneid: %d", err.Error(), zoneid)
            }
            continue
        }

        old := zone.(*Zone)
        if old.Name != new.SZoneName || old.OpenDay != d.Id || old.DbHost != ip || old.IsMerge != new.IMergeToZoneId {
            // 更改了分区信息,更新数据库
            sql := "UPDATE zone SET "
            first := true
            if old.Name != new.SZoneName {
                sql += "zonename='"+new.SZoneName+"'"
                first = false
                log.Debugf("cron [zone] update zone: %d, name: %s->%s", zoneid, old.Name, new.SZoneName)
            }
            if old.OpenDay != d.Id {
                if !first {
                    sql += ","
                }
                sql += "openday_fk="+common.U32toa(d.Id)
                first = false
                log.Debugf("cron [zone] update zone: %d, openDay: %d->%d", zoneid, old.OpenDay, d.Id)
            }
            if old.DbHost != ip {
                if !first {
                    sql += ","
                }
                sql += "logdbhost='"+ip+"'"
                first = false
                log.Debugf("cron [zone] update zone: %d, dbhost: %s->%s", zoneid, old.DbHost, ip)
            }
            if old.IsMerge != new.IMergeToZoneId {
                if !first {
                    sql += ","
                }
                is_merge := uint32(0)
                if new.IMergeToZoneId != 0 {
                    is_merge = 1
                }
                sql += "is_merge="+common.U32toa(is_merge)
                first = false
                log.Debugf("cron [zone] update zone: %d, is_merge: %d->%d", zoneid, old.IsMerge, new.IMergeToZoneId)
            }
            sql += " WHERE id="+common.U32toa(old.Id)

            mu.Lock()
            if _,err := conn.ExecContext(ctx, sql); err != nil {
                mu.Unlock()
                log.Errorf("cron [zone] update err: %s, sql: %s", err.Error(), sql)
                continue
            }
            mu.Unlock()
            zones.Remove(zoneid)
        }
    }
}

func Get(zoneid uint32) *Zone {
    z, _ := zones.Get(zoneid)
    if z == nil {
        return nil
    }
    return z.(*Zone)
}
