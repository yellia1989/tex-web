package sync

import (
    "fmt"
    "time"
    "math"
    "bytes"
    dsql "database/sql"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cron/date"
    acc "github.com/yellia1989/tex-web/backend/cron/account"
    rrole "github.com/yellia1989/tex-web/backend/cron/role"
)

type recharge struct {
    buff bytes.Buffer   // 等待执行的sql
    rows uint32 // 同步条数
    rid uint32  // 上一次保存进度
    init bool   // 是否初始化成功
}

func (t *recharge) name() string {
    return "recharge"
}

func (t *recharge) sync(from *dsql.DB, to *dsql.Conn, zoneid uint32, zoneidFk uint32) error {
    if !t.init {
        var rid dsql.NullInt64
        if err := to.QueryRowContext(ctx, "SELECT rid FROM sync_rid WHERE `table`='recharge' and zoneid=?", zoneid).Scan(&rid); err != nil {
            if err != dsql.ErrNoRows {
                return fmt.Errorf("cron [sync][recharge] scan err: %s, zoneid: %d", err.Error(), zoneid)
            }
        }
        t.rid = uint32(rid.Int64)
        t.init = true
    }

    if t.buff.Len() > 0 {
        if err := t.save(to, zoneid); err != nil {
            return fmt.Errorf("cron [sync][recharge] save err: %s, zoneid: %d", err.Error(), zoneid)
        }
    }

    rows, err := from.QueryContext(ctx, "SELECT _rid,roleid,time,usercreatetime,product_id,price,moneytotal FROM iap_recharge WHERE _rid > ? order by _rid limit 10000", t.rid)
    if err != nil {
        return fmt.Errorf("cron [sync][recharge] query err: %s, zoneid: %d", err.Error(), zoneid)
    }
    defer rows.Close()

    var _rid uint32
    var roleid uint32
    var st string
    var regst string
    var productid uint32
    var price float32
    var moneytotal float32
    var first uint32
    var buff bytes.Buffer
    size := uint32(0)
    for rows.Next() {
        if err := rows.Scan(&_rid, &roleid, &st, &regst, &productid, &price, &moneytotal); err != nil {
            return fmt.Errorf("sync recharge scan err: %s", err.Error())
        }
        if moneytotal-price < math.SmallestNonzeroFloat32 {
            first = 1
        } else {
            first = 0
        }
        t := common.ParseTimeInLocal("2006-01-02 15:04:05", st)
        //regt := common.ParseTimeInLocal("2006-01-02 15:04:05", regst)
        d := date.Get(t)
        if d == nil {
            // 日期还没准备好
            return nil
        }
        account := acc.Get(roleid)
        if account == nil {
            //if isAccountMissed(regt) {
                log.Errorf("cron [sync][recharge] can't find account, accountid: %d", roleid)
                continue
            //}
            //return nil
        }
        r := rrole.Get(zoneidFk, account.Id)
        if r == nil {
            //if isRoleMissed(regt) {
                log.Errorf("cron [sync][recharge] can't find role, roleid: %d, reg time: %s, zoneid: %d", roleid, regst, zoneid)
                continue
            //}
            //return nil
        }
        if d.Id < r.RegDateFk {
            // 日志不对
            continue
        }

        if buff.Len() > 0 {
            buff.WriteString(",")
        }
        daytime := t.Hour()*3600+t.Minute()*60+t.Second()
        buff.WriteString(fmt.Sprintf("(%d,%d,%d,%d,%d,%d,%d)", zoneidFk, account.Id, d.Id, daytime,productid,uint32(price*100),first))
        size++
    }

    if _rid == 0 {
        return nil
    }

    t.buff.WriteString(fmt.Sprintf("REPLACE INTO sync_rid(`table`,zoneid,rid) VALUES('recharge',%d,%d);",zoneid, _rid))

    if size != 0 {
        t.buff.WriteString("INSERT INTO recharge(zoneid_fk,accountid_fk,date_fk,daytime,product_id,money,first) VALUES")
        t.buff.WriteString(buff.String())
        buff.Reset()
        t.rows = size
    }

    if err := t.save(to, zoneid); err != nil {
        return fmt.Errorf("cron [sync][recharge] save err: %s, zoneid: %d", err.Error(), zoneid)
    }
    t.rid = _rid

    log.Debugf("cron [sync][recharge] rid: %d, zoneid: %d", t.rid, zoneid)

    return nil
}

func (t *recharge) save(to *dsql.Conn, zoneid uint32) error {
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
    log.Debugf("cron [sync][recharge] cost: %.2f s, size: %.2f KB, rows: %d, affect rows: %d, zoneid: %d", t2.Sub(t1).Seconds(), size, t.rows, rowsAffected, zoneid)

    t.buff.Reset()
    t.rows = 0

    return nil
}
