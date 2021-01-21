package gm

import (
	"regexp"
	"strings"
	"time"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func parseIDStr(src, sep string, out *string) {
	*out = ""
	reg, _ := regexp.Compile("\\d{5,}")
	vStr := reg.FindAllString(src, -1)

	if len(vStr) != 0 {
		*out = strings.Join(vStr, sep)
	}
}

func WhiteList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT account_id FROM t_whitelist WHERE del_time is NULL")
	if err != nil {
		return err
	}
	defer rows.Close()

	var vStr []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return err
		}

		vStr = append(vStr, id)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse(strings.Join(vStr, ";"))
}

func WhiteAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	if input == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}

	parseIDStr(input, "),(", &input)
	if len(input) == 0 {
		return ctx.SendError(-1,"用户ID格式不正确")
	}
	sql := "REPLACE INTO t_whitelist(account_id) VALUES(" + input + ");"
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("添加白名单用户成功")
}

func WhiteDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	if input == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}

	parseIDStr(input, ",", &input)
	if len(input) == 0 {
		return ctx.SendError(-1,"用户ID格式不正确")
	}
	sql := "DELETE FROM t_whitelist WHERE del_time is NULL AND account_id IN(" + input + ");"
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("删除白名单用户成功")
}

func WhiteReplace(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM t_whitelist WHERE del_time is NULL")
	if err != nil {
		return err
	}

	parseIDStr(input, "),(", &input)
	if len(input) == 0 {
		return ctx.SendError(-1,"用户ID格式不正确")
	}
	sql := "REPLACE INTO t_whitelist(account_id) VALUES(" + input + ");"
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("覆盖白名单用户成功")
}


func TmpWhiteList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT account_id,del_time FROM t_whitelist WHERE del_time is not NULL")
	if err != nil {
		return err
	}
	defer rows.Close()

	var vStr []string
	for rows.Next() {
		var id string
		var delst string
		err = rows.Scan(&id,&delst)
		if err != nil {
			return err
		}

		delt := common.ParseTimeInLocal("2006-01-02 15:04:05", delst)
		now := time.Now()
		if now.After(delt) {
			continue
		}
		vStr = append(vStr, id)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse(strings.Join(vStr, ";"))
}

func WhiteAddTmp(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	if input == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}

	reg, _ := regexp.Compile("\\d{5,}")
	vStr := reg.FindAllString(input, -1)
	if len(vStr) == 0 {
		return ctx.SendError(-1,"用户ID格式不正确")
	}
	now := time.Now();
	delt := now.Add(time.Duration(8) * time.Hour)
	delst := delt.Format("2006-01-02 15:04:05")

	sql := "INSERT IGNORE t_whitelist(account_id,del_time) VALUES(?,?);"
	for _, sId := range vStr {
		_, err = tx.Exec(sql,sId,delst)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("添加临时白名单用户成功")
}

func WhiteDelTmp(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	if input == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_loginserver")
	if err != nil {
		return err
	}

	parseIDStr(input, ",", &input)
	if len(input) == 0 {
		return ctx.SendError(-1,"用户ID格式不正确")
	}
	sql := "DELETE FROM t_whitelist WHERE del_time is not NULL AND account_id IN(" + input + ");"
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("删除白名单用户成功")
}
