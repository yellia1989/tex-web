package gm

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func ActivityList(c echo.Context) error {
	ctx := c.(*mid.Context)
	db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16)/db_loginserver")
	defer db.Close()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT * FROM t_whitelist;")
	if err != nil {
		return err
	}
	defer rows.Close()

	str := ""
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
		str += strconv.Itoa(id)
	}

	return ctx.SendResponse(str)
}
