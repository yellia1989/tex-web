package gm

import (
    "time"
    "strings"
    "strconv"
    sq "database/sql"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/common"
)

type AutoHefuServer struct {
    IId int `json:"iId"`
    IStatus int `json:"status"`
    SSrcZoneIds string `json:"sSrcZoneIds"`
    IMergeToZoneId int `json:"iMergeToZoneId"`
    SHefuTime string `json:"sHefuTime"`
    SPrepareTime string `json:"sPrepareTime"`
    SCreateTime string `json:"sCreateTime"`
    SStartTime string `json:"sStartTime"`
    SEndTime string `json:"sEndTime"`
    SResult string `json:"sResult"`
}

func formatTime(t int64) string {
    return common.FormatTimeInLocal("2006-01-02 15:04:05", time.Unix(t, 0))
}

func parseTime(v string) int64 {
    return common.ParseTimeInLocal("2006-01-02 15:04:05", v).Unix()
}

func AutoHefuServerList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT id,status,src_zoneids,to_zoneid,create_time,prepare_time,hefu_time,start_time,end_time,result FROM t_hefu"
    sql += " ORDER BY id desc"
    var total int
    err := db.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

	c.Logger().Debug(sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]AutoHefuServer, 0)
	for rows.Next() {
		var r AutoHefuServer
        var sSrcZoneIds sq.NullString
        var sResult sq.NullString
        var sCreateTime sq.NullString
        var sHefuTime sq.NullString
        var sPrepareTime sq.NullString
        var sStartTime sq.NullString
        var sEndTime sq.NullString
		if err := rows.Scan(&r.IId, &r.IStatus, &sSrcZoneIds, &r.IMergeToZoneId, &sCreateTime, &sPrepareTime, &sHefuTime, &sStartTime, &sEndTime, &sResult); err != nil {
			return err
		}
        r.SSrcZoneIds = sSrcZoneIds.String
        r.SResult = sResult.String
        r.SCreateTime = sCreateTime.String
        r.SPrepareTime = sPrepareTime.String
        r.SHefuTime = sHefuTime.String
        r.SStartTime = sStartTime.String
        r.SEndTime = sEndTime.String
		logs = append(logs, r)
    }

	if err := rows.Err(); err != nil {
		return err
	}

    return ctx.SendArray(logs, total)
}

func AutoHefuServerAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    sSrcZoneIds := ctx.FormValue("sSrcZoneIds")
    sMergeToZoneId := ctx.FormValue("iToZoneId")
    sHefuTime := ctx.FormValue("sHefuTime")
    sPrepareTime := ctx.FormValue("sPrepareTime")

    if sSrcZoneIds == "" || sMergeToZoneId == "" ||  sHefuTime == "" || sPrepareTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    iPrepareTime := parseTime(sPrepareTime)
    iHefuTime := parseTime(sHefuTime)
    if iHefuTime < iPrepareTime {
        return ctx.SendError(-1, "合服时间不能小于展示时间")
    }

    iMergeToZoneId := common.Atoi(sMergeToZoneId)
    now := time.Now().Unix()
    _, err := db.Exec("INSERT INTO t_hefu(src_zoneids, to_zoneid, prepare_time, hefu_time, create_time) VALUES(?,?,?,?,?)", sSrcZoneIds, iMergeToZoneId, formatTime(iPrepareTime), formatTime(iHefuTime), formatTime(now))
	if err != nil {
		return err
	}

    return ctx.SendResponse("添加成功")
}

func AutoHefuServerDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    zoneid := ctx.FormValue("idsStr")
    if zoneid == "" {
		return ctx.SendError(-1, "参数非法")
    }

    zoneids := strings.Split(zoneid, ",")

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    for _, zone := range zoneids {
        _, err := db.Exec("DELETE FROM t_hefu WHERE id IN (?)", zone)
	    if err != nil {
	  	    return err
	    }
    }

    return ctx.SendResponse("删除成功")
}

func AutoHefuServerUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    id := ctx.FormValue("iId")

    sSrcZoneIds := ctx.FormValue("sSrcZoneIds")
    sMergeToZoneId := ctx.FormValue("iToZoneId")
    sHefuTime := ctx.FormValue("sHefuTime")
    sPrepareTime := ctx.FormValue("sPrepareTime")

    if sSrcZoneIds == "" || sMergeToZoneId == "" || sHefuTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    // 正在进行中或已经合服结束的不能修改
	sql := "SELECT id FROM t_hefu WHERE (status = 0 or status = 3) and id = " + id
    var total int
    err := db.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }
    if (total == 0) {
        return ctx.SendError(-1, "任务已结束或正在进行中")
    }

    iPrepareTime := parseTime(sPrepareTime)
    iHefuTime := parseTime(sHefuTime)
    if iHefuTime < iPrepareTime {
        return ctx.SendError(-1, "合服时间不能小于展示时间")
    }

    iMergeToZoneId := common.Atoi(sMergeToZoneId)
    _, err = db.Exec("UPDATE t_hefu SET src_zoneids=?,to_zoneid=?,prepare_time=?,hefu_time=? WHERE id=?", sSrcZoneIds, iMergeToZoneId, formatTime(iPrepareTime), formatTime(iHefuTime), id)
	if err != nil {
		return err
	}

    return ctx.SendResponse("修改成功")
}
