package gm

import (
    "fmt"
    "strconv"
    "strings"
    "bytes"
    "encoding/json"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/api/sys"
    "github.com/yellia1989/tex-web/backend/cfg"
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
    scmd := strings.ReplaceAll(strings.TrimSpace(ctx.FormValue("cmd")), "\t", " ")

    if szoneid == "" || scmd == "" {
        return ctx.SendError(-1, "参数非法")
    }

    u := ctx.GetUser()
    if u == nil {
        return ctx.SendError(-1, "账号不存在")
    }

    if !u.CheckGmPermission(scmd) {
        return ctx.SendError(-1, "账号GM权限不足")
    }

    buff := bytes.Buffer{}

    comm := cfg.Comm
    app := cfg.App

    zoneids := strings.Split(szoneid, ",")
    for _,zoneid := range zoneids {
        izoneid := common.Atou32(zoneid)
        gamePrx := new(rpc.GameService)
        if izoneid != 0 {
            comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+zoneid, gamePrx)
        }

        cmds := strings.Split(scmd, "\n")
        for _, cmd := range cmds {
            cmd := strings.Trim(strings.ReplaceAll(cmd, "   ", ""), " ")

            result := ""
            var ret int32
            var err error

            buff.WriteString("zone["+zoneid + "] > " + cmd + "\n")

            if izoneid != 0 {
                ret, err = gamePrx.DoGmCmd(u.UserName, cmd, &result)
            }
            if ret != 0 || err != nil {
                serr := ""
                if err != nil {
                    serr = err.Error()
                }
                result = fmt.Sprintf("ret:%s, err:%s\n", rpc.ErrorCode(ret), serr)
            }
            buff.WriteString(result+"\n")
        }
    }

    sys.LogAdd(u.UserName, "gm", "[" + szoneid + "]>" + scmd)

    return ctx.SendResponse(buff.String())
}

func cmd(ctx *mid.Context, zoneid string, cmd string, result *string) error {
    u := ctx.GetUser()

    zoneid2 := getZoneId(common.Atou32(zoneid)) 
    if zoneid2 == 0 {
        return fmt.Errorf("can't find merge zoneid: %s", zoneid)
    }

    return Cmd(u.UserName, zoneid, "0", cmd, result)
}

func Cmd(userName string, zoneid string, mapid string, cmd string, result *string) error {
    comm := cfg.Comm
    app := cfg.App

    gamePrx := new(rpc.GameService)

    comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+zoneid, gamePrx)

    cmd = strings.Trim(strings.ReplaceAll(cmd, "   ", ""), " ")

    var ret int32
    var err error

    ret, err = gamePrx.DoGmCmd(userName, cmd, result)
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        return fmt.Errorf("ret:%d, err:%s, rsp: %s", ret, serr, *result)
    }

    if cmd != "iap_list" && cmd != "item_list" {
        sys.LogAdd(userName, "gm", "[" + zoneid + "]>" + cmd)
    }

    return nil
}

func IAPRecharge(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.FormValue("zoneid")
    roleid := ctx.FormValue("roleid")
    productid := ctx.FormValue("productid")
    scmd := "recharge " + roleid + "," + zoneid + " " + productid

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
    key := ctx.QueryParam("key")

    zones := updateZoneList(false)
    if (len(zones) == 0) {
        return ctx.SendError(-1, "分区列表为空")
    }

    var zoneid uint64 = uint64(zones[0].IZoneId)
    scmd := "item_list"
    if key != "" {
        scmd += " " + key
    }
    var result string
    err := cmd(ctx, strconv.FormatUint(zoneid, 10), scmd, &result)
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
