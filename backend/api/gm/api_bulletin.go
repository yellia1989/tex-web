package gm

import (
    "sort"
    "strings"
    "strconv"
    "encoding/json"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type BulletinSort []rpc.BulletinDataInfo
func (s BulletinSort) Len() int {
    return len(s)
}
func (s BulletinSort) Less(i, j int) bool {
    return s[i].IBulletinId < s[j].IBulletinId
}
func (s BulletinSort) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func BulletinList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    comm := cfg.Comm

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

    var vBulletin []rpc.BulletinDataInfo
    ret, err := bulletinPrx.GetAllBulletin(&vBulletin)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    sort.Sort(sort.Reverse(BulletinSort(vBulletin)))

    vPage := common.GetPage(vBulletin, page, limit)
    return ctx.SendArray(vPage, len(vBulletin))
}

func BulletinAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    bulletin := rpc.NewBulletinDataInfo()
    if err := ctx.Bind(bulletin); err != nil {
        return err
    }

    sTitle := ctx.FormValue("sTitle")
    sContent := ctx.FormValue("sContent")
    sHtmlContent := ctx.FormValue("sHtmlContent")
    sBeginTime := ctx.FormValue("sBeginTime")
    sEndTime := ctx.FormValue("sEndTime")

    if sTitle == "" || sContent == "" || sHtmlContent == "" || sBeginTime == "" || sEndTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    sLangContent := ctx.FormValue("sLangContent")
    if sLangContent != "" {
        json.Unmarshal([]byte(sLangContent), &bulletin.MLangContent)
    }

    comm := cfg.Comm

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

    ret, err := bulletinPrx.AddBulletin(*bulletin.Copy())
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("添加公告成功")
}

func BulletinDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := strings.Split(ctx.FormValue("idsStr"), ",")

    if len(ids) == 0 {
        return ctx.SendError(-1, "公告不存在")
    }

    comm := cfg.Comm

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

    for _, id := range ids  {
        id, _ := strconv.ParseUint(id, 10, 32)
        ret, err := bulletinPrx.DelBulletin(uint32(id))
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("删除公告成功")
}

func BulletinUpdate(c echo.Context) error {
    ctx := c.(*mid.Context)

    bulletin := rpc.NewBulletinDataInfo()
    if err := ctx.Bind(bulletin); err != nil {
        return err
    }

    sTitle := ctx.FormValue("sTitle")
    sContent := ctx.FormValue("sContent")
    sHtmlContent := ctx.FormValue("sHtmlContent")
    sBeginTime := ctx.FormValue("sBeginTime")
    sEndTime := ctx.FormValue("sEndTime")

    if sTitle == "" || sContent == "" || sHtmlContent == "" || sBeginTime == "" || sEndTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    sLangContent := ctx.FormValue("sLangContent")
    if sLangContent != "" {
        json.Unmarshal([]byte(sLangContent), &bulletin.MLangContent)
    }

    comm := cfg.Comm

    bulletinPrx := new(rpc.BulletinService)
    comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

    ret, err := bulletinPrx.ModifyBulletin(*bulletin.Copy())
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return ctx.SendResponse("修改公告成功")
}
