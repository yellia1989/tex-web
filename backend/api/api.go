package api

import (
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-web/backend/api/game"
    "github.com/yellia1989/tex-web/backend/api/gm"
    "github.com/yellia1989/tex-web/backend/api/stat"
    "github.com/yellia1989/tex-web/backend/api/sys"
)

func RegisterHandler(group *echo.Group) {
    group.POST("/login", sys.UserLogin) // 登陆

    group.GET("/menu/list", sys.MenuList)               // 菜单列表
    group.POST("/menu/update", sys.MenuUpdate)          // 更新菜单
    group.GET("/user/list", sys.UserList)               // 用户列表
    group.POST("/user/add", sys.UserAdd)                // 用户增加
    group.POST("/user/del", sys.UserDel)                // 用户删除
    group.POST("/user/update/role", sys.UserUpdateRole) // 用户角色变更
    group.POST("/user/update", sys.UserUpdate)          // 用户更新
    group.POST("/user/pwd", sys.UserPwd)                // 用户改密码
    group.GET("/perm/list", sys.PermList)               // 权限列表
    group.POST("/perm/add", sys.PermAdd)                // 权限增加
    group.POST("/perm/del", sys.PermDel)                // 权限删除
    group.POST("/perm/update", sys.PermUpdate)          // 权限编辑
    group.GET("/role/list", sys.RoleList)               // 角色列表
    group.POST("/role/add", sys.RoleAdd)                // 角色增加
    group.POST("/role/del", sys.RoleDel)                // 角色删除
    group.POST("/role/update", sys.RoleUpdate)          // 角色更新
    group.GET("/log/list", sys.LogList)                 // 系统日志

    group.POST("/gm/game/cmd", gm.GameCmd)               // 执行gm命令
    group.GET("/gm/zone/simplelist", gm.ZoneSimpleList)  // 获取分区列表
    group.GET("/gm/zone/list", gm.ZoneList)              // 获取分区列表
    group.POST("/gm/zone/add", gm.ZoneAdd)               // 增加新分区
    group.POST("/gm/zone/del", gm.ZoneDel)               // 删除分区
    group.POST("/gm/zone/update", gm.ZoneUpdate)         // 更新分区
    group.POST("/gm/zone/version", gm.ZoneUpdateVersion) // 批量更新分区版本号

    group.GET("/gm/channel/list", gm.ChannelList)      // 获取渠道列表
    group.POST("/gm/channel/add", gm.ChannelAdd)       // 增加新渠道
    group.POST("/gm/channel/del", gm.ChannelDel)       // 删除渠道
    group.POST("/gm/channel/update", gm.ChannelUpdate) // 更新渠道

    group.GET("/gm/registry/list", gm.RegistryList) // 获取registry列表
    group.POST("/gm/registry/add", gm.RegistryAdd)  // 增加registry
    group.POST("/gm/registry/del", gm.RegistryDel)  // 删除registry

    group.GET("/gm/mail/list", gm.MailList)          // 获取邮件列表
    group.POST("/gm/mail/testsend", gm.MailTestSend) // 发送测试邮件
    group.POST("/gm/mail/send", gm.MailSend)         // 发送邮件
    group.POST("/gm/mail/upload", gm.MailUpload)     // 上传玩家列表
    group.POST("/gm/mail/del", gm.MailDel)           // 删除邮件
    group.POST("/gm/mail/send2", gm.MailSend2)       // 发送邮件

    group.GET("/gm/item/list", gm.ItemList) // 获取道具列表

    group.GET("/gm/bulletin/list", gm.BulletinList)      // 获取公告列表
    group.POST("/gm/bulletin/add", gm.BulletinAdd)       // 增加公告
    group.POST("/gm/bulletin/del", gm.BulletinDel)       // 删除公告
    group.POST("/gm/bulletin/update", gm.BulletinUpdate) // 更新公告

    group.GET("/gm/notice/list", gm.NoticeList)      // 获取跑马灯列表
    group.POST("/gm/notice/add", gm.NoticeAdd)       // 增加跑马灯
    group.POST("/gm/notice/del", gm.NoticeDel)       // 删除跑马灯
    group.POST("/gm/notice/update", gm.NoticeUpdate) // 更新跑马灯

    group.GET("/gm/cdk/list", gm.CDKList)      // 获取cdk列表
    group.POST("/gm/cdk/add", gm.CDKAdd)       // 增加cdk
    group.POST("/gm/cdk/update", gm.CDKUpdate) // 更新cdk
    group.POST("/gm/cdk/export", gm.CDKExport) // 导出cdk

    group.GET("/gm/whitelist/list", gm.WhiteList)        // 获取白名单列表
    group.POST("/gm/whitelist/add", gm.WhiteAdd)         // 增加白名单用户
    group.POST("/gm/whitelist/del", gm.WhiteDel)         // 删除白名单用户
    group.POST("/gm/whitelist/replace", gm.WhiteReplace) // 覆盖白名单用户
    group.GET("/gm/whitelist/tmplist", gm.TmpWhiteList)  // 获取临时白名单列表
    group.POST("/gm/whitelist/addtmp", gm.WhiteAddTmp)   // 增加临时白名单用户
    group.POST("/gm/whitelist/deltmp", gm.WhiteDelTmp)   // 删除白名单用户

    group.POST("/gm/dirty/test", gm.DirtyTest) // 屏蔽字测试

    group.GET("/gm/activity/list", gm.ActivityList)             // 获取活动列表
    group.POST("/gm/activity/add", gm.ActivityAdd)              // 增加活动
    group.POST("/gm/activity/edit", gm.ActivityEdit)            // 编辑活动
    group.POST("/gm/activity/del", gm.ActivityDel)              // 删除活动
    group.POST("/gm/activity/import", gm.ActivityImport)        // 批量导入活动
    group.POST("/gm/activity/lock", gm.ActivityLock)            // 活动解锁
    group.GET("/gm/activity/onlineZone", gm.ActivityOnlineZone) // 查询活动生效分区

    group.GET("/gm/iap/list", gm.IAPList)          // 获取补单商品
    group.POST("/gm/iap/recharge", gm.IAPRecharge) // 补单
    group.POST("/gm/ban/speak", gm.BanSpeak)       // 禁言
    group.POST("/gm/ban/login", gm.BanLogin)       // 禁止登陆

    group.GET("/gm/map/list", gm.MapList)  // 地图列表
    group.POST("/gm/map/add", gm.MapAdd)   // 地图增加
    group.POST("/gm/map/edit", gm.MapEdit) // 地图编辑
    group.POST("/gm/map/del", gm.MapDel)   // 地图删除

    group.GET("/gm/push/list", gm.PushList)          // 推送任务列表
    group.POST("/gm/push/testsend", gm.PushTestSend) // 增加测试推送任务
    group.POST("/gm/push/send", gm.PushSend)         // 增加推送任务
    group.POST("/gm/push/pause", gm.PushPause)       // 推送任务暂停

    group.GET("/gm/welfare/tasklist", gm.WelfareTaskList)      // 福利任务列表
    group.POST("/gm/welfare/taskadd", gm.WelfareTaskAdd)       // 新增一个福利任务
    group.POST("/gm/welfare/taskpause", gm.WelfareTaskPause)   // 暂停福利任务
    group.POST("/gm/welfare/taskresume", gm.WelfareTaskResume) // 恢复福利任务
    group.POST("/gm/welfare/taskupdate", gm.WelfareTaskUpdate) // 更新福利任务
    group.POST("/gm/welfare/taskdel", gm.WelfareTaskDel)       // 删除福利任务
    group.GET("/gm/welfare/rolelist", gm.WelfareRoleList)      // 玩家福利

    group.GET("/game/role/list", game.RoleList)     // 玩家列表
    group.GET("/game/coin/addlog", game.CoinAddLog) // 金币日志
    group.GET("/game/coin/sublog", game.CoinSubLog)
    group.GET("/game/diamond/addlog", game.DiamondAddLog) // 钻石日志
    group.GET("/game/diamond/sublog", game.DiamondSubLog)
    group.GET("/game/hero/addlog", game.HeroAddLog) // 英雄日志
    group.GET("/game/item/addlog", game.ItemAddLog) // 道具日志
    group.GET("/game/item/sublog", game.ItemSubLog)
    group.GET("/game/mail/sendlog", game.MailSendLog) // 邮件日志
    group.GET("/game/mail/revlog", game.MailRevLog)
    group.GET("/game/mail/dellog", game.MailDelLog)

    group.GET("/game/real/online", game.RealOnline)
    group.GET("/game/real/newadd", game.RealNewadd)
    group.GET("/game/real/income", game.RealIncome)
    group.GET("/game/real/stageverify", game.RealStageVerify)
    group.GET("/game/real/fightverify", game.RealFightVerify)
    group.GET("/game/real/stat", stat.RealStat)
    group.GET("/game/real/map", gm.RealMap)
    group.POST("/game/real/mapobj", gm.RealMapObj)
    group.GET("/game/online/time", game.OnlineTime) // 在线时间记录
    group.GET("/game/role/detail", game.RoleDeatil) //获取特定玩家数据
    group.GET("/game/recharge/trace", game.RechargeTrace)
    group.GET("/game/recharge/receipt", gm.IAPDetail)
    group.GET("/game/client_err/err_info", game.ErrInfo)                 // 客户端报错信息
    group.GET("/game/client_err/err_detail", game.ErrDetail)             // 错误信息详情
    group.POST("/game/client_err/dispose", game.ErrDispose)              // 客户端错误开始处理
    group.GET("/game/client_err/dispose_info", game.DisposeList)         // 获取处理情况列表
    group.POST("/game/client_err/add_dispose_note", game.AddDisposeNote) // 增加错误处理备注
    group.POST("/game/client_err/dispose_finish", game.FinishDispose)    // 错误处理完成

    group.GET("/game/chat/getnewest", game.ChatGetNewest)
    group.GET("/game/chat/gethistory", game.ChatGetHistory)
    group.GET("/game/chat/getnewestmask", game.ChatGetMaskNewest)
    group.GET("/game/chat/getmaskhistory", game.ChatGetMaskLogs)
    group.GET("/game/chat/getmaskword", game.ChatGetMaskWord)
    group.POST("/game/chat/setmaskword", game.ChatSetMaskWord)

    group.GET("/stat/all/list", stat.AllList)
    group.GET("/stat/all/ltv", stat.LtvList)
    group.GET("/stat/newadd/list", stat.NewaddList)
    group.GET("/stat/remain/list", stat.RemainList)
    group.GET("/stat/remain/loss", stat.LossList)
    group.GET("/stat/recharge/list", stat.RechargeList)
    group.GET("/stat/recharge/track", stat.RechargeTrack)
    group.GET("/stat/zone/list", stat.ZoneList)
    group.GET("/stat/date/marklist", stat.MarkList)

    group.GET("/gm/res/list", gm.ResControlList)     // 资源监控列表
    group.GET("/gm/res/actionlist", gm.ActionList)   // 资源获取途径监控列表
    group.POST("/gm/res/add", gm.ActionAdd)          // 增加资源监控项
    group.POST("/gm/res/edit", gm.ActionEdit)        // 编辑资源监控项
    group.POST("/gm/res/del", gm.ActionDel)          // 删除资源监控项
    group.GET("/gm/res_err/err_info", gm.ResErrInfo) // 资源获取途径异常信息
    group.GET("/gm/res_err/err_detail", gm.ResErrDetail)
    group.POST("/gm/res/add_res_control", gm.ResAppendResControl) // 增加资源获取途径
    group.POST("/gm/res/add_res_action", gm.ResAppendAction)      // 增加获取途径项
    group.GET("/gm/res_num_err/err_info", gm.ResNumErrInfo)       // 资源数量获取异常异常信息
    group.GET("/gm/res_num_err/err_detail", gm.ResNumErrDetail)

    group.GET("/game/fight-verify/err_info", game.FightErrInfo) // 战斗验证失败日志列表
    group.POST("/game/fight-verify/export-report", game.FightExportReport) // 导出战斗日志
    group.POST("/game/fight-verify/export-log", game.FightExportLog) // 导出战斗日志

    group.GET("/public/gm/server/get_list", gm.GetServerList) // 获取所有地区的服务器信息
    group.GET("/public/gm/zone/get_list", gm.GetZoneList) // 获取可选的服务器列表
    group.GET("/public/gm/dump_role", gm.DumpRole) // 复制玩家数据
    group.POST("/public/gm/load_role", gm.LoadRole) // 粘贴玩家数据
}
