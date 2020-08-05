package gm

import (
	"strconv"
	"strings"
	"github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/util"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func PushList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    comm := common.GetLocator()

	pushPrx := new(rpc.PushService)
	comm.StringToProxy(common.GetApp()+".PushServer.PushServiceObj", pushPrx)

	var vTask []rpc.PushTaskInfo
	ret, err := pushPrx.GetAllPushTaskInfo(&vTask)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	vPage := common.GetPage(vTask, page, limit)
	return ctx.SendArray(vPage, len(vTask))
}

func PushTestSend(c echo.Context) error {
	ctx := c.(*mid.Context)

    sTaskName := ctx.FormValue("sTaskName")
    roleid, _ := strconv.Atoi(ctx.FormValue("iRoleId"))
    spayload := ctx.FormValue("payload")

    comm := common.GetLocator()

	pushPrx := new(rpc.PushService)
	comm.StringToProxy(common.GetApp()+".PushServer.PushServiceObj", pushPrx)

    var vTarget []rpc.PushTargetAccountInfo
    var payload rpc.PushPayloadInfo
    var taskid uint32

    var target rpc.PushTargetAccountInfo
    target.IAccountId = uint64(roleid)
    vTarget = append(vTarget, target)
    payload.SUPushPayload = spayload

	ret, err := pushPrx.AddPushTask(vTarget, sTaskName, payload, &taskid)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("添加推送任务成功")
}

func PushSend(c echo.Context) error {
	ctx := c.(*mid.Context)

    sTaskName := ctx.FormValue("sTaskName")
    spayload := ctx.FormValue("payload")
    filename := ctx.FormValue("filepath")

    var vTarget []rpc.PushTargetAccountInfo
    content, err := util.LoadFromFile(filename)
    if err != nil {
        return err
    }
    ids := strings.Split(string(content), "\n")
    for _, v := range ids {
        roleid, _ := strconv.Atoi(v)
        var target rpc.PushTargetAccountInfo
        target.IAccountId = uint64(roleid)
        vTarget = append(vTarget, target)
    }

    var payload rpc.PushPayloadInfo
    payload.SUPushPayload = spayload

    comm := common.GetLocator()

	pushPrx := new(rpc.PushService)
	comm.StringToProxy(common.GetApp()+".PushServer.PushServiceObj", pushPrx)

    var taskid uint32
	ret, err := pushPrx.AddPushTask(vTarget, sTaskName, payload, &taskid)
	if err := checkRet(ret, err); err != nil {
		return err
	}

	return ctx.SendResponse("添加推送任务成功")
}

func PushPause(c echo.Context) error {
	ctx := c.(*mid.Context)

	ids := strings.Split(ctx.FormValue("idsStr"), ",")
	if len(ids) == 0 {
		return ctx.SendError(-1, "任务不存在")
	}
    tmp, _ := strconv.Atoi(ctx.FormValue("pause"))
    pause := false
    if tmp == 1 {
        pause = true
    }

    comm := common.GetLocator()

	pushPrx := new(rpc.PushService)
	comm.StringToProxy(common.GetApp()+".PushServer.PushServiceObj", pushPrx)

	for _, id := range ids {
		id, _ := strconv.ParseUint(id, 10, 32)
		ret, err := pushPrx.PausePushTask(uint32(id), pause)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	}

	return ctx.SendResponse("暂停或回复任务成功")
}
