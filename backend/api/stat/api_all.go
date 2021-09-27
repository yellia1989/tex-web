package stat

import (
    "time"
    "sort"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-go/tools/util"
)

type all struct {
    Statymd string `json:"statymd"`
    Zonename string `json:"zone_name"`
    Zoneopenday uint16 `json:"zone_openday"`
    AccountnumTotal uint32 `json:"accountnum_total"`
    RolenumTotal uint32 `json:"rolenum_total"`
    Newadd uint32 `json:"newadd"`
    Active uint32 `json:"active"`
    LoginTimes uint32 `json:"login_times"`
    RgeRoleNum uint32 `json:"rge_rolenum"`
    RgeTotal uint32 `json:"rge_total"`
}

type allVal struct {
    newadd uint32
    active uint32
    loginTimes uint32
    rgeRoles map[uint32]bool
    rgeTotal uint32
}

func AllList(c echo.Context) error {
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

    m := make(map[dateZoneKey]*allVal)
    vzoneid2 := make([]int,0)

    // 获取指定日期总注册
    accountNumTotal, err := getAccountUtilDate(endDate.ID, accountCond)
    if err != nil {
        return err
    }

    // 获取指定日期范围内的创建的账号
    accounts, err := getAccountByDate(startDate.ID, endDate.ID, accountCond)
    if err != nil {
        return err
    }

    // 获取指定日期总创角
    roleTotals, err := getRoleNumUtilDate(vzoneid, endDate.ID, roleCond)
    if err != nil {
        return err
    }

    // 获取指定日期范围内创建的角色
    roles, err := getRoleNumByDate(vzoneid, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    for k, r := range roles {
        v,ok := m[k]
        if !ok {
            v = &allVal{}
            v.rgeRoles = make(map[uint32]bool)
            m[k] = v
        }

        v.newadd = r

        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }

    // 获取指定日期范围内的活跃
    actives, err := getActiveByDate(vzoneid, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    for k, a := range actives {
        v,ok := m[k]
        if !ok {
            v = &allVal{}
            v.rgeRoles = make(map[uint32]bool)
            m[k] = v
        }

        v.active = a

        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }

    // 获取指定日期范围内的登陆次数
    loginTimes, err := getLoginTimesByDate(vzoneid, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    for k, lt := range loginTimes {
        v,ok := m[k]
        if !ok {
            v = &allVal{}
            v.rgeRoles = make(map[uint32]bool)
            m[k] = v
        }

        v.loginTimes = lt

        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }

    // 获取日期范围内充值玩家
    rgeRecords, err := getRgeRecordByDate(vzoneid, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    for k, records := range rgeRecords {
        v,ok := m[k]
        if !ok {
            v = &allVal{}
            v.rgeRoles = make(map[uint32]bool)
            m[k] = v
        }

        for _,r := range *records {
            v.rgeRoles[r.accountidFk] = true
            v.rgeTotal += r.money
        }

        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(vzoneid2)))

    alls := make([]*all,0)
    for t := endDate.ID; t >= startDate.ID; t-- {
        d := getDate(t)
        for i := 0; i < len(vzoneid2); i++ {
            zoneid := uint32(vzoneid2[i])
            k := dateZoneKey{t, zoneid}
            z := getZone(zoneid)

            v, ok := m[k]
            if !ok {
                continue
            }

            roleNumTotal := roleTotals[zoneid]
            roleTotals[zoneid] -= v.newadd

            a := all{d.Time, z.Name, z.OpenDay, accountNumTotal, roleNumTotal, v.newadd, v.active, v.loginTimes, uint32(len(v.rgeRoles)), v.rgeTotal}
            alls = append(alls, &a)
        }

        // 某一天的新注册有可能是0
        accountNum, _ := accounts[t]
        accountNumTotal -= accountNum
    }

    return ctx.SendArray(common.GetPage(alls, page, limit), len(alls))
}

func RealStat(c echo.Context) error {
    ctx := c.(*mid.Context)
    now := getDateByTime(time.Now())

    data := make(map[string]uint32, 0)

    var rgeToday uint32 // 今日充值
    var rgeNewRoleNumToday uint32 // 今日新增充值人数
    rgeRolesToday := make(map[uint32]bool) // 今日充值人数
    rgeRecords, err := getRgeRecordByDate(getAllRealZoneId(), now.ID, now.ID, roleCond)
    if err != nil {
        return err
    }
    for _, records := range rgeRecords {
        for _, r := range *records {
            rgeToday += r.money
            rgeRolesToday[r.accountidFk] = true
            if r.first {
                rgeNewRoleNumToday += 1
            }
        }
    }
    data["rgeToday"] = rgeToday
    data["rgeRoleNumToday"] = uint32(len(rgeRolesToday))
    data["rgeNewRoleNumToday"] = rgeNewRoleNumToday

    var rgeTotal uint32 // 累计充值
    var rgeRoleNumTotal uint32 // 累计充值人数
    rgeTotals, err := getRgeUtilDate(getAllRealZoneId(), now.ID, roleCond)
    if err != nil {
        return err
    }
    for _, rge := range rgeTotals {
        rgeTotal += rge.total
        rgeRoleNumTotal += rge.rolenum
    }
    data["rgeTotal"] = rgeTotal
    data["rgeRoleNumTotal"] = rgeRoleNumTotal

    var activeToday uint32 // 今日活跃
    actives, err := getActiveByDate(nil, now.ID, now.ID, roleCond)
    if err != nil {
        return err
    }
    for _, v := range actives {
        activeToday += v
    }
    data["activeToday"] = activeToday

    var newaddToday uint32 // 今日新增
    rolesToday, err := getRoleNumByDate(nil, now.ID, now.ID, roleCond)
    if err != nil {
        return err
    }
    for _, v := range rolesToday {
        newaddToday += v
    }
    data["newaddToday"] = newaddToday

    var newaddTotal uint32 // 累计新增
    newaddTotal, err = getRoleNumUtilNow(roleCond)
    if err != nil {
        return err
    }
    data["newaddTotal"] = newaddTotal

    var accountToday uint32 // 今日注册
    accountsToday, err := getAccountByDate(now.ID, now.ID, accountCond)
    if err != nil {
        return err
    }
    for _, v := range accountsToday {
        accountToday += v
    }
    data["accountToday"] = accountToday

    var accountTotal uint32 // 累计注册
    accountTotal, err = getAccountUtilNow(accountCond)
    if err != nil {
        return err
    }
    data["accountTotal"] = accountTotal

    return ctx.SendResponse(data)
}

type ltv struct {
    Zonename string `json:"zone_name"`
    Zoneopenday uint16 `json:"zone_openday"`
    RoleNum uint32 `json:"rolenum"`
    RgeRoleNum uint32 `json:"rge_rolenum"`
    RgeTotal uint32 `json:"rge_total"`
    RgeTotal3   uint32  `json:"rge3"`
    RgeTotal7   uint32  `json:"rge7"`
    RgeTotal15   uint32  `json:"rge15"`
    RgeTotal30   uint32  `json:"rge30"`
    RgeTotal60   uint32  `json:"rge60"`
    RgeTotal90   uint32  `json:"rge90"`
    RgeTotal120   uint32  `json:"rge120"`
    RgeTotal150   uint32  `json:"rge150"`
    RgeTotal180   uint32  `json:"rge180"`
    RoleNum3   uint32  `json:"rolenum3"`
    RoleNum7   uint32  `json:"rolenum7"`
    RoleNum15   uint32  `json:"rolenum15"`
    RoleNum30   uint32  `json:"rolenum30"`
    RoleNum60   uint32  `json:"rolenum60"`
    RoleNum90   uint32  `json:"rolenum90"`
    RoleNum120   uint32  `json:"rolenum120"`
    RoleNum150   uint32  `json:"rolenum150"`
    RoleNum180   uint32  `json:"rolenum180"`
}

type ltvVal struct {
    roleNum uint32
    rgeRoleNum uint32
    rgeTotal uint32
    rgeTotal3 uint32
    rgeTotal7 uint32
    rgeTotal15 uint32
    rgeTotal30 uint32
    rgeTotal60 uint32
    rgeTotal90 uint32
    rgeTotal120 uint32
    rgeTotal150 uint32
    rgeTotal180 uint32
    roleNum3 uint32
    roleNum7 uint32
    roleNum15 uint32
    roleNum30 uint32
    roleNum60 uint32
    roleNum90 uint32
    roleNum120 uint32
    roleNum150 uint32
    roleNum180 uint32
}

func LtvList(c echo.Context) error {
    ctx := c.(*mid.Context)
    vzoneid := common.Atou32v(ctx.QueryParam("zoneid"), ",")
    page := common.Atoi(ctx.QueryParam("page"))
    limit := common.Atoi(ctx.QueryParam("limit"))

    now := getDateByTime(time.Now())

    m := make(map[uint32]*ltvVal)

    vzoneid2 := make([]int,0)

    // 获取分区每日累计充值
    rgeDates, err := getRgeByDate(vzoneid, now.ID, roleCond)
    if err != nil {
        return err
    }
    for k, rge := range rgeDates {
        v, ok := m[k.zoneidFk]
        if !ok {
            v = &ltvVal{}
            m[k.zoneidFk] = v
        }
        z := getZone(k.zoneidFk)
        d := k.dateFk - z.openTimeID + 1
        if d <= 3 {
            v.rgeTotal3 += rge.total
        }
        if d <= 7 {
            v.rgeTotal7 += rge.total
        }
        if d <= 15 {
            v.rgeTotal15 += rge.total
        }
        if d <= 30 {
            v.rgeTotal30 += rge.total
        }
        if d <= 60 {
            v.rgeTotal60 += rge.total
        }
        if d <= 90 {
            v.rgeTotal90 += rge.total
        }
        if d <= 120 {
            v.rgeTotal120 += rge.total
        }
        if d <= 150 {
            v.rgeTotal150 += rge.total
        }
        if d <= 180 {
            v.rgeTotal180 += rge.total
        }
        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }

    // 获取分区累计充值
    rges, err := getRgeUtilDate(vzoneid, now.ID, roleCond)
    if err != nil {
        return err
    }
    for zoneid, rge := range rges {
        v, ok := m[zoneid]
        if !ok {
            v = &ltvVal{}
            m[zoneid] = v
        }

        v.rgeTotal = rge.total
        v.rgeRoleNum = rge.rolenum
        if (!util.Contain(vzoneid2, int(zoneid))) {
            vzoneid2 = append(vzoneid2, int(zoneid))
        }
    }

    // 获取分区累计总创角
    roles, err := getRoleNumUtilDate(vzoneid, now.ID, roleCond)
    if err != nil {
        return err
    }
    for zoneid, r := range roles {
        v, ok := m[zoneid]
        if !ok {
            v = &ltvVal{}
            m[zoneid] = v
        }
        v.roleNum = r
        if (!util.Contain(vzoneid2, int(zoneid))) {
            vzoneid2 = append(vzoneid2, int(zoneid))
        }
    }

    // 获取分区每日创角
    roleDates, err := getRoleNumByDate(vzoneid, 0, now.ID, roleCond)
    for k, r := range roleDates {
        v, ok := m[k.zoneidFk]
        if !ok {
            v = &ltvVal{}
            m[k.zoneidFk] = v
        }

        z := getZone(k.zoneidFk)
        d :=  k.dateFk - z.openTimeID + 1
        if d <= 3 {
            v.roleNum3 += r
        }
        if d <= 7 {
            v.roleNum7 += r
        }
        if d <= 15 {
            v.roleNum15 += r
        }
        if d <= 30 {
            v.roleNum30 += r
        }
        if d <= 60 {
            v.roleNum60 += r
        }
        if d <= 90 {
            v.roleNum90 += r
        }
        if d <= 120 {
            v.roleNum120 += r
        }
        if d <= 150 {
            v.roleNum150 += r
        }
        if d <= 180 {
            v.roleNum180 += r
        }

        if (!util.Contain(vzoneid2, int(k.zoneidFk))) {
            vzoneid2 = append(vzoneid2, int(k.zoneidFk))
        }
    }

    sort.Sort(sort.Reverse(sort.IntSlice(vzoneid2)))

    ltvs := make([]*ltv, 0)
    for _, zoneid := range vzoneid2 {
        v := m[uint32(zoneid)]
        z := getZone(uint32(zoneid))

        l := &ltv{z.Name, z.OpenDay, v.roleNum, v.rgeRoleNum, v.rgeTotal, v.rgeTotal3, v.rgeTotal7, v.rgeTotal15, v.rgeTotal30, v.rgeTotal60, v.rgeTotal90, v.rgeTotal120, v.rgeTotal150, v.rgeTotal180, v.roleNum3, v.roleNum7, v.roleNum15, v.roleNum30, v.roleNum60, v.roleNum90, v.roleNum120, v.roleNum150, v.roleNum180}
        ltvs = append(ltvs, l)
    }

    return ctx.SendArray(common.GetPage(ltvs, page, limit), len(ltvs))
}
