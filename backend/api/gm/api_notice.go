package gm

import (
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func NoticeList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	bulletinPrx := new(rpc.BulletinService)
	comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

	var vNotice []rpc.NoticeDataInfo
	ret, err := bulletinPrx.GetAllNotice(&vNotice)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	vPage := common.GetPage(vNotice, page, limit)
	return ctx.SendArray(vPage, len(vNotice))
}

func NoticeAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	notice := rpc.NoticeDataInfo{}
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

	return ctx.SendResponse("添加跑马灯成功")
}

func NoticeDel(c echo.Context) error {
	ctx := c.(*mid.Context)

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

	return ctx.SendResponse("删除跑马灯成功")
}

func NoticeUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

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

	return ctx.SendResponse("修改跑马灯成功")
}
