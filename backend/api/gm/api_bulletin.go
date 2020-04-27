package gm

import (
    "strings"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
)

func BulletinList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

    var vBulletin []rpc.BulletinDataInfo
    ret, err := bulletinPrx.GetAllBulletin(&vBulletin)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    vPage := common.GetPage(vBulletin, page, limit)
    return ctx.SendArray(vPage, len(vBulletin))
}

func BulletinAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    bulletin := rpc.BulletinDataInfo{}
    if err := ctx.Bind(&bulletin); err != nil {
        return err
    }

    sTitle := ctx.FormValue("sTitle")
    sContent := ctx.FormValue("sContent")
    iBeginTime := ctx.FormValue("iBeginTime")
    iEndTime := ctx.FormValue("iEndTime")

    if sTitle == "" || sContent == "" || iBeginTime == "" || iEndTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

    ret, err := bulletinPrx.AddBulletin(bulletin)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("添加公告成功")
}

func BulletinDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := strings.Split(ctx.FormValue("idsStr"), ",")

    if len(ids) == 0 {
        return ctx.SendError(-1, "公告不存在")
    }

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

    for _, id := range ids  {
        id, _ := strconv.ParseUnit(id, 10, 32)
        ret, err := bulletinPrx.DelBulletin(uint32(id))
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("删除公告成功")
}

func BulletinUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)

    bulletin := rpc.BulletinDataInfo{}
    if err := ctx.Bind(&bulletin); err != nil {
        return err
    }

    sTitle := ctx.FormValue("sTitle")
    sContent := ctx.FormValue("sContent")
    iBeginTime := ctx.FormValue("iBeginTime")
    iEndTime := ctx.FormValue("iEndTime")

    if sTitle == "" || sContent == "" || iBeginTime == "" || iEndTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

    ret, err := bulletinPrx.ModifyBulletin(bulletin)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("修改公告成功")
}
