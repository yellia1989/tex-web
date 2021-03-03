package game

import (
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"strconv"
	"strings"
	"sync"
)

var MaskMutex *sync.RWMutex

var isMaskInit = false

var MaskWord = make([]string,0)

type chatLog struct {
	Id     uint32 `json:"id"`
	Time   string `json:"time"`
	ZoneId uint32 `json:"zoneId"`
	RoleId uint32 `json:"roleId"`
	Content string `json:"content"`
	Type   uint32  `json:"type"`
	DirtyWord string  `json:"dirtyWord"`
}

func GetMaskWord() []string  {
	if !isMaskInit{
		db := cfg.StatDb
		if db == nil {
			log.Errorf("logDb is Null")
			goto Error
		}
		var wordStr string
		err := db.QueryRow("select word from dirty_word limit 1").Scan(&wordStr)
		if err != nil {
			log.Errorf("getDirtyWord fail:",err.Error())
			goto Error
		}
		isMaskInit = true
		defer MaskMutex.Unlock()
		defer MaskMutex.Lock()
		MaskWord = strings.Split(wordStr,";")
	}
	Error:
	return MaskWord
}

func ChatGetNewest(c echo.Context) error {
	ctx := c.(*mid.Context)
	db := cfg.LogDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	var limit  = 30
	sql:= "select _rid from chat order by _rid desc limit 1"
	var maxId  = 0
	err:=db.QueryRow(sql).Scan(&maxId)
	if err != nil {
		return err
	}
	limitId := common.MaxInt(maxId - limit,0)
	sql = "select _rid,time,zoneid,roleid,content,type from chat where _rid > ? limit ?"
	rows, err := db.Query(sql,limitId,limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]chatLog, 0)
	for rows.Next() {
		var r chatLog
		if err := rows.Scan(&r.Id, &r.Time, &r.ZoneId, &r.RoleId, &r.Content,&r.Type); err != nil {
			return err
		}
		logs = append(logs, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, limit)
}

func ChatGetHistory(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	roleid := ctx.QueryParam("roleid")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")
	if zoneid == "" || roleid == "" || startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}
	db := cfg.LogDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	log.Infof("sql: %s", roleid)
	log.Infof("sql: %s", startTime)
	log.Infof("sql: %s", endTime)
	sqlcount := "SELECT count(_rid) as total FROM chat WHERE roleid= ? AND time between ? AND ?"
	var total int
	err := db.QueryRow(sqlcount,roleid,startTime,endTime).Scan(&total)
	if err != nil {
		return err
	}
	log.Infof("total: %s", total)
	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "select _rid,time,zoneid,roleid,content,type from chat"
	sql += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	sql += " ORDER BY time desc, _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]chatLog, 0)
	for rows.Next() {
		var r chatLog
		if err := rows.Scan(&r.Id, &r.Time, &r.ZoneId, &r.RoleId, &r.Content, &r.Type); err != nil {
			return err
		}
		logs = append(logs, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func ChatGetMaskLogs(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	db := cfg.StatDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	sql := "SELECT count(id) FROM `chat_dirty_word`"
	count:=0
	err:=db.QueryRow(sql).Scan(&count)
	if err != nil {
		return err
	}
	sql = "select id,time,zoneid,roleid,content,dirtyword from chat_dirty_word ORDER BY id desc limit ?,?"
	rows, err := db.Query(sql,(page-1)*limit,limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]chatLog, 0)
	for rows.Next() {
		var r chatLog
		if err := rows.Scan(&r.Id, &r.Time, &r.ZoneId, &r.RoleId, &r.Content,&r.DirtyWord); err != nil {
			return err
		}
		logs = append(logs, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, count)
}

func ChatGetMaskWord(c echo.Context) error {
	ctx := c.(*mid.Context)
	maskWord := GetMaskWord()
	return ctx.SendResponse(strings.Join(maskWord,"\n"))
}

func ChatSetMaskWord(c echo.Context) error {
	ctx := c.(*mid.Context)
	wordStr := ctx.FormValue("input")
	stringArr:= strings.Split(wordStr,"\n")
	db := cfg.StatDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	tempStr := strings.Join(stringArr,";")
	_, err := db.Exec("insert into dirty_word (id,word) values (1,?) on duplicate key update word = ?",tempStr,tempStr)
	if err != nil {
		return err
	}
	if isMaskInit{
		defer MaskMutex.Unlock()
		defer MaskMutex.Lock()
		MaskWord = stringArr
	}
	return ctx.SendResponse("设置屏蔽字成功")
}

