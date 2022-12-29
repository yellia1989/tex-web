package stat

import (
    "fmt"
    "bytes"
    "time"
    "context"
    dsql "database/sql"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
)

var loginRid uint32
var rechargeRid uint32
var ridinit bool
var ctx context.Context
var conn2 *dsql.Conn
var buff bytes.Buffer
var row_size int

func init() {
    ctx = context.Background()
}

func checkConn2() (err error) {
    if conn2 != nil {
        err = conn2.PingContext(ctx)
        if err != nil {
            conn2.Close()
            conn2 = nil
        } else {
            return
        }
    }

    if conn2 == nil {
        conn2, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            return
        }
    }
    return
}

func Cron(now time.Time) {
    if err := checkConn2(); err != nil {
        log.Errorf("cron [stat] check conn err: %s", err.Error())
        return
    }

    // 初始化进度
    if !ridinit {
        var rid dsql.NullInt64
        if err := conn2.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='role_login' and zoneid=0").Scan(&rid); err != nil {
            if err != dsql.ErrNoRows {
                log.Errorf("cron [stat] role_login scan err: %s", err.Error())
                return
            }
        }
        loginRid = uint32(rid.Int64)

        var rid2 dsql.NullInt64
        if err := conn2.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='role_recharge' and zoneid=0").Scan(&rid2); err != nil {
            if err != dsql.ErrNoRows {
                log.Errorf("cron [stat] role_recharge scan err: %s", err.Error())
                return
            }
        }
        rechargeRid = uint32(rid2.Int64)
        ridinit = true
    }

    if buff.Len() > 0 {
        //上次的进度没有同步完
        if err := save(); err != nil {
            log.Errorf("cron [stat] save err: %s", err.Error())
            return
        }
    }

    // 开始同步
    syncRoles := make(map[zoneAccountKey]*role)
    var zoneid_fk uint32
    var accountid_fk uint32
    var date_fk uint32
    var tmp_loginrid uint32
    rows, err := conn2.QueryContext(ctx, "SELECT rid,zoneid_fk,accountid_fk,date_fk FROM login WHERE rid>? order by rid limit 30000", loginRid)
    if err != nil {
        log.Errorf("cron [stat] login query err: %s", err.Error())
        return
    }
    defer rows.Close()
    for rows.Next() {
        if err := rows.Scan(&tmp_loginrid,&zoneid_fk,&accountid_fk,&date_fk); err != nil {
            log.Errorf("cron [stat] login scan err: %s", err.Error())
            return
        }

        r := get(zoneid_fk,accountid_fk)
        if r != nil {
            r.login(date_fk)
            syncRoles[zoneAccountKey{zoneid_fk,accountid_fk}] = r
        }
    }
    rows.Close()

    var tmp_rechargerid uint32
    var money uint32
    rows2, err := conn2.QueryContext(ctx, "SELECT rid,zoneid_fk,accountid_fk,date_fk,money FROM recharge WHERE rid>? order by rid limit 10000", rechargeRid)
    if err != nil {
        log.Errorf("cron [stat] recharge query err: %s", err.Error())
        return
    }
    defer rows2.Close()
    for rows2.Next() {
        if err := rows2.Scan(&tmp_rechargerid,&zoneid_fk,&accountid_fk,&date_fk, &money); err != nil {
            log.Errorf("cron [stat] recharge scan err: %s", err.Error())
            return
        }

        r := get(zoneid_fk,accountid_fk)
        if r != nil {
            r.rge(date_fk,money)
            syncRoles[zoneAccountKey{zoneid_fk,accountid_fk}] = r
        }
    }
    rows2.Close()

    row_size = len(syncRoles)
    if row_size == 0 {
        return
    }

    for _, r := range syncRoles {
        r.str(&buff)
    }
    if tmp_loginrid != 0 {
        buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('role_login',0,%d);", tmp_loginrid))
    }

    if tmp_rechargerid != 0 {
        buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('role_recharge',0,%d);", tmp_rechargerid))
    }
    if err := save(); err != nil {
        log.Errorf("cron [stat] save err: %s", err.Error())
        return
    }

    if tmp_loginrid != 0 {
        loginRid = tmp_loginrid
    }
    if tmp_rechargerid != 0 {
        rechargeRid = tmp_rechargerid
    }
    log.Debugf("cron [stat] loginrid:%d, rechargerid:%d", loginRid, rechargeRid)
}

func save() (err error) {
    tx, err2 := conn2.BeginTx(ctx, nil)
    if err2 != nil {
        err = fmt.Errorf("begintx err: %s", err2.Error())
        return
    }

    defer tx.Rollback()

    size := float32(buff.Len())/1024
    t1 := time.Now()

    var result dsql.Result
    if result, err2 = tx.ExecContext(ctx, buff.String()); err2 != nil {
        err = fmt.Errorf("exec err: %s, sql: %s", err2.Error(), buff.String())
        return
    }

    if err2 := tx.Commit(); err2 != nil {
        err = fmt.Errorf("commit err: %s, sql: %s", err2.Error(), buff.String())
        return
    }

    t2 := time.Now()

    rowsAffected,_ := result.RowsAffected()
    log.Debugf("cron [stat] save cost: %.2f ms, size: %.2f KB, rows: %d, affect rows: %d", t2.Sub(t1).Seconds(), size, row_size, rowsAffected)

    buff.Reset()
    row_size = 0

    return
}
