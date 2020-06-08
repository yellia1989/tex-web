package gm

import (
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _activityData struct {
	ActivityID    uint32 `json:"activity_id"`
	ApplyZone     string `json:"apply_zone"`
	ApplyUser     string `json:"apply_user"`
	ConfigureData string `json:"configure_data"`
	ConfigureDesc string `json:"configure_desc"`
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

	sql := "SELECT * FROM t_activity where ORDER BY activity_id DESC;"
	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	c.Logger().Error(sql)

	logs := make([]_activityData, 0)
	for rows.Next() {
		var r _activityData
		if err := rows.Scan(&r.ActivityID, &r.ApplyZone, &r.ApplyUser, &r.ConfigureData, &r.ConfigureDesc); err != nil {
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

	return ctx.SendResponse(logs, len(logs))
}
