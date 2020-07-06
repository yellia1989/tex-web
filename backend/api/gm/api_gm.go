package gm

import (
    "fmt"
    "strings"
    "bytes"
    "encoding/json"
    "github.com/labstack/echo"
    tex "github.com/yellia1989/tex-go/service"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/api/sys"
)

var (
    comm = tex.NewCommunicator("tex.mfwregistry.QueryObj@tcp -h 192.168.0.16 -p 2000 -t 3600000")
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

    zoneids := strings.Split(szoneid, ",")
    for _,zoneid := range zoneids {
        gamePrx := new(rpc.GameService)
        gfPrx := new(rpc.GFService)
        if zoneid != "0" {
            comm.StringToProxy("aqua.GameServer.GameServiceObj%aqua.zone."+zoneid, gamePrx)
        } else {
            comm.StringToProxy("aqua.GFServer.GFServiceObj", gfPrx)
        }

        cmds := strings.Split(scmd, "\n")
        for _, cmd := range cmds {
            cmd := strings.Trim(strings.ReplaceAll(cmd, "   ", ""), " ")

            result := ""
            var ret int32
            var err error

            buff.WriteString("zone["+zoneid + "] > " + cmd + "\n")

            if zoneid != "0" {
                ret, err = gamePrx.DoGmCmd(u.UserName, cmd, &result)
            } else {
                ret, err = gfPrx.DoGmCmd(u.UserName, cmd, &result)
            }
            if ret != 0 || err != nil {
                serr := ""
                if err != nil {
                    serr = err.Error()
                }
                result = fmt.Sprintf("ret:%d, err:%s", ret, serr)
            }
            buff.WriteString(result+"\n")
        }
    }

    sys.LogAdd(ctx, "gm", "[" + szoneid + "]>" + scmd)

    return ctx.SendResponse(buff.String())
}

func cmd(ctx *mid.Context, zoneid string, cmd string, result *string) error {
    gamePrx := new(rpc.GameService)
    comm.StringToProxy("aqua.GameServer.GameServiceObj%aqua.zone."+zoneid, gamePrx)

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

    zoneid := "1"
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
