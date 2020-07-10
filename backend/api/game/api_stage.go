package game

import (
	"strconv"

	mysql "database/sql"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _stagelog struct {
	Id           uint32 `json:"id"`
	Time         string `json:"time"`
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
	sqlcount += " WHERE roleid=" + roleid + " AND elite=0 AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = tx.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,time,stageid,cost_time,restore_times,revive_times,giveup,win,star,consume_cards,left_cards,first FROM stage_challenge_finish"
	sql += " WHERE roleid=" + roleid + " AND elite=0" + " AND time between '" + startTime + "' AND '" + endTime + "'"
    sql += " ORDER BY _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_stagelog, 0)
	for rows.Next() {
		var r _stagelog
		if err := rows.Scan(&r.Id, &r.Time, &r.StageId, &r.CostTime, &r.RestoreTimes, &r.ReviveTimes, &r.GiveUp, &r.Win, &r.Star, &r.ConsumeCards, &r.LeftCards, &r.First); err != nil {
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
	sqlcount += " WHERE roleid=" + roleid + " AND elite=1 AND time between '" + startTime + "' AND '" + endTime + "'"
	var total int
	err = tx.QueryRow(sqlcount).Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql := "SELECT _rid as id,time,stageid,cost_time,restore_times,revive_times,giveup,win,star,consume_cards,left_cards,first FROM stage_challenge_finish"
	sql += " WHERE roleid=" + roleid + " AND elite=1" + " AND time between '" + startTime + "' AND '" + endTime + "'"
    sql += " ORDER BY _rid desc"
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_stagelog, 0)
	for rows.Next() {
		var r _stagelog
		if err := rows.Scan(&r.Id, &r.Time, &r.StageId, &r.CostTime, &r.RestoreTimes, &r.ReviveTimes, &r.GiveUp, &r.Win, &r.Star, &r.ConsumeCards, &r.LeftCards, &r.First); err != nil {
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

type _stagePass struct {
	StageID                uint32  `json:"stageid"`
	StageFirstStartNum     uint32  `json:"stage_first_start_num"`
	StageFirstPassNum      uint32  `json:"stage_first_pass_num"`
	StageFirstPassStar1Num uint32  `json:"stage_first_pass_star1_num"`
	StageFirstPassStar2Num uint32  `json:"stage_first_pass_star2_num"`
	StageFirstPassStar3Num uint32  `json:"stage_first_pass_star3_num"`
	StageTotalPassNum      uint32  `json:"stage_total_pass_num"`
	StageLossRate          float32 `json:"stage_loss_rate"`
	RoleNum                uint32  `json:"role_num"`
	StageTotalLossRate     float32 `json:"stage_total_loss_rate"`
}
type _star struct {
	star1 uint32
	star2 uint32
	star3 uint32
}

func StagePass(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")

	if zoneid == "" {
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

	var roleNum uint32
	sql := "SELECT count(DISTINCT roleid) as roleNum FROM create_role"
	err = tx.QueryRow(sql).Scan(&roleNum)
	if err != nil {
		return err
	}

	sql = "SELECT stageid, star, count(DISTINCT roleid) AS first_pass_num FROM stage_challenge_finish WHERE `elite` = 0 AND `first` = 1 AND win = 1 GROUP BY stageid,star"
	rows2, err := tx.Query(sql)
	if err != nil {
		return err
	}

	c.Logger().Debug(sql)

	stage2StarNum := make(map[uint32]*_star)
	for rows2.Next() {
		var stage, star, num uint32
		if err := rows2.Scan(&stage, &star, &num); err != nil {
			return nil
		}

		if _, ok := stage2StarNum[stage]; !ok {
			stage2StarNum[stage] = &_star{0, 0, 0}
		}
		switch star {
		case 1:
			stage2StarNum[stage].star1 = num
		case 2:
			stage2StarNum[stage].star2 = num
		case 3:
			stage2StarNum[stage].star3 = num
		}
	}
	rows2.Close()

	sql = "SELECT a.stageid, a.first_start_num, b.first_pass_num, c.total_pass_num FROM "
	sql += "( SELECT stageid, count(*) AS first_start_num FROM stage_challenge_start WHERE `elite` = 0 AND `first` = 1 GROUP BY stageid ) a"
	sql += " LEFT JOIN (SELECT stageid, count(*) AS first_pass_num FROM stage_challenge_finish WHERE `elite` = 0 AND `first` = 1 AND win = 1 GROUP BY stageid ) b ON a.stageid = b.stageid"
	sql += " LEFT JOIN ( SELECT stageid, count( DISTINCT roleid ) AS total_pass_num FROM stage_challenge_finish WHERE `elite` = 0 AND win = 1 GROUP BY stageid ) c ON a.stageid = c.stageid"

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	c.Logger().Debug(sql)

	logs := make([]_stagePass, 0)
	total := 0
	for rows.Next() {
		total++
		var r _stagePass
		var firstPassNum, totalPassNum mysql.NullInt32
		if err := rows.Scan(&r.StageID, &r.StageFirstStartNum, &firstPassNum, &totalPassNum); err != nil {
			return err
		}
		r.StageFirstPassNum = uint32(firstPassNum.Int32)
		r.StageTotalPassNum = uint32(totalPassNum.Int32)

		r.RoleNum = roleNum

		if _, ok := stage2StarNum[r.StageID]; !ok {
			stage2StarNum[r.StageID] = &_star{0, 0, 0}
		}
		r.StageFirstPassStar1Num = stage2StarNum[r.StageID].star1
		r.StageFirstPassStar2Num = stage2StarNum[r.StageID].star2
		r.StageFirstPassStar3Num = stage2StarNum[r.StageID].star3

		r.StageLossRate = float32(r.StageFirstStartNum-r.StageTotalPassNum) / float32(r.StageFirstStartNum)
		r.StageTotalLossRate = float32(roleNum-r.StageTotalPassNum) / float32(roleNum)
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
