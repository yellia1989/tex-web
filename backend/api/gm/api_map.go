package gm

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/util"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/common"
)

type _dbData struct {
    ID uint32 `json:"id"`
    VZoneId []uint32 `json:"vZoneId"`
    DbHost  string   `json:"dbHost"`
	DbUser  string   `json:"dbUser"`
	DbPwd   string   `json:"dbPwd"`
	DbPort  string   `json:"dbPort"`
}

func DbList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT id,zoneids,dbhost,dbport,dbuser,dbpwd FROM t_dblist"
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

	logs := make([]_dbData, 0)
	for rows.Next() {
		var r _dbData
        var ids string
		if err := rows.Scan(&r.ID, &ids,&r.DbHost,&r.DbPort,&r.DbUser,&r.DbPwd); err != nil {
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

    return ctx.SendArray(logs, total)
}

func DbAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneids := ctx.FormValue("zoneids")
	dbHost := ctx.FormValue("dbHost")
	dbPort := ctx.FormValue("dbPort")
	dbUser := ctx.FormValue("dbUser")
	dbPwd := ctx.FormValue("dbPwd")

    if zoneids == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("INSERT INTO t_dblist(zoneids,dbhost,dbport,dbuser,dbpwd) VALUES(?,?,?,?,?)", zoneids,dbHost,dbPort,dbUser,dbPwd)
	if err != nil {
		return err
	}
    
    return ctx.SendResponse("添加db成功")
}

func DbDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := ctx.FormValue("idsStr")
    if ids == "" {
		return ctx.SendError(-1, "参数非法")
    }

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("DELETE FROM t_dblist WHERE id IN (?)", ids)
	if err != nil {
		return err
	}

    return ctx.SendResponse("删除db成功")
}

func DbEdit(c echo.Context) error {
    ctx := c.(*mid.Context)
    id := ctx.FormValue("id")
    zoneids := ctx.FormValue("zoneids")
	dbHost := ctx.FormValue("dbHost")
	dbPort := ctx.FormValue("dbPort")
	dbUser := ctx.FormValue("dbUser")
	dbPwd := ctx.FormValue("dbPwd")

    if id == "" || zoneids == "" {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    _, err := db.Exec("UPDATE t_dblist SET zoneids=?,dbhost=?,dbport=?,dbuser=?,dbpwd=? WHERE id=?", zoneids,dbHost,dbPort,dbUser,dbPwd,id)
	if err != nil {
		return err
	}

    return ctx.SendResponse("修改db成功")
}

func GameDb(zoneid uint32) (error, string) {
	db := cfg.GameGlobalDb
	if db == nil {
		return fmt.Errorf("连接数据库失败"),""
	}

	sql := "SELECT id,zoneids,dbhost,dbport,dbuser,dbpwd FROM t_dblist"
    rows, err := db.Query(sql)
    if err != nil {
        return err,""
    }
    defer rows.Close()

    var id uint32
    var zoneids string
    var dbHost string
    var dbPort string
    var dbUser string
    var dbPwd string
    for rows.Next() {
        if err := rows.Scan(&id, &zoneids,&dbHost,&dbPort,&dbUser,&dbPwd); err != nil {
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

    conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/db_zone_%d", dbUser,dbPwd,dbHost,dbPort,zoneid)

    return nil,conn
}
