package gm

import (
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _activityData struct {
	ActivityID    uint32 `json:"activity_id"`
	ActivityType  uint32 `json:"activity_type"`
	ApplyZone     string `json:"apply_zone"`
	ApplyUser     string `json:"apply_user"`
	ConfigureData string `json:"configure_data"`
	ConfigureDesc string `json:"configure_desc"`
	Ts            string `json:"ts"`
}

type _activityType struct {
	Type uint32 `json:"activity_type"`
	Name string `json:"type_name"`
}

func ActivityTypeList(c echo.Context) []_activityData {
	types := []_activityData{
		{1, "模块控制"},
		{2, "个人消耗"},
		{3, "冲榜"},
		{4, "累计充值"},
		{5, "首冲送英雄"},
		{6, "一元购"},
		{7, "七日登陆"},
		{8, "免费福利"},
		{9, "成长基金"},
		{10, "个性化活动"},
		{11, "在线时长奖励"},
		{12, "英雄众筹"},
		{13, "许愿"},
		{14, "英雄兑换"},
		{15, "幸运转盘"},
		{16, "邀请有礼"},
		{17, "vip福利"},
	}

	return types
}

func ActivityList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db := common.GetLogDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE db_zone_global")
	if err != nil {
		return err
	}

	sql := "SELECT * FROM t_activity ORDER BY activity_id DESC;"
	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	c.Logger().Error(sql)

	logs := make([]_activityData, 0)
	for rows.Next() {
		var r _activityData
		if err := rows.Scan(&r.ActivityID, &r.ApplyZone, &r.ApplyUser, &r.ConfigureData, &r.ConfigureDesc, &r.Ts, &r.ActivityType); err != nil {
			return err
		}

		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendArray(logs, len(logs))
}
