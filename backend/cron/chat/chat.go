package chat

import (
	"context"
	dsql "database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-web/backend/api/game"
	"github.com/yellia1989/tex-web/backend/cfg"
	"strings"
	"time"
)

type chatLog struct {
	Id     uint32 `json:"id"`
	Time   string `json:"time"`
	ZoneId uint32 `json:"zoneId"`
	RoleId uint32 `json:"roleId"`
	Content string `json:"content"`
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
			log.Errorf("%v", err)
		}
	}()

	checkConn()

	var lastCheckId uint32 = 0

	sql := "select lastCheckId from dirty_word"
	err := statConn.QueryRowContext(ctx, sql).Scan(&lastCheckId)
	if err != nil {
		log.Errorf("cron [chatMaskCheckId] query err: %s", err.Error())
		return
	}

	oldCheckId := lastCheckId

	sql = "SELECT _rid,time,zoneid,roleid,content FROM `chat` where _rid > ? limit 1000"
	rows, err := conn.QueryContext(ctx, sql, lastCheckId)
	if err != nil {
		log.Errorf("cron [chatMask] query err: %s", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var clog chatLog
		if err := rows.Scan(&clog.Id, &clog.Time, &clog.ZoneId, &clog.RoleId, &clog.Content); err != nil {
			log.Errorf("cron [chatMask] scan err: %s, logid: %d", err.Error(), clog.Id)
		} else {
			bMask := false
			firstDirty := ""
			for _,v:= range game.GetMaskWord(){
				if strings.Contains(clog.Content,v){
					bMask = true
					firstDirty = v
					break
				}
			}
			if bMask {
				sql = "insert into chat_dirty_word (zoneid,roleid,content,time,dirtyword,globalid) values(?,?,?,?,?,?)"
				if _, err := statConn.ExecContext(ctx,sql,clog.ZoneId,clog.RoleId,clog.Content,clog.Time,firstDirty,clog.Id);err!=nil{
					if driverErr, ok := err.(*mysql.MySQLError); ok {
						if driverErr.Number == 1062 {
							lastCheckId = clog.Id
							continue
						}
					}
					log.Errorf("cron [chatMask] insert log err: %s", err.Error())
					return
				}
			}
			lastCheckId = clog.Id
		}
	}
	if oldCheckId!=lastCheckId{
		sql = "update dirty_word set lastCheckId = ? where id = 1"
		if _, err := statConn.ExecContext(ctx,sql,lastCheckId);err!=nil{
			log.Errorf("cron [chatMask] update lastCheckId err: %s", err.Error())
			return
		}
	}

}
