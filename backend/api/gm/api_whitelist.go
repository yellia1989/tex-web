package gm

import (
	"regexp"
	"strings"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func parseIDStr(src, sep string, out *string) {
	reg, _ := regexp.Compile("\\d{5}")
	vStr := reg.FindAllString(src, -1)

	if len(vStr) != 0 {
		*out = strings.Join(vStr, sep)
	}
}

func WhiteList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_loginserver")
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT * FROM t_whitelist")
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

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_loginserver")
	if err != nil {
		return err
	}

	parseIDStr(input, "),(", &input)
	sql := "INSERT IGNORE INTO t_whitelist VALUES(" + input + ")"
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

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_loginserver")
	if err != nil {
		return err
	}

	parseIDStr(input, ",", &input)
	sql := "DELETE FROM t_whitelist WHERE account_id IN(" + input + ");"
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

    db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_loginserver")
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM t_whitelist")
	if err != nil {
		return err
	}

	parseIDStr(input, "),(", &input)
	sql := "INSERT IGNORE INTO t_whitelist VALUES(" + input + ");"
	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

    if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendResponse("覆盖白名单用户成功")
}
