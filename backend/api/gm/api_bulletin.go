package gm

import (
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"strings"
)

func BulletinList(c echo.Context) error {
	ctx := c.(*mid.Context)

	var bulletins []string
	return ctx.SendResponse(bulletins)
}

func BulletinAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	sTitle := ctx.FormValue("sTitle")
	sContent := ctx.FormValue("sContent")
	iBeginTime := ctx.FormValue("iBeginTime")
	iEndTime := ctx.FormValue("iEndTime")
	iIsDisplay := ctx.FormValue("iIsDisplay")

	if sTitle == "" || sContent == "" || iBeginTime == "" || iEndTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	return ctx.SendResponse("添加公告成功")
}

func BulletinDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	ids := strings.Split(ctx.FormValue("idsStr"), ",")

	if len(ids) == 0 {
		return ctx.SendError(-1, "公告不存在")
	}

	return ctx.SendResponse("删除公告成功")
}

func BulletinUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

	sTitle := ctx.FormValue("sTitle")
	sContent := ctx.FormValue("sContent")
	iBeginTime := ctx.FormValue("iBeginTime")
	iEndTime := ctx.FormValue("iEndTime")
	iIsDisplay := ctx.FormValue("iIsDisplay")

	if sTitle == "" || sContent == "" || iBeginTime == "" || iEndTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	return ctx.SendResponse("修改公告成功")
}
