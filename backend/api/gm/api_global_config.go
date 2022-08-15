package gm

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type globalConfig struct {
	SKey   string `json:"sKey"`
	SValue string `json:"sValue"`
}

func GlobalConfigList(c echo.Context) error {
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

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_zone_global")
	if err != nil {
		return err
	}

	sql := "SELECT skey,value FROM t_config"
	var total int
	err = tx.QueryRow("SELECT count(*) as total FROM (" + sql + ") a").Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	vRet := make([]globalConfig, 0)
	for rows.Next() {
		var r globalConfig
		if err := rows.Scan(&r.SKey, &r.SValue); err != nil {
			return err
		}
		vRet = append(vRet, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendArray(vRet, total)
}

func GlobalConfigAdd(c echo.Context) error {
	ctx := c.(*mid.Context)
	stConfig := globalConfig{}
	stConfig.SKey = ctx.FormValue("sKey")
	stConfig.SValue = ctx.FormValue("sValue")

	if stConfig.SKey == "" || stConfig.SValue == "" {
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

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_zone_global")
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO t_config(skey,value) VALUES(?,?)", stConfig.SKey, stConfig.SValue)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("添加配置成功")
}

func GlobalConfigDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	keys := ctx.FormValue("keys")
	if keys == "" {
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

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_zone_global")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM t_config WHERE skey IN (?)", keys)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("删除配置成功")
}

func GlobalConfigUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)
	stConfig := globalConfig{}
	stConfig.SKey = ctx.FormValue("sKey")
	stConfig.SValue = ctx.FormValue("sValue")

	if stConfig.SKey == "" || stConfig.SValue == "" {
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

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_zone_global")
	if err != nil {
		return err
	}

	sql :="UPDATE t_config SET value = ? WHERE skey = ?"
	c.Logger().Debug(sql)
	_, err = tx.Exec(sql, stConfig.SValue, stConfig.SKey)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("更新配置成功")
}
