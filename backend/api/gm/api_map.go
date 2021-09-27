package gm

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/util"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
)

type _mapData struct {
    IMapId uint32 `json:"iMapId"`
    VZoneId []uint32 `json:"vZoneId"`
    DbHost  string   `json:"dbHost"`
	DbUser  string   `json:"dbUser"`
	DbPwd   string   `json:"dbPwd"`
	DbPort  string   `json:"dbPort"`
}

func MapSimpleList() []rpc.ZoneInfo {
    l := make([]rpc.ZoneInfo, 0)

	db := cfg.GameGlobalDb
	if db == nil {
        return l
	}

	tx, err := db.Begin()
	if err != nil {
        return l
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
        return l
	}

	sql := "SELECT mapid FROM t_maplist ORDER BY mapid desc"
	rows, err := tx.Query(sql)
	if err != nil {
        return l
	}
	defer rows.Close()

	for rows.Next() {
        var r rpc.ZoneInfo
		if err := rows.Scan(&r.IZoneId); err != nil {
            return l
		}
        r.SZoneName = fmt.Sprintf("地图(%d)", r.IZoneId)
		l = append(l, r)
    }

	if err := rows.Err(); err != nil {
		return l
	}

	if err := tx.Commit(); err != nil {
		return l
	}

    return l
}

func MapList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
		return err
	}

	sql := "SELECT mapid,zoneids,dbhost,dbport,dbuser,dbpwd FROM t_maplist"
    sql += " ORDER BY mapid desc"
    var total int
    err = tx.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_mapData, 0)
	for rows.Next() {
		var r _mapData
        var ids string
		if err := rows.Scan(&r.IMapId, &ids,&r.DbHost,&r.DbPort,&r.DbUser,&r.DbPwd); err != nil {
			return err
		}
        for _,v := range strings.Split(ids, ",") {
            u, _ := strconv.Atoi(v)
            r.VZoneId = append(r.VZoneId, uint32(u))
        }
		logs = append(logs, r)
    }

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
    
    return ctx.SendArray(logs, total)
}

func MapAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    mapid := ctx.FormValue("iMapId")
    zoneids := ctx.FormValue("zoneids")
    endpoint := ctx.FormValue("endpoint")
	dbHost := ctx.FormValue("dbHost")
	dbPort := ctx.FormValue("dbPort")
	dbUser := ctx.FormValue("dbUser")
	dbPwd := ctx.FormValue("dbPwd")

    if mapid == "" || zoneids == "" || endpoint == "" {
        return ctx.SendError(-1, "参数非法")
    }

    if err := registryAdd(cfg.App+".MapServer.MapServiceObj", cfg.App+".map."+mapid, endpoint); err != nil {
        return fmt.Errorf("增加MapServer.MapServiceObj失败: %s", err.Error())
    }

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
		return err
	}

    _, err = tx.Exec("INSERT INTO t_maplist(mapid,zoneids,dbhost,dbport,dbuser,dbpwd) VALUES(?,?,?,?,?,?)", mapid, zoneids,dbHost,dbPort,dbUser,dbPwd)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
    
    return ctx.SendResponse("添加地图成功")
}

func MapDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    zoneids := ctx.FormValue("idsStr")
    if zoneids == "" {
		return ctx.SendError(-1, "参数非法")
    }

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
		return err
	}

    _, err = tx.Exec("DELETE FROM t_maplist WHERE mapid IN (?)", zoneids)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("删除地图成功")
}

func MapEdit(c echo.Context) error {
    ctx := c.(*mid.Context)
    mapid := ctx.FormValue("iMapId")
    zoneids := ctx.FormValue("zoneids")
	dbHost := ctx.FormValue("dbHost")
	dbPort := ctx.FormValue("dbPort")
	dbUser := ctx.FormValue("dbUser")
	dbPwd := ctx.FormValue("dbPwd")

    if mapid == "" || zoneids == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
		return err
	}

    _, err = tx.Exec("UPDATE t_maplist SET zoneids=?,dbhost=?,dbport=?,dbuser=?,dbpwd=? WHERE mapid=?", zoneids,dbHost,dbPort,dbUser,dbPwd,mapid)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("修改地图成功")
}

func GameDb(zoneid uint32) (error, string) {
	db := cfg.GameGlobalDb
	if db == nil {
		return fmt.Errorf("连接数据库失败"),""
	}

	tx, err := db.Begin()
	if err != nil {
		return err,""
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
		return err,""
	}

	sql := "SELECT mapid,zoneids,dbhost,dbport,dbuser,dbpwd FROM t_maplist"
    rows, err := tx.Query(sql)
    if err != nil {
        return err,""
    }
    defer rows.Close()

    var mapid uint32
    var zoneids string
    var dbHost string
    var dbPort string
    var dbUser string
    var dbPwd string
    for rows.Next() {
        if err := rows.Scan(&mapid, &zoneids,&dbHost,&dbPort,&dbUser,&dbPwd); err != nil {
            return err,""
        }
        ids := common.Atou32v(zoneids, ",")
        if ids != nil && util.Contain(ids, zoneid) {
            break
        }
    }

    if err := rows.Err(); err != nil {
        return err,""
    }

    conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/db_map_%d", dbUser,dbPwd,dbHost,dbPort,mapid)

    return nil,conn
}
