package game

import (
	"strconv"

    Sql "database/sql"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _herolog struct {
	Id     uint32 `json:"id"`
	Time   string `json:"time"`
	HeroId uint32 `json:"heroid"`
	Level  uint32 `json:"level"`
	Star   uint32 `json:"star"`
	Action string `json:"action"`
}

func HeroAddLog(c echo.Context) error {
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

	sqlcount := "SELECT count(*) as total FROM add_hero"
	sqlcount += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = tx.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,time,heroid,hero_level,hero_star,operate as action FROM add_hero"
	sql += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Error(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_herolog, 0)
	for rows.Next() {
		var r _herolog
        var star, level Sql.NullInt32
		if err := rows.Scan(&r.Id, &r.Time, &r.HeroId, &level, &star, &r.Action); err != nil {
			return err
		}
        r.Level = uint32(level.Int32)
        r.Star = uint32(star.Int32)
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
