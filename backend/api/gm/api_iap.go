package gm

import (
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"github.com/yellia1989/tex-web/backend/service"
)

func IAPDetail(c echo.Context) error {
	ctx := c.(*mid.Context)
	flowid := ctx.FormValue("flowid")

	comm := cfg.Comm

	iapPrx := new(rpc.IAPService)
	comm.StringToProxy(service.GetIAPServiceName(), iapPrx)

	var stIAPReceiptInAll rpc.IAPReceiptInAll
	ret, err := iapPrx.GetReceiptStatusByFlow(flowid, &stIAPReceiptInAll)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse(stIAPReceiptInAll)
}
