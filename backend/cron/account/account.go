package account

import (
    "sync"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
)

var ctx context.Context
var conn *dsql.Conn
var accounts gcache.Cache
var mu sync.Mutex

func createNewConn() (err error) {
    if conn != nil {
        conn.Close()
    }
    conn, err = cfg.StatDb.Conn(ctx)
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
        if err != dsql.ErrNoRows {
            log.Errorf("get account err: %s", err.Error())
        }
        return nil
    }
    return a.(*Account)
}
