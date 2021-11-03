package server

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type serverData struct {
	App                 string `json:"app"`
	Server              string `json:"server"`
	Division            string `json:"division"`
	Node                string `json:"node"`
	SettingStat         int    `json:"setting_stat"`
	CurStat             int    `json:"cur_stat"`
	ProfileConfTemplate string `json:"profile_conf_template"`
	TemplateName        string `json:"template_name"`
	Pid                 int    `json:"pid"`
}

func ServerList(c echo.Context) error {
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

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_tex")
	if err != nil {
		return err
	}

	sql := "SELECT app, server, division, node, setting_stat, cur_stat, profile_conf_template, template_name, pid FROM t_server"
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

	logs := make([]serverData, 0)
	for rows.Next() {
		var r serverData
		if err := rows.Scan(&r.App, &r.Server, &r.Division, &r.Node, &r.SettingStat, &r.CurStat, &r.ProfileConfTemplate, &r.TemplateName, &r.Pid); err != nil {
			return err
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

func ServerStart(ctx echo.Context) error {
	return nil
}

func ServerStop(ctx echo.Context) error {
	return nil
}

func ServerRestart(ctx echo.Context) error {
	return nil
}
