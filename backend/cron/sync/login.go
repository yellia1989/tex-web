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

func (l *login) sync(from *dsql.Conn, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error {
    if err := to.PingContext(ctx); err != nil {
        return fmt.Errorf("sync login ping err: %s", err.Error())
    }

    if !l.init {
        var rid dsql.NullInt64
        if err := to.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='login' and zoneid=?", zoneid).Scan(&rid); err != nil {
            if err != dsql.ErrNoRows {
                return fmt.Errorf("sync login scan err: %s", err.Error())
            }
        }
        l.rid = uint32(rid.Int64)
        l.init = true
    }

    if l.buff.Len() > 0 {
        if err := l.save(to, zoneid); err != nil {
            return fmt.Errorf("sync login save err: %s", err.Error())
        }
    }

    if err := from.PingContext(ctx); err != nil {
        return fmt.Errorf("sync login ping err: %s", err.Error())
    }
    
    rows, err := from.QueryContext(ctx, "SELECT _rid,roleid,time,usercreatetime FROM login WHERE _rid > ? order by _rid limit 10000", l.rid)
    if err != nil {
        return fmt.Errorf("sync login query err: %s", err.Error())
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
            return fmt.Errorf("sync login scan err: %s", err.Error())
        }
        t := common.ParseTimeInLocal("2006-01-02 15:04:05", st)
        regt := common.ParseTimeInLocal("2006-01-02 15:04:05", regst)
        d := date.Get(t)
        if d == nil {
            // 日期还没准备好
            log.Errorf("can't find date: %s", st)
            return nil
        }
        account := acc.Get(roleid)
        if account == nil {
            if isAccountMissed(regt) {
                // 日志丢失了
                log.Errorf("account create log missed, accountid: %d, time: %s", roleid, regst)
                continue
            }
            // 账号还没准备好
            log.Errorf("account not ready, accountid: %d, time: %s", roleid, regst)
            return nil
        }
        r := rrole.Get(zoneidFk, account.Id)
        if r == nil {
            if isRoleMissed(regt) {
                // 日志丢失了
                log.Errorf("role create log missed, zoneid: %d, roleid: %d, time: %s", zoneid, roleid, st)
                continue
            }
            // 角色信息还没准备好
            log.Errorf("role not ready, zoneid: %d, roleid: %d, time: %s", zoneid, roleid, st)
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

    if _rid != 0 {
        l.buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('login',%d,%d);",zoneid, _rid))
    }

    if size > 0 {
        l.buff.WriteString("INSERT INTO login(zoneid_fk,accountid_fk,date_fk,daytime) VALUES")
        l.buff.WriteString(buff.String())
        buff.Reset()
        l.rows = size
    }

    if l.buff.Len() == 0 {
        return nil
    }

    if err := l.save(to, zoneid); err != nil {
        return fmt.Errorf("sync login save err: %s", err.Error())
    }

    l.rid = _rid
    log.Debugf("sync login rid: %d, zoneid: %d", l.rid, zoneid)

    return nil
}

func (l *login) save(to *dsql.Conn, zoneid uint32) error {
    tx, err := to.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    size := float32(l.buff.Len())/1024
    t1 := time.Now()

    var result dsql.Result
    if result, err = tx.ExecContext(ctx, l.buff.String()); err != nil {
        return fmt.Errorf("sync login sql: %s, err: %s", l.buff.String(), err.Error())
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("sync login sql: %s, err: %s", l.buff.String(), err.Error())
    }

    t2 := time.Now()

    rowsAffected,_ := result.RowsAffected()
    log.Debugf("sync login cost: %.2f ms, size: %.2f KB, rows: %d, affect rows: %d, zoneid: %d", t2.Sub(t1).Seconds(), size, l.rows, rowsAffected, zoneid)

    l.buff.Reset()
    l.rows = 0

    return nil
}
