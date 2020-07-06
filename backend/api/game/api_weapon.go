package game

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _weaponlog struct {
	Id       uint32 `json:"id"`
	Time     string `json:"time"`
	WeaponId uint32 `json:"weaponId"`
	Action   string `json:"action"`
}

func WeaponAddLog(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	roleid := ctx.QueryParam("roleid")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")

	if zoneid == "" || roleid == "" || startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := common.GetLogDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE log_zone_" + zoneid)
	if err != nil {
		return err
	}

	sqlcount := "SELECT count(*) as total FROM add_weapon"
	sqlcount += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = tx.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,time,weaponId,operate as action FROM add_weapon"
	sql += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
    sql += " ORDER BY _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_weaponlog, 0)
	for rows.Next() {
		var r _weaponlog
		if err := rows.Scan(&r.Id, &r.Time, &r.WeaponId, &r.Action); err != nil {
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

	return ctx.SendArray(logs, total)
}
