package game

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _stagelog struct {
	Id           uint32 `json:"id"`
	StageId      uint32 `json:"stageid"`
	CostTime     uint32 `json:"cost_time"`
	RestoreTimes uint32 `json:"restore_times"`
	ReviveTimes  uint32 `json:"revive_times"`
	GiveUp       uint32 `json:"give_up`
	Win          uint32 `json:"win"`
	Star         uint32 `json:"star"`
	ConsumeCards uint32 `json:"consume_cards"`
	LeftCards    uint32 `json:"left_cards"`
	First        uint32 `json:"first"`
}

func StageAddLog(c echo.Context) error {
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

	sqlcount := "SELECT count(*) as total FROM stage_challenge_finish"
	sqlcount += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = tx.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,stageid,cost_time,restore_times,revive_times,giveup,win,star,consume_cards,left_cards,first FROM stage_challenge_finish"
	sql += " WHERE roleid=" + roleid + " AND elite=0" + " AND time between '" + startTime + "' AND '" + endTime + "'"
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Error(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_stagelog, 0)
	for rows.Next() {
		var r _stagelog
		if err := rows.Scan(&r.Id, &r.StageId, &r.CostTime, &r.RestoreTimes, &r.ReviveTimes, &r.GiveUp, &r.Win, &r.Star, &r.ConsumeCards, &r.LeftCards, &r.First); err != nil {
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

func EliteStageAddLog(c echo.Context) error {
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

	sqlcount := "SELECT count(*) as total FROM stage_challenge_finish"
	sqlcount += " WHERE roleid=" + roleid + " AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = tx.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,stageid,cost_time,restore_times,revive_times,giveup,win,star,consume_cards,left_cards,first FROM stage_challenge_finish"
	sql += " WHERE roleid=" + roleid + " AND elite=1" + " AND time between '" + startTime + "' AND '" + endTime + "'"
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Error(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_stagelog, 0)
	for rows.Next() {
		var r _stagelog
		if err := rows.Scan(&r.Id, &r.StageId, &r.CostTime, &r.RestoreTimes, &r.ReviveTimes, &r.GiveUp, &r.Win, &r.Star, &r.ConsumeCards, &r.LeftCards, &r.First); err != nil {
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
