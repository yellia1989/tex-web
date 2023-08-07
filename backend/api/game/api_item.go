package game

import (
	sq "database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-go/tools/log"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"strconv"
)

type itemlog struct {
	Id     uint32 `json:"id"`
	Time   string `json:"time"`
	BaseId uint32 `json:"baseId"`
	AddNum uint32 `json:"add_num"`
	CurNum uint32 `json:"cur_num"`
	Action string `json:"action"`
}

func ItemAddLog(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	roleid := ctx.QueryParam("roleid")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")
	itemid := ctx.QueryParam("itemid")
	turnMerge := ctx.QueryParam("turnMergeZone")

	if zoneid == "" || roleid == "" || startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db, err := zoneLogDbWithOptions(zoneid, ZoneLogOptions{trunMerge: turnMerge == "on"})

	if err != nil {
		return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
	}
	defer db.Close()

	whereStr := " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"

	if itemid != "" {
		whereStr += " AND id=" + itemid
	}

	sqlcount := "SELECT count(*) as total FROM add_item"
	sqlcount += whereStr
	var total int
	err = db.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid,time,id,add_num,cur_num,action FROM add_item"
	sql += whereStr
	sql += " ORDER BY time desc, _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]itemlog, 0)
	for rows.Next() {
		var r itemlog
		if err := rows.Scan(&r.Id, &r.Time, &r.BaseId, &r.AddNum, &r.CurNum, &r.Action); err != nil {
			return err
		}
		logs = append(logs, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func ItemSubLog(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	roleid := ctx.QueryParam("roleid")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")
	itemid := ctx.QueryParam("itemid")
	turnMerge := ctx.QueryParam("turnMergeZone")

	if zoneid == "" || roleid == "" || startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db, err := zoneLogDbWithOptions(zoneid, ZoneLogOptions{trunMerge: turnMerge == "on"})

	whereStr := " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"

	if itemid != "" {
		whereStr += " AND id=" + itemid
	}

	if err != nil {
		return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
	}
	defer db.Close()

	sqlcount := "SELECT count(*) as total FROM sub_item"
	sqlcount += whereStr
	var total int
	err = db.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid,time,id,sub_num,cur_num,action FROM sub_item"
	sql += whereStr
	sql += " ORDER BY time desc, _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]itemlog, 0)
	for rows.Next() {
		var r itemlog
		var subNum sq.NullInt32
		if err := rows.Scan(&r.Id, &r.Time, &r.BaseId, &subNum, &r.CurNum, &r.Action); err != nil {
			return err
		}
		r.AddNum = (uint32)(subNum.Int32)
		logs = append(logs, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}
