package gm

import (
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/cfg"
    "strings"
)

type AutoOpenServer struct {
    IZoneId uint32 `json:"iZoneId"`
    SNode string `json:"sNode"`
    SDBInfo string `json:"sDBInfo"`
    SConnAddr string `json:"sConnaddr"`
    IMaxNum uint32 `json:"iMaxNum"`
    IMaxOnlineNUm uint32 `json:"iMaxOnline"`
    IOpenServerNum uint32 `json:"iOpenServerNum"`
}

func AutoOpenServerList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT zoneid,node,dbinfo,connaddr,createrole,onlinerole,openrole FROM t_kaifu"
    sql += " ORDER BY zoneid desc"
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

	logs := make([]AutoOpenServer, 0)
	for rows.Next() {
		var r AutoOpenServer
		if err := rows.Scan(&r.IZoneId, &r.SNode,&r.SDBInfo,&r.SConnAddr,&r.IMaxNum,&r.IMaxOnlineNUm,&r.IOpenServerNum); err != nil {
			return err
		}
		logs = append(logs, r)
    }

	if err := rows.Err(); err != nil {
		return err
	}

    return ctx.SendArray(logs, total)
}

func AutoOpenServerAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("iZoneId")
	sNode := ctx.FormValue("sNode")
	sDBInfo := ctx.FormValue("sDBInfo")
	sConnaddr := ctx.FormValue("sConnaddr")
	iMaxNum := ctx.FormValue("iMaxNum")
	iMaxOnline := ctx.FormValue("iMaxOnline")
	iOpenServerNum := ctx.FormValue("iOpenServerNum")

    if zoneid == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("INSERT INTO t_kaifu(zoneid,node,dbinfo,connaddr,createrole,onlinerole,openrole) VALUES(?,?,?,?,?,?,?)", zoneid,sNode,sDBInfo,sConnaddr,iMaxNum, iMaxOnline, iOpenServerNum)
	if err != nil {
		return err
	}

    return ctx.SendResponse("添加服务器成功")
}

func AutoOpenServerDel(c echo.Context) error {
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
      _, err := db.Exec("DELETE FROM t_kaifu WHERE zoneid IN (?)", zone)
	  if err != nil {
	  	return err
	  }
  }



    return ctx.SendResponse("删除服务器成功")
}

func AutoOpenServerUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("iZoneId")
	sNode := ctx.FormValue("sNode")
	sDBInfo := ctx.FormValue("sDBInfo")
	sConnaddr := ctx.FormValue("sConnaddr")
	iMaxNum := ctx.FormValue("iMaxNum")
	iMaxOnline := ctx.FormValue("iMaxOnline")
	iOpenServerNum := ctx.FormValue("iOpenServerNum")

    if zoneid == ""{
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.ServerGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("UPDATE t_kaifu SET zoneid=?,node=?,dbinfo=?,connaddr=?,createrole=?,onlinerole=?,openrole=? WHERE zoneid=?", zoneid,sNode,sDBInfo,sConnaddr,iMaxNum,iMaxOnline,iOpenServerNum,zoneid)
	if err != nil {
		return err
	}

    return ctx.SendResponse("修改服务器成功")
}
