package stat

import (
    "time"
    "sort"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-go/tools/util"
)

type remain struct {
    StatYmd string `json:"statymd"`
    ZoneName string `json:"zone_name"`
    ZoneOpenDay uint16 `json:"zone_openday"`
    Days []uint32 `json:"days"`
    Newadd uint32 `json:"newadd"`
    RgeTotal7 uint32 `json:"rge_total7"`
    RgeTotal30 uint32 `json:"rge_total30"`
    RgeRolenum uint32 `json:"rge_rolenum"`
    RgeTotal uint32 `json:"rge_total"`
    RgeR uint32 `json:"rge_r"`
}

type remainVal struct {
    roleNum uint32
    rgeTotal uint32
    rgeTotal7 uint32
    rgeTotal30 uint32
    active []uint32
    r uint32
    rgeRoleNum uint32
}

func RemainList(c echo.Context) error {
    ctx := c.(*mid.Context)
    vzoneid := common.Atou32v(ctx.QueryParam("zoneid"), ",")
    page := common.Atoi(ctx.QueryParam("page"))
    limit := common.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    datatype := common.Atoi(ctx.QueryParam("datatype"))

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
    roles, err := getRoles(vzoneid, startDate.ID, endDate.ID, datatype == 2, roleCond)
    if err != nil {
        return err
    }

    m := make(map[dateZoneKey]*remainVal)
    vzoneid2 := make([]int,0)
    for _, r := range roles {
        k := dateZoneKey{r.regDateFk,r.zoneidFk}
        d := nowDate.ID-r.regDateFk+1
        if d > 90 {
            // 只统计90天留存
            d = 90
        }
        v,ok := m[k]
        if !ok {
            v = &remainVal{}
            m[k] = v
            // 根据创角时间距离当前时间来创建90日活跃数组，这样可以节省空间
            v.active = make([]uint32, d, d)
        }

        v.roleNum += 1
        v.rgeTotal += r.rgeTotal
        v.rgeTotal7 += r.rgeTotal7
        v.rgeTotal30 += r.rgeTotal30
        for i := uint32(0); i < d; i++ {
            if r.isActive(i+1) {
                v.active[i] += 1
            }
        }
        if r.isR() {
            v.r += 1
        }
        if r.rgeTotal > 0 {
            v.rgeRoleNum += 1
        }
        if (!util.Contain(vzoneid2, int(r.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(r.zoneidFk))
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(vzoneid2)))

    remains := make([]*remain,0)
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
            r := remain{d.Time, z.Name, z.OpenDay, v.active, v.roleNum, v.rgeTotal7, v.rgeTotal30, v.rgeRoleNum, v.rgeTotal, v.r}
            remains = append(remains, &r)
        }
    }

    return ctx.SendArray(common.GetPage(remains, page, limit), len(remains))
}

type loss struct {
    Statymd string `json:"statymd"`
    Zonename string `json:"zone_name"`
    Zoneopenday uint16 `json:"zone_openday"`
    Newadd uint32   `json:"newadd"`
    RgeRoleNum uint32 `json:"rge_rolenum"`
    WeekActive uint32 `json:"week_active"`
    DWeekActive uint32 `json:"dweek_active"`
    MonthActive uint32 `json:"month_active"`
    PayWeekActive uint32 `json:"pay_week_active"`
    PayDWeekActive uint32 `json:"pay_dweek_active"`
    PayMonthActive uint32 `json:"pay_month_active"`
}

type lossVal struct {
    newadd uint32
    rgeRoleNum uint32
    weekActive uint32
    dweekActive uint32
    monthActive uint32
    payWeekActive uint32
    payDWeekActive uint32
    payMonthActive uint32
}

func LossList(c echo.Context) error {
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

    // 获取创角玩家
    roles, err := getRoles(vzoneid, startDate.ID, endDate.ID, false, roleCond)
    if err != nil {
        return err
    }

    m := make(map[dateZoneKey]*lossVal)
    vzoneid2 := make([]int,0)
    for _, r := range roles {
        k := dateZoneKey{r.regDateFk,r.zoneidFk}
        v,ok := m[k]
        if !ok {
            v = &lossVal{}
            m[k] = v
        }

        v.newadd += 1
        if r.rgeTotal > 0 {
            v.rgeRoleNum += 1
        }
        if r.isWeekActive() {
            if r.rgeTotal > 0 {
                v.payWeekActive += 1
            } else {
                v.weekActive += 1
            }
        }
        if r.isDWeekActive() {
            if r.rgeTotal > 0 {
                v.payDWeekActive += 1
            } else {
                v.dweekActive += 1
            }
        }
        if r.isMonthActive() {
            if r.rgeTotal > 0 {
                v.payMonthActive += 1
            } else {
                v.monthActive += 1
            }
        }
        if (!util.Contain(vzoneid2, int(r.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(r.zoneidFk))
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(vzoneid2)))

    losses := make([]*loss,0)
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
            l := loss{d.Time, z.Name, z.OpenDay, v.newadd, v.rgeRoleNum, v.weekActive, v.dweekActive, v.monthActive, v.payWeekActive, v.payDWeekActive, v.payMonthActive}
            losses = append(losses, &l)
        }
    }

    return ctx.SendArray(common.GetPage(losses, page, limit), len(losses))
}
