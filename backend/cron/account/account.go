package account

import (
    "sync"
    "errors"
    "context"
    dsql "database/sql"
    "github.com/bluele/gcache"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
)

var ctx context.Context
var conn *dsql.Conn
var connError = errors.New("account连接有准备好")
var accounts gcache.Cache
var mu sync.Mutex

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
                var err error
                conn, err = cfg.StatDb.Conn(ctx)
                if err != nil {
                    return nil, err
                }
            }

            if err := conn.PingContext(ctx); err != nil {
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
        if err != dsql.ErrNoRows {
            log.Errorf("get account err: %s", err.Error())
        }
        return nil
    }
    return a.(*Account)
}
