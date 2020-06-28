package gm

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func CDKList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	mpPrx := new(rpc.MPService)
	comm.StringToProxy("aqua.MPServer.MPServiceObj", mpPrx)

	var vCDK []rpc.CDKeyConfig
	str := "aqua"
	var a uint32
	ret, err := mpPrx.GetCDKeyList(str, 0, uint32(page*limit), &vCDK, &a)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	vPage := common.GetPage(vCDK, page, limit)
	return ctx.SendArray(vPage, len(vCDK))
}

func CDKAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	CDKey := rpc.CDKeyConfig{}
	if err := ctx.Bind(&CDKey); err != nil {
		return err
	}
    CDKey.ICreateMode = 1
    CDKey.IDeliveryMode = 1

	sBeginTime := ctx.FormValue("iBeginTime")
	sEndTime := ctx.FormValue("iEndTime")

	if sBeginTime == "" || sEndTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	MPPrx := new(rpc.MPService)
	comm.StringToProxy("aqua.MPServer.MPServiceObj", MPPrx)

	projectConfig := rpc.MPProjectConfig{}
	if ok, _ := MPPrx.GetProject("aqua", &projectConfig); ok != 0 {
		projectConfig.SProjectId = "aqua"
		projectConfig.SProjectName = "aqua"

		ret, err := MPPrx.CreateProject(projectConfig)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	CDKey.SProjectId = "aqua"
	var iCDKeyId uint32
	ret, err := MPPrx.CreateCDKey(CDKey, &iCDKeyId)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("添加CDK成功")
}

func CDKUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

	CDKey := rpc.CDKeyConfig{}
	if err := ctx.Bind(&CDKey); err != nil {
		return err
	}

	MPPrx := new(rpc.MPService)
	comm.StringToProxy("aqua.MPServer.MPServiceObj", MPPrx)

	ret, err := MPPrx.ModifyCDKey(CDKey)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("修改CDK成功")
}
