package gm

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
)

func ZoneMap() map[uint32]rpc.ZoneInfo {
    comm := common.GetLocator()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(common.GetApp()+".DirServer.DirServiceObj", dirPrx)

    mzone := make(map[uint32]rpc.ZoneInfo)

    var zones []rpc.ZoneInfo
    dirPrx.GetAllZone(&zones)
    for _, v := range zones {
        mzone[v.IZoneId] = *(v.Copy())
    }
    return mzone
}

func zoneList(c echo.Context) ([]rpc.ZoneInfo) {
    comm := common.GetLocator()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(common.GetApp()+".DirServer.DirServiceObj", dirPrx)

    var zones []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&zones)
    checkRet(ret, err)
    
    return zones
}

func ZoneSimpleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    gf := ctx.FormValue("gf")
    all := ctx.FormValue("all")

    zones := zoneList(c)

    for i,_ := range zones {
        zones[i].SZoneName = fmt.Sprintf("%s(%d)", zones[i].SZoneName, zones[i].IZoneId)
    }

    if gf != "" {
        zones = append(zones, rpc.ZoneInfo{IZoneId: 0, SZoneName: "GFServer"})
    }

    var zones2 []rpc.ZoneInfo
    if all != "" {
        zones2 = append(zones2, rpc.ZoneInfo{IZoneId: 99999, SZoneName: "全服"})
    }
    zones2 = append(zones2, zones...)

    return ctx.SendResponse(zones2)
}

func ZoneList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    zones := zoneList(c)
    vPage := common.GetPage(zones, page, limit)

    return ctx.SendArray(vPage, len(zones))
}

func ZoneAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    zone := rpc.NewZoneInfo()
    if err := ctx.Bind(zone); err != nil {
        return err
    }

    sDivision := fmt.Sprintf(common.GetApp()+".zone.%d", zone.IZoneId)
    sHandleConnEp := ctx.FormValue("sHandleConn")
    sConnServiceObjEp := ctx.FormValue("sConnServiceObj")
    sGameServiceObjEp := ctx.FormValue("sGameServiceObj")

    if err := registryAdd(common.GetApp()+".ConnServer.HandleConn", sDivision, sHandleConnEp); err != nil {
        return err
    }
    if err := registryAdd(common.GetApp()+".ConnServer.ConnServiceObj", sDivision, sConnServiceObjEp); err != nil {
        registryDel(common.GetApp()+".ConnServer.HandleConn", sDivision, sHandleConnEp)
        return err
    }
    if err := registryAdd(common.GetApp()+".GameServer.GameServiceObj", sDivision, sGameServiceObjEp); err != nil {
        registryDel(common.GetApp()+".ConnServer.HandleConn", sDivision, sHandleConnEp)
        registryDel(common.GetApp()+".ConnServer.ConnServiceObj", sDivision, sConnServiceObjEp)
        return err
    }

    comm := common.GetLocator()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(common.GetApp()+".DirServer.DirServiceObj", dirPrx)
    ret, err := dirPrx.CreateZone(*zone.Copy())
    if err := checkRet(ret, err); err != nil {
        registryDel(common.GetApp()+".ConnServer.HandleConn", sDivision, sHandleConnEp)
        registryDel(common.GetApp()+".ConnServer.ConnServiceObj", sDivision, sConnServiceObjEp)
        registryDel(common.GetApp()+".GameServer.GameServiceObj", sDivision, sGameServiceObjEp)
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

    comm := common.GetLocator()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(common.GetApp()+".DirServer.DirServiceObj", dirPrx)

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

    zone := rpc.NewZoneInfo()
    if err := ctx.Bind(zone); err != nil {
        return err
    }

    comm := common.GetLocator()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(common.GetApp()+".DirServer.DirServiceObj", dirPrx)
    ret, err := dirPrx.ModifyZone(*zone.Copy(), rpc.ZoneModifyInfo{})
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

    comm := common.GetLocator()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(common.GetApp()+".DirServer.DirServiceObj", dirPrx)

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        zone := rpc.NewZoneInfo()
        ret, err := dirPrx.GetZone(uint32(id), zone)
        if err := checkRet(ret, err); err != nil {
            return err
        }

        zone.SClientVersion = clientVersion
        zone.SForceUpdateVersion = forceUpdateVersion
        zone.SAndClientVersion = andClientVersion
        zone.SAndForceUpdateVersion = andForceUpdateVersion
        ret, err = dirPrx.ModifyZone(*zone.Copy(), rpc.ZoneModifyInfo{})
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("批量修改分区成功")
}
