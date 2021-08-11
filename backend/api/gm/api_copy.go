package gm

import (
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/util"
    "github.com/yellia1989/tex-go/tools/log"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "errors"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type Server struct {
    Name string `json:"name"`
    Ip string `json:"ip"`
    Port string `json:"port"`
}

func (s *Server) GetAddress() string {
    return fmt.Sprintf("%s:%s", s.Ip, s.Port)
}

var AllServer = make(map[string]Server)

func init() {
    servers, err := util.LoadFromFile("data/server.json")
    if err != nil {
        log.Error("servers init failed, %s", err.Error())
    }
    err = json.Unmarshal(servers, &AllServer)
    if err != nil {
        log.Error("servers init failed, %s", err.Error())
    }
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
    vServers := make([]server, 0)

    for id, v := range AllServer {
        ser := server{id, v.Name}
        vServers = append(vServers, ser)
    }

    return ctx.SendArray(vServers, len(vServers))
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
    }

    // 查询其他服务器列表
    address, err := getOtherServerAddress(server)
    if err != nil {
        return err
    }

    path := "/api/public/gm/get_zone_list"
    param := "server=" + server

    p := make(map[string]string)
    p["server"] = server
    param += "&sign=" + mid.Sign(path, p)

    url := "http://" + address + path + "/?" + param
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    type apiResponse struct {
        Code int `json:"code"`
        Msg string `json:"msg"`
        Data interface{} `json:"data"`
    }

    var data apiResponse
    err = json.Unmarshal(body, &data)
    if err != nil {
        return err
    }

    if data.Code != 0 {
        return fmt.Errorf(data.Msg)
    }

    return ctx.SendResponse(data.Data)
}

func DumpRole(c echo.Context) error {
    ctx := c.(*mid.Context)
    server := ctx.QueryParam("server")
    zone := ctx.QueryParam("zone")
    role := ctx.QueryParam("role")

    if server == "" || zone == "" || role == "" {
        return ctx.SendError(-1, "选择的玩家错误")
    }

    if server == cfg.ServerID {
        result := new (string)
        cmdstr := "dump_role_http " + role
        err := Cmd("public_gm", zone, "0", cmdstr, result)
        if err != nil {
            return err
        }
        return ctx.SendResponse(*result)
    }

    // 从其他服务器获取玩家数据
    address, err := getOtherServerAddress(server)
    if err != nil {
        return err
    }

    path := "/api/public/gm/dump_role"
    param := "server=" + server
    param += "&zone=" + zone
    param += "&role=" + role

    p := make(map[string]string)
    p["server"] = server
    p["zone"] = zone
    p["role"] = role

    param += "&sign=" + mid.Sign(path, p)

    url := "http://" + address + path + "/?" + param
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    type apiResponse struct {
        Code int `json:"code"`
        Msg string `json:"msg"`
        Data interface{} `json:"data"`
    }

    var data apiResponse
    err = json.Unmarshal(body, &data)
    if err != nil {
        return err
    }

    if data.Code != 0 {
        return fmt.Errorf(data.Msg)
    }

    return ctx.SendResponse(data.Data)
}

func LoadRole(c echo.Context) error {
    // 执行gm粘贴玩家数据
    ctx := c.(*mid.Context)
    zone := ctx.FormValue("zone")
    role := ctx.FormValue("role")
    roleData := ctx.FormValue("data")

    if zone == "" || role == "" || roleData == "" {
        return ctx.SendError(-1, "选择的玩家错误")
    }

    result := new (string)
    cmdstr := "load_role_http " + role + " " + roleData
    if err := cmd(ctx, zone, cmdstr, result); err != nil {
        return err
    }

    return ctx.SendResponse(*result)
}
