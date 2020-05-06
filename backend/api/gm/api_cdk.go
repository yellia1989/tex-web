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
	ret, err := mpPrx.GetAllNotice(&vCDK)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	vPage := common.GetPage(vCDK, page, limit)
	return ctx.SendArray(vPage, len(vCDK))
}

func CDKAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		cdk := rpc.NoticeDataInfo{}
		if err := ctx.Bind(&notice); err != nil {
			return err
		}

		sBeginTime := ctx.FormValue("sBeginTime")
		sEndTime := ctx.FormValue("sEndTime")

		zones := strings.Split(ctx.FormValue("zoneid"), ",")
		for _, zone := range zones {
			zone, _ := strconv.ParseUint(zone, 10, 32)
			notice.VZoneId = append(notice.VZoneId, uint32(zone))
		}

		if sBeginTime == "" || sEndTime == "" || len(notice.VZoneId) == 0 {
			return ctx.SendError(-1, "参数非法")
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		ret, err := bulletinPrx.AddNotice(notice)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	*/

	return ctx.SendResponse("添加CDK成功")
}

func CDKDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		ids := strings.Split(ctx.FormValue("idsStr"), ",")

		if len(ids) == 0 {
			return ctx.SendError(-1, "跑马灯不存在")
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		for _, id := range ids {
			id, _ := strconv.ParseUint(id, 10, 32)
			ret, err := bulletinPrx.DelNotice(uint32(id))
			if err := checkRet(ret, err); err != nil {
				return err
			}
		}
	*/

	return ctx.SendResponse("删除CDK成功")
}

func CDKUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		notice := rpc.NoticeDataInfo{}
		if err := ctx.Bind(&notice); err != nil {
			return err
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		ret, err := bulletinPrx.ModifyNotice(notice)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	*/

	return ctx.SendResponse("修改CDK成功")
}
