package account

import (
    "sync"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
)

var mu sync.Mutex
var ctx context.Context
var conn *dsql.Conn

var accounts gcache.Cache

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

type Account struct {
    Id uint32
    AccountId uint32
}

func init() {
    ctx = context.Background()
    accounts = gcache.New(10000).
        LFU().
        LoaderFunc(func(key interface{}) (interface{}, error) {
            mu.Lock()
            defer mu.Unlock()

            if err := checkConn(); err != nil {
                return nil, err
            }

            a := Account{}
            if err := conn.QueryRowContext(ctx, "SELECT id,accountid FROM account WHERE accountid=?", key).Scan(&a.Id, &a.AccountId); err != nil {
                return nil, err
            }
            return &a,nil
        }).
        Build()
}

func Get(accountid uint32) *Account {
    a, err := accounts.Get(accountid)
    if a == nil {
        //if err != dsql.ErrNoRows {
            log.Errorf("cron [account] get cache account err: %s, accountid: %d", err.Error(), accountid)
        //}
        return nil
    }
    return a.(*Account)
}
