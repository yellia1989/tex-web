package stat

import (
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type newadd struct {
    Statymd string `json:"statymd"`
    AccountnumTotal uint32 `json:"accountnum_total"`
    RolenumTotal uint32 `json:"rolenum_total"`
    Accountnum uint32 `json:"accountnum"`
    Rolenum uint32 `json:"rolenum"`
}

func NewaddList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page := common.Atoi(ctx.QueryParam("page"))
    limit := common.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    startDate := getDateByString(startTime)
    endDate := getDateByString(endTime)
    if startDate == nil || endDate == nil {
        return ctx.SendError(-1, "请选择想要查找的服务器并选择日期范围")
    }

    // 获取账号总数量
    accountNumTotal, err := getAccountUtilDate(endDate.ID, accountCond)
    if err != nil {
        return err
    }

    // 获取角色总数量
    roleNumTotals, err := getRoleNumUtilDate(nil, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    // 全服总创角
    var roleNumTotal uint32
    for _, v := range roleNumTotals {
        roleNumTotal += v
    }

    // 获取指定日期范围内的创建的账号
    accounts, err := getAccountByDate(startDate.ID, endDate.ID, accountCond)
    if err != nil {
        return err
    }

    // 获取指定日期范围内创建的角色
    roles, err := getRoleNumByDate(nil, startDate.ID, endDate.ID, roleCond)
    if err != nil {
        return err
    }
    date2Roles := make(map[uint32]uint32)
    for k, r := range roles {
        date2Roles[k.dateFk] += r
    }

    newadds := make([]*newadd,0)
    for t := endDate.ID; t >= startDate.ID; t-- {
        d := getDate(t)

        accountnum, ok := accounts[d.ID]
        if !ok {
            continue
        }

        var nd newadd
        nd.Statymd = d.Time
        nd.Accountnum = accountnum
        nd.Rolenum,_ = date2Roles[d.ID]

        nd.RolenumTotal = roleNumTotal
        nd.AccountnumTotal = accountNumTotal

        accountNumTotal -= nd.Accountnum
        roleNumTotal -= nd.Rolenum

        newadds = append(newadds, &nd)
    }

    return ctx.SendArray(common.GetPage(newadds, page, limit), len(newadds))
}
