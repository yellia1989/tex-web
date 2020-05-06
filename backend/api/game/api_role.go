package game

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

func RoleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    name := ctx.QueryParam("name")

    _ = zoneid
    _ = name

    return ctx.SendError(-1, "暂未实现")
}
