package stat

import (
    "sort"
    "time"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-go/tools/util"
)

type recharge struct {
    Statymd string `json:"statymd"`
    Zonename string `json:"zone_name"`
    Zoneopenday uint16 `json:"zone_openday"`
    Active uint32 `json:"active"`
    RgeRolenum int `json:"rge_rolenum"`
    RgeTotal uint32 `json:"rge_total"`
    RgeNewRolenum uint32 `json:"rge_newrolenum"`
}

type rechargeVal struct {
    active uint32
    rgeRoles map[uint32]bool
    rgeTotal uint32
    rgeNewRoleNum uint32
}

func RechargeList(c echo.Context) error {
    ctx := c.(*mid.Context)
    vzoneid := common.Atou32v(ctx.QueryParam("zoneid"), ",")
    page := common.Atoi(ctx.QueryParam("page"))
    limit := common.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    startDate := getDateByString(startTime)
    endDate := getDateByString(endTime)
    if startDate == nil || endDate == nil {
        return ctx.SendError(-1, "请指定日期范围")
    }

    if len(vzoneid) == 0 {
        return ctx.SendError(-1, "请指定分区")
    }

    // 获取日期范围内活跃玩家
    dateZoneActives, err := getActiveByDate(vzoneid, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }

    // 获取日期范围内充值玩家
    rgeRecords, err := getRgeRecordByDate(vzoneid, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    
    vzoneid2 := make([]int,0)
    m := make(map[dateZoneKey]*rechargeVal)
    for k,active := range dateZoneActives {
        im, ok := m[k]
        if !ok {
            im = &rechargeVal{}
            im.rgeRoles = make(map[uint32]bool)
            m[k] = im
        }
        im.active = active
        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }
    for k,records := range rgeRecords {
        im, ok := m[k]
        if !ok {
            im = &rechargeVal{}
            im.rgeRoles = make(map[uint32]bool)
            m[k] = im
        }
        for _,r := range *records {
            im.rgeRoles[r.accountidFk] = true
            im.rgeTotal += r.money
            if r.first {
                im.rgeNewRoleNum += 1
            }
        }
        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(vzoneid2)))

    recharges := make([]*recharge,0)
    for t := endDate.ID; t >= startDate.ID; t-- {
        d := getDate(t)
        for i := 0; i < len(vzoneid2); i++ {
            zoneid := uint32(vzoneid2[i])
            k := dateZoneKey{t, zoneid}
            v, ok := m[k]
            if !ok {
                continue
            }
            z := getZone(zoneid)
            im := recharge{d.Time, z.Name, z.OpenDay, v.active, len(v.rgeRoles), v.rgeTotal, v.rgeNewRoleNum}
            recharges = append(recharges, &im)
        }
    }

    return ctx.SendArray(common.GetPage(recharges, page, limit), len(recharges))
}

type rechargeTrack struct {
    Statymd string `json:"statymd"`
    Zonename string `json:"zone_name"`
    Zoneopenday uint16 `json:"zone_openday"`
    Newadd uint32 `json:"newadd"`
    RgeTotal uint32 `json:"rge_total"`
    RgeTotal3 uint32 `json:"rge_total3"`
    RgeTotal7 uint32 `json:"rge_total7"`
    RgeTotal14 uint32 `json:"rge_total14"`
    RgeTotal30 uint32 `json:"rge_total30"`
    Rges []uint32 `json:"days"`
}

type rechargeTrackVal struct {
    newadd uint32
    rgeTotal uint32
    rgeTotal3 uint32
    rgeTotal7 uint32
    rgeTotal14 uint32
    rgeTotal30 uint32
    rges []uint32
}

func RechargeTrack(c echo.Context) error {
    ctx := c.(*mid.Context)
    vzoneid := common.Atou32v(ctx.QueryParam("zoneid"), ",")
    page := common.Atoi(ctx.QueryParam("page"))
    limit := common.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    startDate := getDateByString(startTime)
    endDate := getDateByString(endTime)
    if startDate == nil || endDate == nil {
        return ctx.SendError(-1, "请指定日期范围")
    }

    if len(vzoneid) == 0 {
        return ctx.SendError(-1, "请指定分区")
    }

    nowDate := getDateByTime(time.Now())

    // 获取创角玩家
    roles, err := getRoles(vzoneid, startDate.ID, endDate.ID, false, roleCond)
    if err != nil {
        return err
    }

    m := make(map[dateZoneKey]*rechargeTrackVal)
    vzoneid2 := make([]int,0)
    for _, r := range roles {
        k := dateZoneKey{r.regDateFk,r.zoneidFk}
        d := nowDate.ID-r.regDateFk+1
        if d > 90 {
            // 只统计90天充值
            d = 90
        }
        v,ok := m[k]
        if !ok {
            v = &rechargeTrackVal{}
            m[k] = v
            // 根据创角时间距离当前时间来创建90日活跃数组，这样可以节省空间
            v.rges = make([]uint32, d, d)
        }

        v.newadd += 1
        v.rgeTotal += r.rgeTotal
        v.rgeTotal3 += r.rgeTotal3
        v.rgeTotal7 += r.rgeTotal7
        v.rgeTotal14 += r.rgeTotal14
        v.rgeTotal30 += r.rgeTotal30
        for i := uint32(0); i < d; i++ {
            v.rges[i] += r.getRge(i+1)
        }
        if (!util.Contain(vzoneid2, int(r.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(r.zoneidFk))
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(vzoneid2)))

    rechargeTracks := make([]*rechargeTrack,0)
    for t := endDate.ID; t >= startDate.ID; t-- {
        d := getDate(t)
        for i := 0; i < len(vzoneid2); i++ {
            zoneid := uint32(vzoneid2[i])
            k := dateZoneKey{t, zoneid}
            v, ok := m[k]
            if !ok {
                continue
            }
            z := getZone(zoneid)
            imt := rechargeTrack{d.Time, z.Name, z.OpenDay, v.newadd, v.rgeTotal, v.rgeTotal3, v.rgeTotal7, v.rgeTotal14, v.rgeTotal30, make([]uint32, len(v.rges), len(v.rges))}
            copy(imt.Rges, v.rges)
            rechargeTracks = append(rechargeTracks, &imt)
        }
    }

    return ctx.SendArray(common.GetPage(rechargeTracks, page, limit), len(rechargeTracks))
}
