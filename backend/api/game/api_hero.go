package game

import (
    "fmt"
	"strconv"
	Sql "database/sql"
	"github.com/labstack/echo/v4"
	mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/log"
)

type herolog struct {
	Id     uint32 `json:"id"`
	Time   string `json:"time"`
	HeroId uint32 `json:"heroid"`
	Star   uint32 `json:"star"`
    Step   uint32 `json:"step"`
    Quality uint32 `json:"quality"`
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
	
    db, err := zoneLogDb(zoneid)

    if err != nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

	sqlcount := "SELECT count(*) as total FROM log_zone_"+zoneid+".add_hero"
	sqlcount += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = db.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,time,heroid,star,step,quality,operate as action FROM log_zone_"+zoneid+".add_hero"
	sql += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
    sql += " ORDER BY time desc, _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]herolog, 0)
	for rows.Next() {
		var r herolog
		var star, quality, step Sql.NullInt32
		if err := rows.Scan(&r.Id, &r.Time, &r.HeroId, &star, &step, &quality, &r.Action); err != nil {
			return err
		}
		r.Star = uint32(star.Int32)
		r.Step = uint32(step.Int32)
		r.Quality = uint32(quality.Int32)
		logs = append(logs, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}
