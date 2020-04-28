package gm

import (
    "strconv"
    "sort"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/common"
)

type MailSorter []rpc.MailDataInfo
func (s MailSorter) Len() int {
    return len(s)
}
func (s MailSorter) Less(i, j int) bool {
    return s[i].IMailId < s[j].IMailId
}
func (s MailSorter) Swap(i, j int) {
    tmp := s[i].Copy()
    s[i] = *(s[j].Copy())
    s[j] = *tmp
}

func MailList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    mailPrx := new(rpc.MailService)
    comm.StringToProxy("aqua.MailServer.MailServiceObj", mailPrx)

    var vMail []rpc.MailDataInfo
    ret, err := mailPrx.GetAllMail(&vMail)
    if err := checkRet(ret, err); err != nil {
        return err
    }
    sort.Sort(sort.Reverse(MailSorter(vMail)))

    vPage := common.GetPage(vMail, page, limit)
    return ctx.SendArray(vPage, len(vMail))
}

func MailAdd(c echo.Context) error {
    ctx := c.(*mid.Context)

    return ctx.SendResponse("添加registry成功")
}

func MailDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    /*
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "registry不存在")
    }

    queryPrx := new(rpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    for _, id := range ids {
        tmp := strings.Split(id, "$")
        if len(tmp) != 3 {
            return ctx.SendError(-1, "参数非法")
        }
        ret, err := queryPrx.RemoveEndpoint(tmp[0], tmp[1], tmp[2])
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }
    */

    return ctx.SendResponse("删除registry成功")
}
