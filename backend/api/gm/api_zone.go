package gm

import (
    "fmt"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
)

func ZoneSimpleList(c echo.Context) error {
    ctx := c.(*mid.Context)

    dirPrx := new(rpc.DirService)
    comm.StringToProxy("aqua.DirServer.DirServiceObj", dirPrx)

    var zones []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&zones)
    if ret != 0 || err != nil {
        ctx.SendError(-1, fmt.Sprintf("ret:%d, err:%s", ret, err.Error()))
    }

    return ctx.SendResponse(zones)
}
