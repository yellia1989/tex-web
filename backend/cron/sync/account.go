package sync

import (
    "fmt"
    "time"
    "bytes"
    dsql "database/sql"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cron/date"
)

type account struct {
    buff bytes.Buffer   // 等待执行的sql
    rows uint32 // 同步条数
    rid uint32  // 上一次保存进度
    init bool   // 是否初始化成功
}

func (a *account) name() string {
    return "account"
}

func (a *account) sync(from *dsql.DB, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error {
    if !a.init {
        var rid dsql.NullInt64
        if err := to.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='account' and zoneid=?", zoneid).Scan(&rid); err != nil {
            if err != dsql.ErrNoRows {
                return fmt.Errorf("cron [sync][account] scan err: %s", err.Error())
            }
        }
        a.rid = uint32(rid.Int64)
        a.init = true
    }

    if a.buff.Len() > 0 {
        if err := a.save(to, zoneid); err != nil {
            return fmt.Errorf("cron [sync][account] save err: %s", err.Error())
        }
    }
    
    rows, err := from.QueryContext(ctx, "SELECT _rid,time,accountid,ip,ostype FROM account_create WHERE _rid > ? limit 10000", a.rid)
    if err != nil {
        return fmt.Errorf("cron [sync][account] account_create query err: %s", err.Error())
    }
    defer rows.Close()

    var _rid uint32
    var st string
    var accountid uint32
    var ip string
    var ostype uint32
    var buff bytes.Buffer
    size := uint32(0)
    var maxt time.Time

    for rows.Next() {
        if err := rows.Scan(&_rid, &st, &accountid, &ip, &ostype); err != nil {
            return fmt.Errorf("cron [sync][account] scan err: %s", err.Error())
        }
        t := common.ParseTimeInLocal("2006-01-02 15:04:05", st)
        d := date.Get(t)
        if d == nil {
            // 日期还没准备好
            return nil
        }
        if buff.Len() > 0 {
            buff.WriteString(",")
        }
        daytime := t.Hour()*3600+t.Minute()*60+t.Second()
        buff.WriteString(fmt.Sprintf("(%d,%d,'','%s',%d,%d,'')", accountid, ostype, ip, d.Id, daytime))
        size++

        if t.After(maxt) {
            maxt = t
        }
    }

    if !maxt.IsZero() {
        UpdateAccountMaxTime(maxt)
    }

    if buff.Len() > 0 {
        buff.WriteString("ON DUPLICATE KEY UPDATE ostype=VALUES(ostype), channel=VALUES(channel), ip=VALUES(ip), date_fk=VALUES(date_fk), daytime=VALUES(daytime), lang=VALUES(lang);")
        buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('account',0,%d)", _rid))
        a.buff.WriteString("INSERT INTO account(accountid,ostype,channel,ip,date_fk,daytime,lang) VALUES")
        a.buff.WriteString(buff.String())
        buff.Reset()
        a.rows = size

        if err := a.save(to, zoneid); err != nil {
            return fmt.Errorf("cron [sync][account] save err: %s", err.Error())
        }
        a.rid = _rid
    }

    return nil
}

func (a *account) save(to *dsql.Conn, zoneid uint32) error {
    tx, err := to.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    size := float32(a.buff.Len())/1024
    t1 := time.Now()

    var result dsql.Result
    if result, err = tx.ExecContext(ctx, a.buff.String()); err != nil {
        return fmt.Errorf("exec err: %s, sql: %s", err.Error(), a.buff.String())
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("commit err: %s, sql: %s", err.Error(), a.buff.String())
    }

    t2 := time.Now()

    rowsAffected,_ := result.RowsAffected()
    log.Debugf("cron [sync][account] cost: %.2f s, size: %.2f KB, rows: %d, affected rows: %d", t2.Sub(t1).Seconds(), size, a.rows, rowsAffected)

    a.buff.Reset()
    a.rows = 0

    return nil
}
