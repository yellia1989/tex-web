package zone

import (
    "time"
    "sync"
    "errors"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
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
}

var mu sync.Mutex
var ctx context.Context
var conn *dsql.Conn
var zones gcache.Cache
var connError = errors.New("conn数据库没有准备好")

func init() {
    ctx = context.Background()
    zones = gcache.New(0).
        Simple().
        LoaderFunc(func(key interface{}) (interface{}, error) {
            mu.Lock()
            defer mu.Unlock()

            if conn == nil {
                return nil, connError
            }
            if err := conn.PingContext(ctx); err != nil {
                return nil, err
            }
            z := Zone{}
            if err := conn.QueryRowContext(ctx, "SELECT zone.id,zoneid,zonename,openday_fk,logdbhost FROM zone WHERE zoneid=?", key).Scan(&z.Id, &z.ZoneId, &z.Name, &z.OpenDay, &z.DbHost); err != nil {
                return nil, err
            }
            return &z,nil
        }).
        Build()
}

func Cron(now time.Time) {
    if conn == nil {
        var err error
        conn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            log.Errorf("create zone conn err: %s", err.Error())
            return
        }
    }

    mzoneip, _ := gm.RegistryIp()
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
                conn.PingContext(ctx)
                sql := "INSERT INTO zone(zoneid,zonename,openday_fk,logdbhost) VALUES(?,?,?,?)"
                if _,err := conn.ExecContext(ctx, sql, zoneid, new.SZoneName, d.Id, ip); err != nil {
                    log.Errorf("zone cron: %s", err.Error())
                    continue
                }
                log.Debugf("zone cron add new zone: %d", zoneid)
            } else {
                log.Errorf("zone cron: %s", err.Error())
            }
            continue
        }

        old := zone.(*Zone)
        if old.Name != new.SZoneName || old.OpenDay != d.Id || old.DbHost != ip {
            // 更改了分区信息,更新数据库
            conn.PingContext(ctx)
            sql := "UPDATE zone SET "
            first := true
            if old.Name != new.SZoneName {
                sql += "zonename='"+new.SZoneName+"'"
                first = false
                log.Debugf("zone cron update zone: %d, name: %s->%s", zoneid, old.Name, new.SZoneName)
            }
            if old.OpenDay != d.Id {
                if !first {
                    sql += ","
                }
                sql += "openday_fk="+common.U32toa(d.Id)
                first = false
                log.Debugf("zone cron update zone: %d, openDay: %d->%d", zoneid, old.OpenDay, d.Id)
            }
            if old.DbHost != ip {
                if !first {
                    sql += ","
                }
                sql += "logdbhost='"+ip+"'"
                first = false
                log.Debugf("zone cron update zone: %d, dbhost: %s->%s", zoneid, old.DbHost, ip)
            }
            sql += " WHERE id="+common.U32toa(old.Id)
            if _,err := conn.ExecContext(ctx, sql); err != nil {
                log.Errorf("zone cron: %s, sql: %s", err.Error(), sql)
                continue
            }
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
