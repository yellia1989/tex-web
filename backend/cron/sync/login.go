package sync

import (
    "fmt"
    "time"
    "bytes"
    dsql "database/sql"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cron/date"
    acc "github.com/yellia1989/tex-web/backend/cron/account"
    rrole "github.com/yellia1989/tex-web/backend/cron/role"
)

type login struct {
    buff bytes.Buffer   // 等待执行的sql
    rows uint32 // 同步条数
    rid uint32  // 上一次保存进度
    init bool   // 是否初始化成功
}

func (l *login) name() string {
    return "login"
}

func (l *login) sync(from *dsql.DB, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error {
    if !l.init {
        var rid dsql.NullInt64
        if err := to.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='login' and zoneid=?", zoneid).Scan(&rid); err != nil {
            if err != dsql.ErrNoRows {
                return fmt.Errorf("cron [sync][login] scan err: %s, zoneid: %d", err.Error(), zoneid)
            }
        }
        l.rid = uint32(rid.Int64)
        l.init = true
    }

    if l.buff.Len() > 0 {
        if err := l.save(to, zoneid); err != nil {
            return fmt.Errorf("cron [sync][login] save err: %s, zoneid: %d", err.Error(), zoneid)
        }
    }

    rows, err := from.QueryContext(ctx, "SELECT _rid,roleid,time,usercreatetime FROM login WHERE _rid > ? order by _rid limit 10000", l.rid)
    if err != nil {
        return fmt.Errorf("cron [sync][login] query err: %s, zoneid: %d", err.Error(), zoneid)
    }
    defer rows.Close()

    var _rid uint32
    var roleid uint32
    var st string
    var regst string
    var buff bytes.Buffer
    size := uint32(0)
    for rows.Next() {
        if err := rows.Scan(&_rid, &roleid, &st, &regst); err != nil {
            return fmt.Errorf("cron [sync][login] scan err: %s, zoneid: %d", err.Error(), zoneid)
        }
        t := common.ParseTimeInLocal("2006-01-02 15:04:05", st)
        regt := common.ParseTimeInLocal("2006-01-02 15:04:05", regst)
        d := date.Get(t)
        if d == nil {
            return nil
        }
        account := acc.Get(roleid)
        if account == nil {
            if isAccountMissed(regt) {
                log.Errorf("cron [sync][login] can't find account, accountid: %d", roleid)
                continue
            }
            return nil
        }
        r := rrole.Get(zoneidFk, account.Id)
        if r == nil {
            if isRoleMissed(regt) {
                log.Errorf("cron [sync][login] can't find role, roleid: %d, reg time: %s, zoneid: %d", roleid, regst, zoneid)
                continue
            }
            return nil
        }
        if d.Id < r.RegDateFk {
            // 日志不对
            continue
        }

        if buff.Len() > 0 {
            buff.WriteString(",")
        }
        daytime := t.Hour()*3600+t.Minute()*60+t.Second()
        buff.WriteString(fmt.Sprintf("(%d,%d,%d,%d)", zoneidFk, account.Id, d.Id, daytime))
        size++
    }

    if _rid == 0 {
        return nil
    }

    l.buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('login',%d,%d);",zoneid, _rid))

    if size > 0 {
        l.buff.WriteString("INSERT INTO login(zoneid_fk,accountid_fk,date_fk,daytime) VALUES")
        l.buff.WriteString(buff.String())
        buff.Reset()
        l.rows = size
    }

    if err := l.save(to, zoneid); err != nil {
        return fmt.Errorf("cron [sync][login] save err: %s, zoneid: %d", err.Error(), zoneid)
    }

    l.rid = _rid
    log.Debugf("cron [sync][login] rid: %d, zoneid: %d", l.rid, zoneid)

    return nil
}

func (l *login) save(to *dsql.Conn, zoneid uint32) error {
    if l.buff.Len() == 0 {
        return nil
    }

    tx, err := to.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    size := float32(l.buff.Len())/1024
    t1 := time.Now()

    var result dsql.Result
    if result, err = tx.ExecContext(ctx, l.buff.String()); err != nil {
        return fmt.Errorf("exec err: %s, sql: %s", err.Error(), l.buff.String())
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("commit err: %s, sql: %s", err.Error(), l.buff.String())
    }

    t2 := time.Now()

    rowsAffected,_ := result.RowsAffected()
    log.Debugf("cron [sync][login] cost: %.2f s, size: %.2f KB, rows: %d, affect rows: %d, zoneid: %d", t2.Sub(t1).Seconds(), size, l.rows, rowsAffected, zoneid)

    l.buff.Reset()
    l.rows = 0

    return nil
}
