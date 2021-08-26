package gm

import (
    "sort"
	"strconv"
	"strings"
    "encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/common"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type NoticeSort []rpc.NoticeDataInfo
func (s NoticeSort) Len() int {
    return len(s)
}
func (s NoticeSort) Less(i, j int) bool {
    return s[i].INoticeId < s[j].INoticeId
}
func (s NoticeSort) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func NoticeList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    comm := cfg.Comm

	bulletinPrx := new(rpc.BulletinService)
	comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

	var vNotice []rpc.NoticeDataInfo
	ret, err := bulletinPrx.GetAllNotice(&vNotice)
	if err := checkRet(ret, err); err != nil {
		return err
	}

    sort.Sort(sort.Reverse(NoticeSort(vNotice)))

	vPage := common.GetPage(vNotice, page, limit)
	return ctx.SendArray(vPage, len(vNotice))
}

func NoticeAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	notice := rpc.NoticeDataInfo{}
	if err := ctx.Bind(&notice); err != nil {
		return err
	}

	sBeginTime := ctx.FormValue("sBeginTime")
	sEndTime := ctx.FormValue("sEndTime")

    szoneid := ctx.FormValue("zoneid")
    if szoneid != "" {
	    zones := strings.Split(szoneid, ",")
	    for _, zone := range zones {
	    	zone, _ := strconv.ParseUint(zone, 10, 32)
	    	notice.VZoneId = append(notice.VZoneId, uint32(zone))
	    }
    }

	if sBeginTime == "" || sEndTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

    sLangContent := ctx.FormValue("sLangContent")
    if sLangContent != "" {
        json.Unmarshal([]byte(sLangContent), &notice.MLangContent)
    }

    comm := cfg.Comm

	bulletinPrx := new(rpc.BulletinService)
	comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

	ret, err := bulletinPrx.AddNotice(notice)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("添加跑马灯成功")
}

func NoticeDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	ids := strings.Split(ctx.FormValue("idsStr"), ",")

	if len(ids) == 0 {
		return ctx.SendError(-1, "跑马灯不存在")
	}

    comm := cfg.Comm

	bulletinPrx := new(rpc.BulletinService)
	comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

	for _, id := range ids {
		id, _ := strconv.ParseUint(id, 10, 32)
		ret, err := bulletinPrx.DelNotice(uint32(id))
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	return ctx.SendResponse("删除跑马灯成功")
}

func NoticeUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

	notice := rpc.NoticeDataInfo{}
	if err := ctx.Bind(&notice); err != nil {
		return err
	}

    comm := cfg.Comm

	bulletinPrx := new(rpc.BulletinService)
	comm.StringToProxy(cfg.App+".BulletinServer.BulletinServiceObj", bulletinPrx)

	ret, err := bulletinPrx.ModifyNotice(notice)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("修改跑马灯成功")
}
