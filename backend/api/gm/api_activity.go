package gm

import (
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func ActivityList(c echo.Context) error {
	ctx := c.(*mid.Context)

	return ctx.SendResponse("list")
}
