package gm

import (
	"database/sql"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func parseIDStr(src string, out *string) {
	reg, _ := regexp.Compile("/[,;\r\n\t ]/")
	vStr := reg.FindAllString(src, -1)

	*out = strings.Join(vStr, "),(")
}

func WhiteList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16)/db_loginserver?charset=utf8")
	defer db.Close()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT * FROM t_whitelist;")
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

	return ctx.SendResponse(strings.Join(vStr, ";"))
}

func WhiteAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	if input == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16)/db_loginserve?charset=utf8")
	defer db.Close()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT IGNORE INTO t_whitelist VALUES(?)")
	if err != nil {
		return err
	}

	parseIDStr(input, &input)
	_, err = stmt.Exec(input)
	if err != nil {
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

	db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16)/db_loginserve?charset=utf8")
	defer db.Close()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM t_whitelist WHERE account_id IN(?)")
	if err != nil {
		return err
	}

	parseIDStr(input, &input)
	_, err = stmt.Exec(input)
	if err != nil {
		return err
	}

	return ctx.SendResponse("删除白名单用户成功")
}

func WhiteReplace(c echo.Context) error {
	ctx := c.(*mid.Context)

	input := ctx.FormValue("input")

	db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16)/db_loginserve?charset=utf8")
	defer db.Close()
	if err != nil {
		return err
	}

	db.Exec("DELETE FROM t_whitelist;")
	stmt, err := db.Prepare("INSERT IGNORE INTO t_whitelist VALUES(?)")
	if err != nil {
		return err
	}

	parseIDStr(input, &input)
	_, err = stmt.Exec(input)
	if err != nil {
		return err
	}

	return ctx.SendResponse("覆盖白名单用户成功")
}
