package stat

import (
    "sync"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

var ctx context.Context
var mu sync.Mutex
var conn *dsql.Conn

var roles gcache.Cache

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

type zoneAccountKey struct {
    zoneidFk uint32
    accountidFk uint32
}

type Role struct {
    RegDateFk uint32
}

func init() {
    ctx = context.Background()
    roles = gcache.New(100000).
        LFU().
        LoaderFunc(func(key interface{}) (interface{}, error) {
            mu.Lock()
            defer mu.Unlock()

            if err := checkConn(); err != nil {
                return nil, err
            }

            zaKey := key.(zoneAccountKey)
            r := Role{}
            if err := conn.QueryRowContext(ctx, "SELECT reg_date_fk FROM role WHERE zoneid_fk=? AND accountid_fk=?", zaKey.zoneidFk, zaKey.accountidFk).Scan(&r.RegDateFk); err != nil {
                return nil, err
            }
            return &r,nil
        }).
        Build()
}

func Get(zoneid uint32, accountid uint32) *Role {
    r, err := roles.Get(zoneAccountKey{zoneid, accountid})
    if r == nil {
        if err != dsql.ErrNoRows {
            log.Errorf("cron [role] get cache role err: %s", err.Error())
        }
        return nil
    }
    return r.(*Role)
}
