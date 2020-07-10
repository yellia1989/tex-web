package gm

import (
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func DirtyTest(c echo.Context) error {
	ctx := c.(*mid.Context)

	sInput := ctx.FormValue("input")

    comm := common.GetLocator()

	dirtyPrx := new(rpc.DirtyCheckService)
	comm.StringToProxy(common.GetApp() + ".DirtyCheckServer.DirtyCheckServiceObj", dirtyPrx)

	var sOutPut string
	ret, err := dirtyPrx.Filter(sInput, &sOutPut)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse(sOutPut)
}
