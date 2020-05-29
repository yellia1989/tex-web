package api

import (
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/api/gm"
    "github.com/yellia1989/tex-web/backend/api/game"
)

func RegisterHandler(group *echo.Group) {
    group.POST("/login", UserLogin)                     // 登陆

    group.GET("/menu/list", MenuList)                   // 菜单列表
    group.POST("/menu/update", MenuUpdate)              // 更新菜单
    group.GET("/user/list", UserList)                   // 用户列表
    group.POST("/user/add", UserAdd)                    // 用户增加
    group.POST("/user/del", UserDel)                    // 用户删除
    group.POST("/user/update/role", UserUpdateRole)     // 用户角色变更
    group.POST("/user/update", UserUpdate)              // 用户更新
    group.POST("/user/pwd", UserPwd)                    // 用户改密码
    group.GET("/perm/list", PermList)                   // 权限列表
    group.POST("/perm/add", PermAdd)                    // 权限增加
    group.POST("/perm/del", PermDel)                    // 权限删除
    group.POST("/perm/update", PermUpdate)              // 权限编辑
    group.GET("/role/list", RoleList)                   // 角色列表
    group.POST("/role/add", RoleAdd)                    // 角色增加
    group.POST("/role/del", RoleDel)                    // 角色删除
    group.POST("/role/update", RoleUpdate)              // 角色更新

    group.POST("/gm/game/cmd", gm.GameCmd)              // 执行gm命令
    group.GET("/gm/zone/simplelist", gm.ZoneSimpleList) // 获取分区列表
    group.GET("/gm/zone/list", gm.ZoneList)             // 获取分区列表
    group.POST("/gm/zone/add", gm.ZoneAdd)              // 增加新分区
    group.POST("/gm/zone/del", gm.ZoneDel)              // 删除分区
    group.POST("/gm/zone/update", gm.ZoneUpdate)        // 更新分区
    group.POST("/gm/zone/version", gm.ZoneUpdateVersion)        // 批量更新分区版本号

    group.GET("/gm/channel/list", gm.ChannelList)             // 获取渠道列表
    group.POST("/gm/channel/add", gm.ChannelAdd)              // 增加新渠道
    group.POST("/gm/channel/del", gm.ChannelDel)              // 删除渠道
    group.POST("/gm/channel/update", gm.ChannelUpdate)        // 更新渠道

    group.GET("/gm/registry/list", gm.RegistryList)           // 获取registry列表
    group.POST("/gm/registry/add", gm.RegistryAdd)            // 增加registry
    group.POST("/gm/registry/del", gm.RegistryDel)            // 删除registry

    group.GET("/gm/mail/list", gm.MailList)                   // 获取邮件列表
    group.POST("/gm/mail/testsend", gm.MailTestSend)          // 发送测试邮件
    group.POST("/gm/mail/send", gm.MailSend)                  // 发送邮件
    group.POST("/gm/mail/upload", gm.MailUpload)              // 上传玩家列表
    group.POST("/gm/mail/del", gm.MailDel)                    // 删除邮件
    group.POST("/gm/mail/send2", gm.MailSend2)                // 发送邮件

    group.GET("/gm/item/list", gm.ItemList)                   // 获取道具列表

	group.GET("/gm/bulletin/list", gm.BulletinList)           // 获取公告列表
	group.POST("/gm/bulletin/add", gm.BulletinAdd)            // 增加公告
	group.POST("/gm/bulletin/del", gm.BulletinDel)       // 删除公告
	group.POST("/gm/bulletin/update", gm.BulletinUpdate) // 更新公告

	group.GET("/gm/notice/list", gm.NoticeList)      // 获取跑马灯列表
	group.POST("/gm/notice/add", gm.NoticeAdd)       // 增加跑马灯
	group.POST("/gm/notice/del", gm.NoticeDel)       // 删除跑马灯
	group.POST("/gm/notice/update", gm.NoticeUpdate) // 更新跑马灯

	group.GET("/gm/cdk/list", gm.CDKList)      // 获取cdk列表
	group.POST("/gm/cdk/add", gm.CDKAdd)       // 增加cdk
	group.POST("/gm/cdk/update", gm.CDKUpdate) // 更新cdk

    group.GET("/gm/whitelist/list", gm.WhiteList)   // 获取白名单列表
    group.POST("/gm/whitelist/add", gm.WhiteAdd)    // 增加白名单用户
    group.POST("/gm/whitelist/del", gm.WhiteDel)    // 删除白名单用户
    group.POST("/gm/whitelist/replace", gm.WhiteReplace)    // 覆盖白名单用户

    group.POST("/gm/dirty/test", gm.DirtyTest)   // 屏蔽字测试

    group.GET("/gm/activity/list", gm.ActivityList)     // 获取活动列表

    group.GET("/gm/iap/list", gm.IAPList)   // 获取补单商品
    group.POST("/gm/iap/recharge", gm.IAPRecharge) // 补单
    group.POST("/gm/ban/speak", gm.BanSpeak) // 禁言
    group.POST("/gm/ban/login", gm.BanLogin) // 禁止登陆

    group.GET("/game/role/list", game.RoleList)
    group.GET("/game/coin/addlog", game.CoinAddLog)
    group.GET("/game/coin/sublog", game.CoinSubLog)
    group.GET("/game/diamond/addlog", game.DiamondAddLog)
    group.GET("/game/diamond/sublog", game.DiamondSubLog)

    group.GET("/game/real/online", game.RealOnline)
    group.GET("/game/real/newadd", game.RealNewadd)
    group.GET("/game/real/income", game.RealIncome)
    group.GET("/game/real/stageverify", game.RealStageVerify)
    group.GET("/game/real/stat", game.RealStat)
}
