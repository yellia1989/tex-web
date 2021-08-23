package gm

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/util"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
)

type _mapData struct {
    IMapId uint32 `json:"iMapId"`
    VZoneId []uint32 `json:"vZoneId"`
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

	sql := "SELECT mapid,zoneids FROM t_maplist"
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
		if err := rows.Scan(&r.IMapId, &ids); err != nil {
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

    _, err = tx.Exec("INSERT INTO t_maplist(mapid,zoneids) VALUES(?,?)", mapid, zoneids)
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

    _, err = tx.Exec("UPDATE t_maplist SET zoneids=? WHERE mapid=?", zoneids, mapid)
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

	sql := "SELECT mapid,zoneids FROM t_maplist"
    rows, err := tx.Query(sql)
    if err != nil {
        return err,""
    }
    defer rows.Close()

    var mapid uint32
    var zoneids string
    for rows.Next() {
        if err := rows.Scan(&mapid, &zoneids); err != nil {
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

    cfg := cfg.Config
    conn := cfg.GetCfg(fmt.Sprintf("gamedb_%d", mapid),"")
    if conn == "" {
        return fmt.Errorf("缺少数据库配置:%d", zoneid),""
    }

    return nil,conn
}
