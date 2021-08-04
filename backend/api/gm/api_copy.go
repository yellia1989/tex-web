package gm

import (
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

func getOtherServerAddress(id string) error {
    // 解析配置 查找地址

    return nil
}

func GetZoneList(c echo.Context) error {
ctx := c.(*mid.Context)
         server := ctx.QueryParam("server")
    own := ""
    if server == own {

    } else {
        // 查询其他服务器列表
    }

    return nil
}

func CopyRole(c echo.Context) error {
ctx := c.(*mid.Context)
         server := ctx.QueryParam("server")
         zone := ctx.QueryParam("zone")
         role := ctx.QueryParam("role")

         if server == "" || zone == "" || role == "" {
             return ctx.SendError(-1, "选择的玩家错误")
         }

    own := ""
    if server == own {

    } else {
        // 从其他服务器获取玩家数据
    }
    return nil
}

func PasteRole(c echo.Context) error {
    // 执行gm粘贴玩家数据
ctx := c.(*mid.Context)
         zone := ctx.QueryParam("zone")
         role := ctx.QueryParam("role")
         roleData := ctx.QueryParam("data")

         if zone == "" || role == "" || roleData == "" {
             return ctx.SendError(-1, "选择的玩家错误")
         }

    return nil
}
