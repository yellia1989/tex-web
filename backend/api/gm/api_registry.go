package gm

import (
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/sdp/rpc"
)

func RegistryList(c echo.Context) error {
    ctx := c.(*mid.Context)

    queryPrx := new(rpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    var vObj []rpc.ObjEndpoint
    ret, err := queryPrx.GetAllEndpoints(&vObj)
    if err := checkRet(ret, err); err != nil {
        return err
    }
    
    return ctx.SendResponse(vObj)
}

func RegistryAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    return ctx.SendResponse("添加registry成功")
}

func RegistryDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    return ctx.SendResponse("删除registry成功")
}

func RegistryUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)

    return ctx.SendResponse("修改registry成功")
}
