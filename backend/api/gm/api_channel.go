package gm

import (
    "strings"
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/common"
)

func ChannelList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    var channels []rpc.ChannelAddr
    ret, err := loginPrx.GetAllChannel(&channels)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    vPage := common.GetPage(channels, page, limit)
    
    return ctx.SendArray(vPage, len(channels))
}

func ChannelAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    sChannel := ctx.FormValue("sChannel")
    sRes := ctx.FormValue("sRes")
    sShopVer := ctx.FormValue("sShopVer")
    if sChannel == "" || sRes == "" || sShopVer == "" {
        return ctx.SendError(-1, "参数非法")
    }

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    ret, err := loginPrx.AddNewChannel(sChannel, "", sRes, sShopVer)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("添加渠道成功")
}

func ChannelDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "渠道不存在")
    }

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    for _, id := range ids {
        ret, err := loginPrx.DelChannel(id)
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("删除渠道成功")
}

func ChannelUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)

    sChannel := ctx.FormValue("sChannel")
    sRes := ctx.FormValue("sRes")
    sShopVer := ctx.FormValue("sShopVer")
    if sChannel == "" || sRes == "" || sShopVer == "" {
        return ctx.SendError(-1, "参数非法")
    }

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    ret, err := loginPrx.ModifyChannel(sChannel, "", sRes, sShopVer)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("修改渠道成功")
}
