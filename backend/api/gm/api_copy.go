package gm

import (
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/util"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "errors"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
)

type Server struct {
    Name string `json:"name"`
    Ip string `json:"ip"`
    Port int `json:"port"`
}

func (s *Server) GetAddress() string {
    return fmt.Sprintf("%s:%s", s.Ip, s.Port)
}

var AllServer = make(map[string]Server)

func init() {
    servers, err := util.LoadFromFile("data/server.json")
    if err != nil {
        fmt.Printf("servers init failed, %s", err.Error())
    }
    err = json.Unmarshal(servers, AllServer)
    if err != nil {
        fmt.Printf("servers init failed, %s", err.Error())
    }

    fmt.Println(AllServer)
}

func getOtherServerAddress(id string) (string, error) {
    // 查找其他服务器地址
    server, ok := AllServer[id]
    if !ok {
        return "", errors.New("没有找到对应的服务器")
    }

    return server.GetAddress(), nil
}

func GetServerList(c echo.Context) error {
    ctx := c.(*mid.Context)

    type server struct {
        Id string `json:"id"`
        Name string `json:"name"`
    }
    mServers := make(map[string]server)

    for id, v := range AllServer {
        ser := server{id, v.Name}
        mServers[ser.Id] = ser
    }

    return ctx.SendResponse(mServers)
}

func GetZoneList(c echo.Context) error {
    ctx := c.(*mid.Context)
    server := ctx.QueryParam("server")
    if server == cfg.ServerID {
        data := make(map[string][]rpc.ZoneInfo,0)
        zones := zoneList(c)
        for i,_ := range zones {
            zones[i].SZoneName = fmt.Sprintf("%s(%d)", zones[i].SZoneName, zones[i].IZoneId)
        }
        data["game"] = zones

        zones2 := MapSimpleList()
        data["map"] = zones2
        return ctx.SendResponse(data);
    } else {
        // 查询其他服务器列表
        address, err := getOtherServerAddress(server)
        if err != nil {
            return err
        }

        params := url.Values{}
        Url, err := url.Parse(address + "/get")
        if err != nil {
            return err
        }

        params.Set("server", server)
        Url.RawQuery = params.Encode()
        path := Url.String()
        fmt.Println(path)
        resp, err := http.Get(path)
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        return ctx.SendResponse(body)
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

    if server == cfg.ServerID {
        result := new (string)
        cmdstr := "copy_role " + role
        cmd(ctx, zone, cmdstr, result)
        return ctx.SendResponse(*result)
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

    result := new (string)
    cmdstr := "paste_role " + role + " " + roleData
    cmd(ctx, zone, cmdstr, result)

    return ctx.SendResponse(*result)
}
