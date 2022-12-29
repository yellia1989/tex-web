package sync

import (
    "fmt"
    "time"
    "bytes"
    dsql "database/sql"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cron/date"
    zzone "github.com/yellia1989/tex-web/backend/cron/zone"
    acc "github.com/yellia1989/tex-web/backend/cron/account"
)

type role struct {
    buff bytes.Buffer   // 等待执行的sql
    rows uint32 // 同步条数
    rid uint32  // 上一次保存进度
    init bool   // 是否初始化成功
}

func (t *role) name() string {
    return "role"
}

func (t *role) sync(from *dsql.DB, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error {
    if !t.init {
        var rid dsql.NullInt64
        if err := to.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='account_newrole' and zoneid=0").Scan(&rid); err != nil {
            if err != dsql.ErrNoRows {
                return fmt.Errorf("cron [sync][role] scan err: %s", err.Error())
            }
        }
        t.rid = uint32(rid.Int64)
        t.init = true
    }

    if t.buff.Len() > 0 {
        if err := t.save(to, zoneid); err != nil {
            return fmt.Errorf("cron [sync][role] save err: %s", err.Error())
        }
    }
    
    rows, err := from.QueryContext(ctx, "SELECT _rid,zoneid,roleid,time,`first` FROM account_newrole WHERE _rid > ? order by _rid limit 10000", t.rid)
    if err != nil {
        return fmt.Errorf("cron [sync][role] query err: %s", err.Error())
    }
    defer rows.Close()

    var _rid uint32
    var rzoneid uint32
    var roleid uint32
    var st string
    var first uint32
    var buff bytes.Buffer
    size := uint32(0)
    var maxt time.Time

    for rows.Next() {
        if err := rows.Scan(&_rid, &rzoneid, &roleid, &st, &first); err != nil {
            return fmt.Errorf("cron [sync][role] scan err: %s", err.Error())
        }
        t := common.ParseTimeInLocal("2006-01-02 15:04:05", st)
        if t.After(maxt) {
            maxt = t
        }

        d := date.Get(t)
        if d == nil {
            // 日期还没准备好
            return nil
        }
        account := acc.Get(roleid)
        if account == nil {
            //if isAccountMissed(t) {
                log.Errorf("cron [sync][role] can't find account, accountid: %d", roleid)
                continue
            //}
            //return nil
        }
        z := zzone.Get(rzoneid)
        if z == nil {
            continue;
        }
        if buff.Len() > 0 {
            buff.WriteString(",")
        }
        daytime := t.Hour()*3600+t.Minute()*60+t.Second()
        buff.WriteString(fmt.Sprintf("(%d,%d,%d,%d,%d)", z.Id, account.Id, d.Id, daytime, first))
        size++
    }

    if _rid == 0 {
        return nil
    }

    if size != 0 {
        t.buff.WriteString("INSERT INTO role(zoneid_fk,accountid_fk,reg_date_fk,daytime,first) VALUES")
        t.buff.WriteString(buff.String())
        t.buff.WriteString("ON DUPLICATE KEY UPDATE reg_date_fk=VALUES(reg_date_fk), first=VALUES(first);")
        buff.Reset()
        t.rows = size
    }

    t.buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('account_newrole',0,%d)", _rid))

    if err := t.save(to, zoneid); err != nil {
        return fmt.Errorf("cron [sync][role] save err: %s", err.Error())
    }
    t.rid = _rid

    if !maxt.IsZero() {
        UpdateRoleMaxTime(maxt)
    }

    log.Debugf("cron [sync][role] rid: %d, max role time: %s", t.rid, common.FormatTimeInLocal("2006-01-02 15:04:05", maxt))

    return nil
}

func (t *role) save(to *dsql.Conn, zoneid uint32) error {
    if t.buff.Len() == 0 {
        return nil
    }

    tx, err := to.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    size := float32(t.buff.Len())/1024
    t1 := time.Now()

    var result dsql.Result
    if result, err = tx.ExecContext(ctx, t.buff.String()); err != nil {
        return fmt.Errorf("exec err: %s, sql: %s", err.Error(), t.buff.String())
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("commit err: %s, sql: %s", err.Error(), t.buff.String())
    }

    t2 := time.Now()

    rowsAffected,_ := result.RowsAffected()
    log.Debugf("cron [sync][role] cost: %.2f s, size: %.2f KB, rows: %d, affect rows: %d", t2.Sub(t1).Seconds(), size, t.rows, rowsAffected)

    t.buff.Reset()
    t.rows = 0

    return nil
}
