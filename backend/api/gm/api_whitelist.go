package gm

import (
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func WhiteList(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		var vNotice []rpc.NoticeDataInfo
		ret, err := bulletinPrx.GetAllNotice(&vNotice)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	*/

	return ctx.SendResponse("11019")
}

func WhiteAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		notice := rpc.NoticeDataInfo{}
		if err := ctx.Bind(&notice); err != nil {
			return err
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		ret, err := bulletinPrx.AddNotice(notice)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	*/

	return ctx.SendResponse("添加白名单用户成功")
}

func WhiteDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		ids := strings.Split(ctx.FormValue("idsStr"), ",")

		if len(ids) == 0 {
			return ctx.SendError(-1, "白名单用户不存在")
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

	return ctx.SendResponse("删除白名单用户成功")
}

func WhiteReplace(c echo.Context) error {
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

	return ctx.SendResponse("覆盖白名单用户成功")
}
