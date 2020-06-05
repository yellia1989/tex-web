package game

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _onlineTime struct {
	TotalTime    string  `json:"onlineTime_total`
	RoleNum      uint32  `json:"onlineTime_roleNum"`
	RoleNumTotal uint32  `json:"onlineTime_roleNum_total`
	RoleNumRate  float32 `json:"onlineTime_roleNum_rate"`
}

func OnlineTime(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

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

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)

	sql := "SELECT roleid,sum(online_time) as online_time FROM logout GROUP BY roleid ORDER BY sum(online_time)"
	sql += " LIMIT " + limitstart + "," + limitrow

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_onlineTime, 0)
	for rows.Next() {
		var r _onlineTime
		/*
			if err := rows.Scan(&r.StageID, &r.StageFirstStartNum, &firstPassNum, &totalPassNum); err != nil {
				return err
			}
			r.StageFirstPassNum = uint32(firstPassNum.Int32)
			r.StageTotalPassNum = uint32(totalPassNum.Int32)

			r.RoleNum = roleNum
			r.StageFirstPassStar1Num = stage2StarNum[r.StageID].star1
			r.StageFirstPassStar2Num = stage2StarNum[r.StageID].star2
			r.StageFirstPassStar3Num = stage2StarNum[r.StageID].star3

			r.StageLossRate = float32(r.StageFirstStartNum-r.StageTotalPassNum) / float32(r.StageFirstStartNum)
			r.StageTotalLossRate = float32(roleNum-r.StageTotalPassNum) / float32(roleNum)
		*/
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
