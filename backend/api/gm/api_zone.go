package gm

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
)

func checkRet(ret int32, err error, c echo.Context) bool {
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        c.Error(fmt.Errorf("opt zone failed, ret:%d, err:%s", ret, serr))
        return false
    }

    return true
}

func zoneList(c echo.Context) ([]rpc.ZoneInfo) {
    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)

    var zones []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&zones)
    checkRet(ret, err, c)
    
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
    if !checkRet(ret, err, c) {
        return ctx.SendError(-1, "添加分区失败")
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
        if !checkRet(ret, err, c) {
            return ctx.SendError(-1, "删除分区失败")
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
    if !checkRet(ret, err, c) {
        return ctx.SendError(-1, "修改分区失败")
    }

    return ctx.SendResponse("修改分区成功")
}
