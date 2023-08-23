package gm

import (
	sq "database/sql"
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"strconv"
	"time"
)

type ClearLossServer struct {
	IId         int    `json:"iId"`
	IStatus     int    `json:"status"`
	IFromZoneId int    `json:"iFromZoneId"`
	IToZoneId   int    `json:"iToZoneId"`
	SDeleteTime string `json:"sDeleteTime"`
	SCreateTime string `json:"sCreateTime"`
	SStartTime  string `json:"sStartTime"`
	SEndTime    string `json:"sEndTime"`
	SResult     string `json:"sResult"`
}

func ClearLossServerList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT id,status,from_zoneid,to_zoneid,create_time,delete_time,start_time,end_time,result FROM t_clearlossuser"
	sql += " ORDER BY id desc"
	var total int
	err := db.QueryRow("SELECT count(*) as total FROM (" + sql + ") a").Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Debug(sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]ClearLossServer, 0)
	for rows.Next() {
		var r ClearLossServer
		var sResult sq.NullString
		var sCreateTime sq.NullString
		var sDeleteTime sq.NullString
		var sStartTime sq.NullString
		var sEndTime sq.NullString
		if err := rows.Scan(&r.IId, &r.IStatus, &r.IFromZoneId, &r.IToZoneId, &sCreateTime, &sDeleteTime, &sStartTime, &sEndTime, &sResult); err != nil {
			return err
		}
		r.SResult = sResult.String
		r.SCreateTime = sCreateTime.String
		r.SDeleteTime = sDeleteTime.String
		r.SStartTime = sStartTime.String
		r.SEndTime = sEndTime.String
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func ClearLossServerAdd(c echo.Context) error {
	ctx := c.(*mid.Context)
	sFromZoneId := ctx.FormValue("iFromZoneId")
	sToZoneId := ctx.FormValue("iToZoneId")
	sDeleteTime := ctx.FormValue("sDeleteTime")

	if sFromZoneId == "" || sToZoneId == "" || sDeleteTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	iDeleteTime := parseTime(sDeleteTime)

	iFromZoneId := common.Atoi(sFromZoneId)
	iToZoneId := common.Atoi(sToZoneId)
	now := time.Now().Unix()
	_, err := db.Exec("INSERT INTO t_clearlossuser(from_zoneid, to_zoneid, delete_time, create_time) VALUES(?,?,?,?)", iFromZoneId, iToZoneId, formatTime(iDeleteTime), formatTime(now))
	if err != nil {
		return err
	}

	return ctx.SendResponse("添加成功")
}

func ClearLossServerUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)
	id := ctx.FormValue("iId")

	sFromZoneId := ctx.FormValue("iFromZoneId")
	sToZoneId := ctx.FormValue("iToZoneId")
	sDeleteTime := ctx.FormValue("sDeleteTime")

	if sFromZoneId == "" || sToZoneId == "" || sDeleteTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	// 正在进行中或已经合服结束的不能修改
	sql := "SELECT id FROM t_clearlossuser WHERE or status = 3 and id = " + id
	var total int
	err := db.QueryRow("SELECT count(*) as total FROM (" + sql + ") a").Scan(&total)
	if err != nil {
		return err
	}
	if total == 0 {
		return ctx.SendError(-1, "任务已结束或正在进行中")
	}

	iDeleteTime := parseTime(sDeleteTime)

	iFromZoneId := common.Atoi(sFromZoneId)
	iToZoneId := common.Atoi(sToZoneId)
	_, err = db.Exec("UPDATE t_hefu SET from_zoneid=?,to_zoneid=?,delete_time=? WHERE id=?", iFromZoneId, iToZoneId, formatTime(iDeleteTime), id)
	if err != nil {
		return err
	}

	return ctx.SendResponse("修改成功")
}
