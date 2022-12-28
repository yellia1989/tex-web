package gm

import (
    "fmt"
    "time"
    "sync"
    "strings"
    "strconv"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
)

// 缓存分区列表
var zones []rpc.ZoneInfo
var mu sync.Mutex

type _zoneInfo struct {
    rpc.ZoneInfo
    SPublishTime string `json:"sPublishTime"`
}

func ZoneMap() map[uint32]rpc.ZoneInfo {
    mzone := make(map[uint32]rpc.ZoneInfo)

    tmp := updateZoneList(false)
    for _, v := range tmp {
        mzone[v.IZoneId] = *(v.Copy())
    }
    return mzone
}

func updateZoneList(bUpdate bool) ([]rpc.ZoneInfo) {
    mu.Lock()
    if len(zones) == 0 {
        bUpdate = true
    }
    if !bUpdate {
        mu.Unlock()
        return zones
    }

    mu.Unlock()

    comm := cfg.Comm

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(cfg.App+".DirServer.DirServiceObj", dirPrx)

    var tmp []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&tmp)
    if ret == 0 && err == nil {
        mu.Lock()
        zones = tmp[:]
        mu.Unlock()
    }

    return zones
}

func ZoneSimpleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    game := ctx.QueryParam("game")
    all := ctx.QueryParam("all")

    bgame := game != ""
    ball := all != ""

    if game == "" && all == "" {
        bgame = true
    }

    var zones []rpc.ZoneInfo
    if ball {
        zones = append(zones, rpc.ZoneInfo{IZoneId: 99999, SZoneName: "全服"})
    }

    if bgame {
        tmp := updateZoneList(false)
        zones2 := make([]rpc.ZoneInfo, len(tmp), len(tmp))
        copy(zones2, tmp)
        for i,_ := range zones2 {
            zones2[i].SZoneName = fmt.Sprintf("%s(%d)", zones2[i].SZoneName, zones2[i].IZoneId)
        }
        zones = append(zones, zones2...)
    }

    data := make(map[string][]rpc.ZoneInfo,0)
    data["game"] = zones

    return ctx.SendResponse(data)
}

func ZoneList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    tmp := updateZoneList(true)
    zones2 := make([]_zoneInfo, len(tmp))
    for k,v := range tmp {
        zones2[k].IZoneId = v.IZoneId
        zones2[k].SZoneName = v.SZoneName
        zones2[k].SConnServer = v.SConnServer
        zones2[k].SGameServer = v.SGameServer
        zones2[k].IZoneFlag = v.IZoneFlag
        zones2[k].IIsManual = v.IIsManual
        zones2[k].IManualZoneStatus = v.IManualZoneStatus
        zones2[k].IMaxNum = v.IMaxNum
        zones2[k].SPublishTime = common.FormatTimeInLocal("2006-01-02 15:04:05", time.Unix(int64(v.IPublishTime), 0))
        zones2[k].IIsKick = v.IIsKick
        zones2[k].MVersion = v.MVersion
        zones2[k].IMaxOnline = v.IMaxOnline
        zones2[k].ICurNum = v.ICurNum
        zones2[k].ILastReportTime = v.ILastReportTime
        zones2[k].ICurZoneStatus = v.ICurZoneStatus
        zones2[k].ICurOnline = v.ICurOnline
    }

    vPage := common.GetPage(zones2, page, limit)

    return ctx.SendArray(vPage, len(zones2))
}

func ZoneAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    zone := rpc.NewZoneInfo()
    if err := ctx.Bind(zone); err != nil {
        return err
    }

    sPublishTime := c.FormValue("sPublishTime")
    if sPublishTime == "" {
        return ctx.SendError(-1, "开服时间不能为空")
    }
    zone.IPublishTime = uint32(common.ParseTimeInLocal("2006-01-02 15:04:05", sPublishTime).Unix())

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    var channels []rpc.ChannelAddr
    ret, err := loginPrx.GetAllChannel(&channels)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    // 填充每一个渠道的版本号信息
    zone.MVersion = make(map[string]rpc.ZoneVersion)
    for _, v := range channels {
        sResVersion := ctx.FormValue(fmt.Sprintf("s%sResVersion", v.SChannel))
        if (sResVersion == "") {
            return ctx.SendError(-1, "渠道版本号不能为空")
        }
        var ver rpc.ZoneVersion
        ver.SRes = sResVersion
        zone.MVersion[v.SChannel] = ver
    }

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(cfg.App+".DirServer.DirServiceObj", dirPrx)
    ret, err = dirPrx.CreateZone(*zone.Copy())
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("添加分区成功")
}

func ZoneDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "分区不存在")
    }

    comm := cfg.Comm

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(cfg.App+".DirServer.DirServiceObj", dirPrx)

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        ret, err := dirPrx.DeleteZone(uint32(id))
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("删除分区成功")
}

func ZoneUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)

    zone := rpc.NewZoneInfo()
    if err := ctx.Bind(zone); err != nil {
        return err
    }

    sPublishTime := c.FormValue("sPublishTime")
    if sPublishTime == "" {
        return ctx.SendError(-1, "开服时间不能为空")
    }
    zone.IPublishTime = uint32(common.ParseTimeInLocal("2006-01-02 15:04:05", sPublishTime).Unix())

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    var channels []rpc.ChannelAddr
    ret, err := loginPrx.GetAllChannel(&channels)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    // 填充每一个渠道的版本号信息
    zone.MVersion = make(map[string]rpc.ZoneVersion)
    for _, v := range channels {
        sResVersion := ctx.FormValue(fmt.Sprintf("s%sResVersion", v.SChannel))
        if (sResVersion == "") {
            return ctx.SendError(-1, "渠道版本号不能为空")
        }
        var ver rpc.ZoneVersion
        ver.SRes = sResVersion
        zone.MVersion[v.SChannel] = ver
    }

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(cfg.App+".DirServer.DirServiceObj", dirPrx)
    ret, err = dirPrx.ModifyZone(*zone.Copy(), rpc.ZoneModifyInfo{})
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("修改分区成功")
}

func ZoneUpdateVersion(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "分区不存在")
    }

    comm := cfg.Comm

    loginPrx := new(rpc.LoginService)
    comm.StringToProxy(cfg.App+".LoginServer.LoginServiceObj", loginPrx)

    var channels []rpc.ChannelAddr
    ret, err := loginPrx.GetAllChannel(&channels)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    // 填充每一个渠道的版本号信息
    MVersion := make(map[string]rpc.ZoneVersion)
    for _, v := range channels {
        sResVersion := ctx.FormValue(fmt.Sprintf("s%sResVersion", v.SChannel))
        if (sResVersion == "") {
            return ctx.SendError(-1, "渠道版本号不能为空")
        }
        var ver rpc.ZoneVersion
        ver.SRes = sResVersion
        MVersion[v.SChannel] = ver
    }

    iZoneFlag , _ := strconv.ParseUint(ctx.FormValue("iZoneFlag"),10,32)
    iIsManual , _ := strconv.ParseUint(ctx.FormValue("iIsManual"),10,32)
    iManualZoneStatus , _ := strconv.ParseUint(ctx.FormValue("iManualZoneStatus"),10,32)

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(cfg.App+".DirServer.DirServiceObj", dirPrx)

    for _, id := range ids {
        id, _ := strconv.ParseUint(id, 10, 32)
        zone := rpc.NewZoneInfo()
        ret, err := dirPrx.GetZone(uint32(id), zone)
        if err := checkRet(ret, err); err != nil {
            return err
        }
        zone.MVersion = MVersion
        tempZone := *zone.Copy()
        tempZone.IZoneFlag = uint32(iZoneFlag)
        tempZone.IIsManual = uint32(iIsManual)
        tempZone.IManualZoneStatus = uint32(iManualZoneStatus)

        ret, err = dirPrx.ModifyZone(tempZone, rpc.ZoneModifyInfo{})
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("批量修改分区成功")
}
