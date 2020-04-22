package gm

import (
    "strings"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
)

func zoneList(c echo.Context) ([]rpc.ZoneInfo) {
    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)

    var zones []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&zones)
    checkRet(ret, err)
    
    return zones
}

func ZoneSimpleList(c echo.Context) error {
    ctx := c.(*mid.Context)

    zones := zoneList(c)
    zones = append(zones, rpc.ZoneInfo{IZoneId: 0, SZoneName: "GFServer"})

    return ctx.SendResponse(zones)
}

func ZoneList(c echo.Context) error {
    ctx := c.(*mid.Context)

    zones := zoneList(c)

    return ctx.SendResponse(zones)
}

func ZoneAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    zone := rpc.ZoneInfo{}
    if err := ctx.Bind(&zone); err != nil {
        return err
    }

    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)
    ret, err := dirPrx.CreateZone(zone)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("添加分区成功")
}

func ZoneDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "分区不存在")
    }

    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        ret, err := dirPrx.DeleteZone(uint32(id))
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("删除分区成功")
}

func ZoneUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)

    zone := rpc.ZoneInfo{}
    if err := ctx.Bind(&zone); err != nil {
        return err
    }

    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)
    ret, err := dirPrx.ModifyZone(zone, rpc.ZoneModifyInfo{})
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("修改分区成功")
}

func ZoneUpdateVersion(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "分区不存在")
    }

    clientVersion := ctx.FormValue("sClientVersion")
    forceUpdateVersion := ctx.FormValue("sForceUpdateVersion")
    andClientVersion := ctx.FormValue("sAndClientVersion")
    andForceUpdateVersion := ctx.FormValue("sAndForceUpdateVersion")

    if len(strings.Split(clientVersion, ".")) != 5 || len(strings.Split(forceUpdateVersion, ".")) != 3 || len(strings.Split(andClientVersion, ".")) != 5 || len(strings.Split(andForceUpdateVersion, ".")) != 3 {
        return ctx.SendError(-1, "参数非法")
    }

    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        zone := rpc.ZoneInfo{}
        ret, err := dirPrx.GetZone(uint32(id), &zone)
        if err := checkRet(ret, err); err != nil {
            return err
        }

        zone.SClientVersion = clientVersion
        zone.SForceUpdateVersion = forceUpdateVersion
        zone.SAndClientVersion = andClientVersion
        zone.SAndForceUpdateVersion = andForceUpdateVersion
        ret, err = dirPrx.ModifyZone(zone, rpc.ZoneModifyInfo{})
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("批量修改分区成功")
}
