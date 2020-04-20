package gm

import (
    "fmt"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
)

func zoneList(c echo.Context) ([]rpc.ZoneInfo) {
    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)

    var zones []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&zones)
    if ret != 0 || err != nil {
        c.Error(fmt.Errorf("get zone list failed, ret:%d, err:%s", ret, err.Error()))
    }
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
