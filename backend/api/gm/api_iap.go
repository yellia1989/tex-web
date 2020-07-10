package gm

import (
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func IAPDetail(c echo.Context) error {
	ctx := c.(*mid.Context)
	flowid := ctx.FormValue("flowid")

    comm := common.GetLocator()

	iapPrx := new(rpc.IAPService)
	comm.StringToProxy(common.GetApp() + ".IAPServer.IAPServiceObj", iapPrx)

    var stIAPReceiptInAll rpc.IAPReceiptInAll
	ret, err := iapPrx.GetReceiptStatusByFlow(flowid, &stIAPReceiptInAll)
	if err := checkRet(ret, err); err != nil {
		return err
	}

    return ctx.SendResponse(stIAPReceiptInAll)
}
