package chat

import (
    "context"
    dsql "database/sql"
    "fmt"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/api/game"
    "github.com/yellia1989/tex-web/backend/cfg"
    "strings"
    "time"
    "bytes"
)

type chatLog struct {
    Id     uint32 `json:"id"`
    Time   string `json:"time"`
    Type   uint32  `json:"type"`
    FromMapId uint32 `json:"fromMapId"`
    FromZoneId uint32 `json:"fromZoneId"`
    FromRoleId uint32 `json:"fromRoleId"`
    FromRoleName string `json:"fromRoleName"`
    ToZoneId uint32 `json:"toZoneId"`
    ToRoleId uint32 `json:"toRoleId"`
    ToRoleName string `json:"toRoleName"`
    AllianceId uint64 `json:"allianceId"`
    AllianceName string `json:"allianceName"`
    Content string `json:"content"`
    DirtyWord string `json:"dirtyWord"`
}

var ctx context.Context
var conn *dsql.Conn

var statConn *dsql.Conn

func init() {
    ctx = context.Background()
}

func checkConn() {
    var err error
    if conn != nil {
        err = conn.PingContext(ctx)
        if err != nil {
            conn.Close()
            conn = nil
        } else {
            return
        }
    }
    if statConn != nil {
        err = statConn.PingContext(ctx)
        if err != nil {
            statConn.Close()
            statConn = nil
        } else {
            return
        }
    }

    if conn == nil {
        conn, err = cfg.LogDb.Conn(ctx)
        if err != nil {
            panic(fmt.Sprintf("cron [chatMask] create conn err: %s", err.Error()))
        }
    }

    if statConn == nil {
        statConn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            panic(fmt.Sprintf("cron [chatMask] create statConn err: %s", err.Error()))
        }
    }
}

func Cron(now time.Time) {
    defer func() {
        if err := recover(); err != nil {
            log.Errorf("cron [chatMask] recover err: %v", err)
        }
    }()

    checkConn()

    var lastSyncRidNull dsql.NullInt64

    sql := "select sync_rid from chat_dirty_word limit 1"
    err := statConn.QueryRowContext(ctx, sql).Scan(&lastSyncRidNull)
    if err != nil && err != dsql.ErrNoRows {
        log.Errorf("cron [chatMask] query sync_rid err: %s", err.Error())
        return
    }

    lastSyncRid := uint32(lastSyncRidNull.Int64)
    oldSyncRid := lastSyncRid

    sql = "SELECT _rid,time,zoneid,mapid,type,fromroleid,fromrolename,tozoneid,toroleid,torolename,allianceid,alliancename,content FROM `chat` where _rid > ? order by _rid limit 1000"
    rows, err := conn.QueryContext(ctx, sql, lastSyncRid)
    if err != nil {
        log.Errorf("cron [chatMask] query chat log err: %s", err.Error())
        return
    }
    defer rows.Close()

    maskLogs := make([]chatLog, 0)
    for rows.Next() {
        var r chatLog
        if err := rows.Scan(&r.Id, &r.Time, &r.FromZoneId, &r.FromMapId, &r.Type, &r.FromRoleId, &r.FromRoleName, &r.ToZoneId, &r.ToRoleId, &r.ToRoleName, &r.AllianceId, &r.AllianceName, &r.Content); err != nil {
            log.Errorf("cron [chatMask] scan chat log err: %s, rid: %d", err.Error(), r.Id)
        } else {
            oldSyncRid = r.Id

            bMask := false
            for _,v := range game.GetMaskWord() {
                if v != "" && strings.Contains(r.Content, v){
                    bMask = true
                    r.DirtyWord = v
                    break
                }
            }
            if bMask {
                maskLogs = append(maskLogs, r)
            }
        }
    }

    var buff bytes.Buffer

    if oldSyncRid != lastSyncRid {
        buff.WriteString(fmt.Sprintf("insert into chat_dirty_word (id,sync_rid) values (1,%d) on duplicate key update sync_rid=%d;", oldSyncRid, oldSyncRid))
    }

    if len(maskLogs) != 0 {
        buff.WriteString("insert into chat_dirty_history(_rid,time,zoneid,mapid,type,fromroleid,fromrolename,tozoneid,toroleid,torolename,allianceid,alliancename,content,dirtyword) VALUES ")
        for i, v := range maskLogs {
            if i != 0 {
                buff.WriteString(",")
            }
            buff.WriteString(fmt.Sprintf("(%d,'%s',%d,%d,%d,%d,'%s',%d,%d,'%s',%d,'%s','%s','%s')", v.Id, v.Time, v.FromZoneId, v.FromMapId, v.Type, v.FromRoleId, v.FromRoleName, v.ToZoneId, v.ToRoleId, v.ToRoleName, v.AllianceId, v.AllianceName, v.Content, v.DirtyWord))
        }
    }

    if buff.Len() == 0 {
        return
    }

    tx, err := statConn.BeginTx(ctx,nil)
    if err != nil {
        log.Errorf("cron [chatMask] begin tx err: %s", err.Error())
        return
    }

    if _,err := tx.Exec(buff.String()); err != nil {
        if err != nil {
            log.Errorf("cron [chatMask] exec tx err: %s, sql: %s", err.Error(), buff.String())
            if err := tx.Rollback(); err != nil {
                log.Errorf("cron [chatMask] rollback tx err: %s", err.Error())
            }
            return
        }
    }
    if err := tx.Commit(); err != nil {
        log.Errorf("cron [chatMask] commit tx err: %s", err.Error())
    }
}
