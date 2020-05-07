package gm

import (
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func DirtyTest(c echo.Context) error {
	ctx := c.(*mid.Context)

	sInput := ctx.FormValue("input")

	dirtyPrx := new(rpc.DirtyCheckService)
	comm.StringToProxy("aqua.DirtyCheckServer.DirtyCheckService", dirtyPrx)

	var sOutPut string
	ret, err := dirtyPrx.filter(sInput, sOutPut)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse(sOutPut)
}
