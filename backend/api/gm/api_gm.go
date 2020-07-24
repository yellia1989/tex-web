package gm

import (
    "fmt"
    "strings"
    "bytes"
    "strconv"
    "encoding/json"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/api/sys"
    "github.com/yellia1989/tex-web/backend/common"
)

func checkRet(ret int32, err error) error {
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        return fmt.Errorf("opt failed, ret:%d, err:%s", ret, serr)
    }

    return nil
}

func GameCmd(c echo.Context) error {
    ctx := c.(*mid.Context)
    szoneid := ctx.FormValue("zoneids")
    scmd := ctx.FormValue("cmd")

    if szoneid == "" || scmd == "" {
        return ctx.SendError(-1, "参数非法")
    }

    buff := bytes.Buffer{}
    u := ctx.GetUser()

    comm := common.GetLocator()
    app := common.GetApp()

    zoneids := strings.Split(szoneid, ",")
    for _,zoneid := range zoneids {
        izoneid,_ := strconv.Atoi(zoneid)
        gamePrx := new(rpc.GameService)
        gfPrx := new(rpc.GFService)
        mapPrx := new(rpc.MapService)
        if izoneid != 0 {
            if izoneid != 8888 && izoneid != 9999 && izoneid > 1000 {
                comm.StringToProxy(app+".MapServer.MapServiceObj%"+app+".map."+zoneid, mapPrx)
            } else {
                comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+zoneid, gamePrx)
            }
        } else {
            comm.StringToProxy(app+".GFServer.GFServiceObj", gfPrx)
        }

        cmds := strings.Split(scmd, "\n")
        for _, cmd := range cmds {
            cmd := strings.Trim(strings.ReplaceAll(cmd, "   ", ""), " ")

            result := ""
            var ret int32
            var err error

            buff.WriteString("zone["+zoneid + "] > " + cmd + "\n")

            if izoneid != 0 {
                if izoneid == 8888 || izoneid == 9999 || izoneid <= 1000 {
                    ret, err = gamePrx.DoGmCmd(u.UserName, cmd, &result)
                } else {
                    ret, err = mapPrx.DoGmCmd(u.UserName, cmd, &result)
                }
            } else {
                ret, err = gfPrx.DoGmCmd(u.UserName, cmd, &result)
            }
            if ret != 0 || err != nil {
                serr := ""
                if err != nil {
                    serr = err.Error()
                }
                result = fmt.Sprintf("ret:%d, err:%s\n", ret, serr)
            }
            buff.WriteString(result+"\n")
        }
    }

    sys.LogAdd(ctx, "gm", "[" + szoneid + "]>" + scmd)

    return ctx.SendResponse(buff.String())
}

func cmd(ctx *mid.Context, zoneid string, cmd string, result *string) error {
    comm := common.GetLocator()
    app := common.GetApp()

    gamePrx := new(rpc.GameService)
    comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+zoneid, gamePrx)

    cmd = strings.Trim(strings.ReplaceAll(cmd, "   ", ""), " ")

    var ret int32
    var err error

    u := ctx.GetUser()
    ret, err = gamePrx.DoGmCmd(u.UserName, cmd, result)
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        return fmt.Errorf("ret:%d, err:%s", ret, serr)
    }

    if cmd != "iap_list" && cmd != "item_list" {
        sys.LogAdd(ctx, "gm", "[" + zoneid + "]>" + cmd)
    }

    return nil
}

func IAPRecharge(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("zoneid")
    roleid := ctx.FormValue("roleid")
    productid := ctx.FormValue("productid")
    scmd := "recharge " + roleid + " " + productid

    if zoneid == "" || roleid == "" || productid == "" {
        return ctx.SendError(-1, "参数非法")
    }

    var result string
    err := cmd(ctx, zoneid, scmd, &result)
    if err !=  nil {
        return err
    }

    return ctx.SendResponse(result)
}

type _iap struct {
    Id uint32 `json:"id"`
    Type uint32 `json:"type"`
    Name string `json:"name"`
}
func IAPList(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("zoneid")
    scmd := "iap_list"

    if zoneid == "" {
        return ctx.SendError(-1, "参数非法")
    }

    var result string
    err := cmd(ctx, zoneid, scmd, &result)
    if err !=  nil {
        return err
    }

    iaps := make([]_iap,0)
    if err := json.Unmarshal([]byte(result), &iaps); err != nil {
        return err
    }

    return ctx.SendResponse(iaps)
}

type _item struct {
    ID uint32 `json:"value"`
    Name string `json:"name"`
}
func ItemList(c echo.Context) error {
    ctx := c.(*mid.Context)

    zoneid := "3"
    scmd := "item_list"
    var result string
    err := cmd(ctx, zoneid, scmd, &result)
    if err !=  nil {
        return err
    }

    items := make([]_item,0)
    if err := json.Unmarshal([]byte(result), &items); err != nil {
        return err
    }

    return ctx.SendResponse(items)
}

func BanSpeak(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("zoneid")
    roleid := ctx.FormValue("roleid")
    time := ctx.FormValue("time")
    scmd := "ban_speak " + roleid + " " + time

    if zoneid == "" || roleid == "" || time == "" {
        return ctx.SendError(-1, "参数非法")
    }

    var result string
    err := cmd(ctx, zoneid, scmd, &result)
    if err !=  nil {
        return err
    }

    return ctx.SendResponse(result)
}

func BanLogin(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("zoneid")
    roleid := ctx.FormValue("roleid")
    time := ctx.FormValue("time")
    scmd := "ban_login " + roleid + " " + time

    if zoneid == "" || roleid == "" || time == "" {
        return ctx.SendError(-1, "参数非法")
    }

    var result string
    err := cmd(ctx, zoneid, scmd, &result)
    if err !=  nil {
        return err
    }

    return ctx.SendResponse(result)
}

func RealMap(c echo.Context) error {
    ctx := c.(*mid.Context)
    mapid := ctx.FormValue("mapid")

    if mapid == "" {
        return ctx.SendError(-1, "参数非法")
    }

    comm := common.GetLocator()
    app := common.GetApp()

    mapPrx := new(rpc.MapService)
    comm.StringToProxy(app+".MapServer.MapServiceObj%"+app+".map."+mapid, mapPrx)

    cmd := "map_json"

    var result string
    u := ctx.GetUser()
    ret, err := mapPrx.DoGmCmd(u.UserName, cmd, &result)
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        return fmt.Errorf("ret:%d, err:%s", ret, serr)
    }

    return ctx.SendResponse(result)
}

func RealMapObj(c echo.Context) error {
    ctx := c.(*mid.Context)
    mapid := ctx.FormValue("mapid")
    objid := ctx.FormValue("objid")

    if objid == "" || mapid == "" {
        return ctx.SendError(-1, "参数非法")
    }

    comm := common.GetLocator()
    app := common.GetApp()

    mapPrx := new(rpc.MapService)
    comm.StringToProxy(app+".MapServer.MapServiceObj%"+app+".map."+mapid, mapPrx)

    cmd := "see_obj " + objid

    var result string
    u := ctx.GetUser()
    ret, err := mapPrx.DoGmCmd(u.UserName, cmd, &result)
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        return fmt.Errorf("ret:%d, err:%s", ret, serr)
    }

    return ctx.SendResponse(result)
}
