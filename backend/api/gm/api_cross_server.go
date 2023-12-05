package gm

import (
    "strings"
    "strconv"
    sq "database/sql"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type CrossServer struct {
    IId int `json:"iId"`
    IStatus int `json:"status"`
    SZoneIds string `json:"sZoneIds"`
    SNode string `json:"sNode"`
    SDBInfo string `json:"sDBInfo"`
}

func CrossServerList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT mapid,status,zoneids,node,dbinfo FROM t_maplist"
    sql += " ORDER BY mapid desc"
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

	logs := make([]CrossServer, 0)
	for rows.Next() {
		var r CrossServer
        var sZoneIds sq.NullString
		if err := rows.Scan(&r.IId, &r.IStatus, &sZoneIds, &r.SNode, &r.SDBInfo); err != nil {
			return err
		}
        r.SZoneIds = sZoneIds.String
		logs = append(logs, r)
    }

	if err := rows.Err(); err != nil {
		return err
	}

    return ctx.SendArray(logs, total)
}

func CrossServerAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    mapId, _ := strconv.Atoi(ctx.FormValue("iId"))
    sZoneIds := ctx.FormValue("sZoneIds")
    sNode := ctx.FormValue("sNode")
	sDBInfo := ctx.FormValue("sDBInfo")

    if sZoneIds == ""  {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("INSERT INTO t_maplist(mapid, zoneids,node,dbinfo) VALUES(?,?,?,?)", mapId,sZoneIds,sNode,sDBInfo)
	if err != nil {
		return err
	}

    return ctx.SendResponse("添加成功")
}

func CrossServerDel(c echo.Context) error {
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
        _, err := db.Exec("DELETE FROM t_maplist WHERE mapid IN (?)", zone)
	    if err != nil {
	  	    return err
	    }
    }

    return ctx.SendResponse("删除成功")
}

func CrossServerUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    id := ctx.FormValue("iId")

    sZoneIds := ctx.FormValue("sZoneIds")
    sNode := ctx.FormValue("sNode")
	sDBInfo := ctx.FormValue("sDBInfo")

    if sZoneIds == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("UPDATE t_maplist SET zoneids=?,node=?,dbinfo=? WHERE mapid=?", sZoneIds,sNode,sDBInfo,id)
	if err != nil {
		return err
	}

    return ctx.SendResponse("修改成功")
}
