package gm

import (
	"strconv"
    "strings"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func CDKList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    comm := cfg.Comm

	mpPrx := new(rpc.MPService)
	comm.StringToProxy(cfg.App+".MPServer.MPServiceObj", mpPrx)

	var vCDK []rpc.CDKeyConfig
	str := cfg.App+""

	var l uint32
    from := (page-1) * limit
    to := from + limit
	ret, err := mpPrx.GetCDKeyList(str, uint32(from), uint32(to), &vCDK, &l)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendArray(vCDK, int(l))
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

    comm := cfg.Comm

	MPPrx := new(rpc.MPService)
	comm.StringToProxy(cfg.App+".MPServer.MPServiceObj", MPPrx)

	projectConfig := rpc.MPProjectConfig{}
	if ok, _ := MPPrx.GetProject(cfg.App+"", &projectConfig); ok != 0 {
		projectConfig.SProjectId = cfg.App+""
		projectConfig.SProjectName = cfg.App+""

		ret, err := MPPrx.CreateProject(projectConfig)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	CDKey.SProjectId = cfg.App+""
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

    comm := cfg.Comm

	MPPrx := new(rpc.MPService)
	comm.StringToProxy(cfg.App+".MPServer.MPServiceObj", MPPrx)

	ret, err := MPPrx.ModifyCDKey(CDKey)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("修改CDK成功")
}

func CDKExport(c echo.Context) error {
	ctx := c.(*mid.Context)
    idsStr := ctx.FormValue("idsStr")

    if idsStr == "" {
        return ctx.SendError(-1, "参数非法")
    }

    comm := cfg.Comm

	MPPrx := new(rpc.MPService)
	comm.StringToProxy(cfg.App+".MPServer.MPServiceObj", MPPrx)

    data := make(map[int]*string,0)
    for _,ids := range strings.Split(idsStr, ",") {
        id,_ := strconv.Atoi(ids)
        var cdks string
	    ret, err := MPPrx.ExportCDKey(uint32(id), &cdks)
	    if err := checkRet(ret, err); err != nil {
	    	return err
	    }
        data[id] = &cdks
    }

	return ctx.SendResponse(data)
}
