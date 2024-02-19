package gm

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-go/tools/util"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"github.com/yellia1989/tex-web/backend/service"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
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

	comm := cfg.Comm

	mailPrx := new(rpc.MailService)
	comm.StringToProxy(service.GetMailServiceName(), mailPrx)

	var vMail []rpc.MailDataInfo
	ret, err := mailPrx.GetAllMail(&vMail)
	if err := checkRet(ret, err); err != nil {
		return err
	}
	sort.Sort(sort.Reverse(MailSorter(vMail)))

	vPage := common.GetPage(vMail, page, limit)
	return ctx.SendArray(vPage, len(vMail))
}

func MailTestSend(c echo.Context) error {
	ctx := c.(*mid.Context)

	sTitle := ctx.FormValue("sTitle")
	sContent := ctx.FormValue("sContent")
	iCoin, _ := strconv.Atoi(ctx.FormValue("iCoin"))
	iDiamond, _ := strconv.Atoi(ctx.FormValue("iDiamond"))
	iZoneId, _ := strconv.Atoi(ctx.FormValue("iZoneId"))
	iRoleId, _ := strconv.ParseUint(ctx.FormValue("iRoleId"), 10, 64)
	itemstr := ctx.FormValue("items")

	if sTitle == "" || sContent == "" {
		return ctx.SendError(-1, "参数非法")
	}

	m := rpc.NewMailDataInfo()
	m.STitle = sTitle
	m.SContent = sContent
	m.ICoin = uint32(iCoin)
	m.IDiamond = uint32(iDiamond)
	if itemstr != "" {
		item1 := strings.Split(itemstr, ";")
		for _, v := range item1 {
			item2 := strings.SplitN(v, ",", 2)
			id, _ := strconv.ParseUint(item2[0], 10, 32)
			num, _ := strconv.ParseUint(item2[1], 10, 32)
			m.VItems = append(m.VItems, rpc.CmdIDNum{IId: uint32(id), INum: uint32(num)})
		}
	}
	m.VSendZoneIds = append(m.VSendZoneIds, uint32(iZoneId))
	m.VToUser = append(m.VToUser, iRoleId)

	sLangContent := ctx.FormValue("sLangContent")
	if sLangContent != "" {
		json.Unmarshal([]byte(sLangContent), &m.MLangContent)
	}

	comm := cfg.Comm

	mailPrx := new(rpc.MailService)
	comm.StringToProxy(service.GetMailServiceName(), mailPrx)

	ret, err := mailPrx.AddMail(*m.Copy())
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("发送测试邮件成功")
}

func MailSend(c echo.Context) error {
	ctx := c.(*mid.Context)

	sTitle := ctx.FormValue("sTitle")
	sContent := ctx.FormValue("sContent")
	iCoin, _ := strconv.Atoi(ctx.FormValue("iCoin"))
	iDiamond, _ := strconv.Atoi(ctx.FormValue("iDiamond"))
	itemstr := ctx.FormValue("items")

	if sTitle == "" || sContent == "" {
		return ctx.SendError(-1, "参数非法")
	}

	m := rpc.NewMailDataInfo()
	m.STitle = sTitle
	m.SContent = sContent
	m.ICoin = uint32(iCoin)
	m.IDiamond = uint32(iDiamond)
	if itemstr != "" {
		item1 := strings.Split(itemstr, ";")
		for _, v := range item1 {
			item2 := strings.SplitN(v, ",", 2)
			id, _ := strconv.ParseUint(item2[0], 10, 32)
			num, _ := strconv.ParseUint(item2[1], 10, 32)
			m.VItems = append(m.VItems, rpc.CmdIDNum{IId: uint32(id), INum: uint32(num)})
		}
	}

	sLangContent := ctx.FormValue("sLangContent")
	if sLangContent != "" {
		json.Unmarshal([]byte(sLangContent), &m.MLangContent)
	}

	comm := cfg.Comm

	mailPrx := new(rpc.MailService)
	comm.StringToProxy(service.GetMailServiceName(), mailPrx)

	zonestr := ctx.FormValue("zoneids")
	filename := ctx.FormValue("filepath")
	if zonestr == "" && filename == "" {
		return ctx.SendError(-1, "参数非法")
	}

	if zonestr != "" {
		zone1 := strings.Split(zonestr, ",")
		// 指定分区发送
		for _, v := range zone1 {
			id, _ := strconv.Atoi(v)
			m.VSendZoneIds = append(m.VSendZoneIds, uint32(id))
		}

		ret, err := mailPrx.AddMail(*m.Copy())
		if err := checkRet(ret, err); err != nil {
			return err
		}
	} else {
		// 指定玩家发送
		content, err := util.LoadFromFile(filename)
		if err != nil {
			return err
		}
		zoneid2roles := make(map[int][]uint64)
		role1 := strings.Split(string(content), "\n")
		for _, ids := range role1 {
			tmp := strings.Fields(ids)
			if len(tmp) != 2 {
				// 格式错误直接忽略
				continue
			}
			zoneid, _ := strconv.Atoi(tmp[0])
			roleid, _ := strconv.ParseUint(tmp[1], 10, 64)
			zoneid2roles[zoneid] = append(zoneid2roles[zoneid], roleid)
		}

		var mails []rpc.MailDataInfo
		for k, v := range zoneid2roles {
			m.VSendZoneIds = m.VSendZoneIds[:0]
			m.VSendZoneIds = append(m.VSendZoneIds, uint32(k))
			m.VToUser = v[:]
			mails = append(mails, *m.Copy())
		}

		ret, err := mailPrx.AddMails(mails)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	return ctx.SendResponse("发送邮件成功")
}

func MailUpload(c echo.Context) error {
	ctx := c.(*mid.Context)

	fh, err := ctx.FormFile("mailids")
	if err != nil {
		return err
	}
	f, err := fh.Open()
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	filename := "data/upload/" + time.Now().Format("20060102150405")
	err = util.SaveToFile(filename, content, false)
	if err != nil {
		return err
	}

	return ctx.SendResponse(filename)
}

func MailDel(c echo.Context) error {
	ctx := c.(*mid.Context)
	ids := strings.Split(ctx.FormValue("idsStr"), ",")
	if len(ids) == 0 {
		return ctx.SendError(-1, "邮件不存在")
	}

	comm := cfg.Comm

	mailPrx := new(rpc.MailService)
	comm.StringToProxy(service.GetMailServiceName(), mailPrx)

	for _, id := range ids {
		id, _ := strconv.ParseUint(id, 10, 32)
		ret, err := mailPrx.DelMail(uint32(id))
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	return ctx.SendResponse("删除邮件成功")
}

func MailSend2(c echo.Context) error {
	ctx := c.(*mid.Context)

	sTitle := ctx.FormValue("sTitle")
	sContent := ctx.FormValue("sContent")

	if sTitle == "" || sContent == "" {
		return ctx.SendError(-1, "参数非法")
	}

	m := rpc.NewMailDataInfo()
	m.STitle = sTitle
	m.SContent = sContent

	sLangContent := ctx.FormValue("sLangContent")
	if sLangContent != "" {
		json.Unmarshal([]byte(sLangContent), &m.MLangContent)
	}

	comm := cfg.Comm

	mailPrx := new(rpc.MailService)
	comm.StringToProxy(service.GetMailServiceName(), mailPrx)

	// 指定玩家发送
	filename := ctx.FormValue("filepath")
	if filename == "" {
		return ctx.SendError(-1, "参数非法")
	}
	content, err := util.LoadFromFile(filename)
	if err != nil {
		return err
	}
	role1 := strings.Split(string(content), "\n")
	for _, ids := range role1 {
		tmp := strings.Fields(ids)
		if len(tmp) != 3 {
			// 格式错误直接忽略
			continue
		}
		zoneid := 0
		if zoneid, err = strconv.Atoi(tmp[0]); err != nil {
			return fmt.Errorf("非法的分区id")
		}
		roleid := uint64(0)
		if roleid, err = strconv.ParseUint(tmp[1], 10, 64); err != nil {
			return fmt.Errorf("非法的角色id")
		}
		item1 := strings.Split(tmp[2], ";")
		m.VItems = m.VItems[:0]
		for _, v := range item1 {
			item2 := strings.SplitN(v, ",", 2)
			id, _ := strconv.ParseUint(item2[0], 10, 32)
			num, _ := strconv.ParseUint(item2[1], 10, 32)
			m.VItems = append(m.VItems, rpc.CmdIDNum{IId: uint32(id), INum: uint32(num)})
		}

		m.VSendZoneIds = m.VSendZoneIds[:0]
		m.VSendZoneIds = append(m.VSendZoneIds, uint32(zoneid))
		m.VToUser = m.VToUser[:0]
		m.VToUser = append(m.VToUser, uint64(roleid))

		ret, err := mailPrx.AddMail(*m.Copy())
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	return ctx.SendResponse("发送邮件成功")
}
