package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	dsql "database/sql"

	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
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

	app := strings.TrimSpace(ctx.QueryParam("app"))
	server := strings.TrimSpace(ctx.QueryParam("server"))

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
	where := ""
	if app != "" {
		where += "app = '" + app + "'"
	}
	if server != "" {
		if where == "" {
			where += "server = '" + server + "'"
		} else {
			where += " AND server = '" + server + "'"
		}
	}
	if where != "" {
		sql += " WHERE " + where
	}
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
		var profile dsql.NullString
		if err := rows.Scan(&r.App, &r.Server, &r.Division, &r.Node, &r.SettingStat, &r.CurStat, &profile, &r.TemplateName, &r.Pid); err != nil {
			return err
		}
		if profile.Valid {
			r.ProfileConfTemplate = profile.String
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

func ServerOperator(c echo.Context) error {
	ctx := c.(*mid.Context)

	app := ctx.FormValue("app")
	server := ctx.FormValue("server")
	division := ctx.FormValue("division")
	cmd := ctx.FormValue("cmd")

	if app == "" || server == "" || cmd == "" {
		return ctx.SendError(-1, "参数为空")
	}

	req := rpc.PatchTaskReq{}
	req.STaskNo = strconv.FormatInt(time.Now().UnixNano(), 10)
	reqItem := rpc.PatchTaskItemReq{}
	reqItem.STaskNo = req.STaskNo
	reqItem.SApp = app
	reqItem.SServer = server
	reqItem.SDivision = division
	reqItem.SNodeName = "192.168.0.16"
	reqItem.SCommand = cmd
	req.VItem = append(req.VItem, reqItem)

	comm := cfg.Comm
	patchPrx := new(rpc.Patch)
	comm.StringToProxy("tex.mfwpatch.PatchObj", patchPrx)

	ret, err := patchPrx.AddTask(req)
	if ret != 0 || err != nil {
		return fmt.Errorf("opt failed, ret:%d, err:%s", ret, err.Error())
	}

	return ctx.SendResponse(req.STaskNo)
}

func GetTask(c echo.Context) error {
	ctx := c.(*mid.Context)

	taskNo := ctx.FormValue("taskNo")
	if taskNo == "" {
		return ctx.SendError(-1, "参数非法")
	}

	comm := cfg.Comm
	patchPrx := new(rpc.Patch)
	comm.StringToProxy("tex.mfwpatch.PatchObj", patchPrx)

	taskRsp := &rpc.PatchTaskRsp{}
	ret, err := patchPrx.GetTask(taskNo, taskRsp)
	if ret != 0 || err != nil {
		return fmt.Errorf("opt failed, ret:%d, err:%s", ret, err.Error())
	}

	return ctx.SendResponse(taskRsp)
}
