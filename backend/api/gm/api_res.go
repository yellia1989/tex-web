package gm

import (
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type resControl struct {
	ResId      uint32   `json:"iResId"`
	Action     []string `json:"sAction"`
	ActionName []string `json:"sActionName"`
}

type Action struct {
	Vaule string `json:"vaule"`
	Name  string `json:"name"`
}

var vAction []Action
var nextUpdateTime time.Time

func ResControlList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db := cfg.GameGlobalDb
	sql := "SELECT res_id, action, action_name from t_res_control"
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		return err
	}

	vResControl := make([]resControl, 0)
	for rows.Next() {
		var r resControl
		var sAction string
		var sActionName string
		if err := rows.Scan(&r.ResId, &sAction, &sActionName); err != nil {
			return err
		}

		r.Action = strings.Split(sAction, ",")
		r.ActionName = strings.Split(sActionName, ",")
		vResControl = append(vResControl, r)
	}

	return ctx.SendArray(vResControl, len(vResControl))
}

func ActionList(c echo.Context) error {
	ctx := c.(*mid.Context)

	actionList := getAllAction()

	return ctx.SendArray(actionList, len(actionList))
}

func refreshActionList() {
	now := time.Now()
	if now.Before(nextUpdateTime) {
		return
	}

	db := cfg.LogDb
	sql := "SELECT action, action_name from user_action"
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}

	vtmp := make([]Action, 0)
	for rows.Next() {
		var r Action
		if err := rows.Scan(&r.Vaule, &r.Name); err != nil {
			return
		}
		vtmp = append(vtmp, r)
	}
	vAction = vtmp

	nextUpdateTime = now.Add(time.Minute * 5)
}

func getAllAction() []Action {
	refreshActionList()

	allAction := make([]Action, len(vAction))

	copy(allAction, vAction)

	return allAction
}

func ActionAdd(c echo.Context) error {
	ctx := c.(*mid.Context)
	iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))
	sAction := ctx.FormValue("sAction")
	sActionName := ctx.FormValue("sActionName")

	if iResId == 0 || sAction == "" || sActionName == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameGlobalDb
	sql := "INSERT INTO t_res_control (res_id, action, action_name) VALUES(?,?,?)"

	rows, err := db.Query(sql, iResId, sAction, sActionName)
	defer rows.Close()
	if err != nil {
		return err
	}

	return ctx.SendResponse("添加资源监控项成功")
}

func ActionEdit(c echo.Context) error {
	ctx := c.(*mid.Context)
	iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))
	sAction := ctx.FormValue("sAction")
	sActionName := ctx.FormValue("sActionName")

	if iResId == 0 || sAction == "" || sActionName == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameGlobalDb
	sql := "UPDATE t_res_control SET action=?, action_name=? WHERE res_id=?"

	rows, err := db.Query(sql, sAction, sActionName, iResId)
	defer rows.Close()
	if err != nil {
		return err
	}

	return ctx.SendResponse("编辑资源监控项成功")
}

func ActionDel(c echo.Context) error {
	ctx := c.(*mid.Context)
	iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))

	if iResId == 0 {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameGlobalDb
	sql := "DELETE FROM t_res_control WHERE res_id=?"

	rows, err := db.Query(sql, iResId)
	defer rows.Close()
	if err != nil {
		return err
	}
	return ctx.SendResponse("删除资源监控项成功")
}
