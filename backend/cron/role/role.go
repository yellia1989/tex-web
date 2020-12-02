package stat

import (
    "sync"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-go/tools/log"
)

var roles gcache.Cache
var mu sync.Mutex
var conn *dsql.Conn
var ctx context.Context

func createNewConn() (err error) {
    if conn != nil {
        conn.Close()
    }
    conn, err = cfg.StatDb.Conn(ctx)
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

            if conn == nil {
                if err := createNewConn(); err != nil {
                    return nil, err
                }
            }

            if err := conn.PingContext(ctx); err != nil {
                if err := createNewConn(); err != nil {
                    return nil, err
                }
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
            log.Errorf("get role err: %s", err.Error())
        }
        return nil
    }
    return r.(*Role)
}
