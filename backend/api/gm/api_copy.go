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
    } else {
        // 查询其他服务器列表
        address, err := getOtherServerAddress(server)
        if err != nil {
            return err
        }

        cmd := "/api/public/gm/zone/get_list"
        param := "server=" + server
        path := "http://" + address + cmd + "/?" + param
        resp, err := http.Get(path)
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        mAllData := make(map[string]interface{})
        err = json.Unmarshal(body, &mAllData)
        if err != nil {
            return err
        }
        data, ok := mAllData["data"]
        if !ok {
          return errors.New("未找到数据")
        }
        return ctx.SendResponse(data)
    }

    return nil
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
    } else {
        // 从其他服务器获取玩家数据
        address, err := getOtherServerAddress(server)
        if err != nil {
            return err
        }

        cmd := "/api/public/gm/copy_role"
        param := "server=" + server
        param += "&zone=" + zone
        param += "&role=" + role
        path := "http://" + address + cmd + "/?" + param
        resp, err := http.Get(path)
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        mAllData := make(map[string]interface{})
        err = json.Unmarshal(body, &mAllData)
        if err != nil {
            return err
        }
        data, ok := mAllData["data"]
        if !ok {
          return errors.New("未找到数据")
        }
        return ctx.SendResponse(data)
    }
    return nil
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
