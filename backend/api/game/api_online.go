package game

import (
	"strconv"

	"sort"
	"time"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _onlineTime struct {
	TotalTime    string  `json:"onlineTime_total"`
	RoleNum      uint32  `json:"onlineTime_roleNum"`
	RoleNumTotal uint32  `json:"onlineTime_roleNum_total"`
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

	data := time.Now().Format("2006-01-02")
	data = "logymd='" + data + "'"
	sql := "SELECT roleid,sum(online_time) as online_time FROM logout "
	sql += "WHERE " + data + " GROUP BY roleid ORDER BY sum(online_time)"
	sql += " LIMIT " + limitstart + "," + limitrow

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	mRole := make(map[uint64]uint32)
	for rows.Next() {
		var roleID uint64
		var totalTime uint32
		if err := rows.Scan(&roleID, &totalTime); err != nil {
			return err
		}

		mRole[roleID] = totalTime
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	roleNum := uint32(len(mRole))

	// 初始化一套时间轴
	times := map[uint32]uint32{300: 0, 600: 0, 1800: 0, 3600: 0, 7200: 0, 3600 * 4: 0, 3600 * 8: 0, 3600 * 12: 0, 3600 * 24: 0, 3600*24 + 1: 0}
	// 统计数据
	for _, t := range mRole {
		switch {
		case t <= 300:
			times[300]++
		case t <= 600:
			times[600]++
		case t <= 1800:
			times[1800]++
		case t <= 3600:
			times[3600]++
		case t <= 7200:
			times[7200]++
		case t <= 3600*4:
			times[3600*4]++
		case t <= 3600*8:
			times[3600*8]++
		case t <= 3600*12:
			times[3600*12]++
		case t <= 3600*24:
			times[3600*24]++
		default:
			times[3600*24+1]++
		}
	}
	tmp := make([]int, 0, 10)
	for k := range times {
		tmp = append(tmp, int(k))
	}
	sort.Ints(tmp)

	logs := make([]_onlineTime, 0)
	for _, t := range tmp {
		r := _onlineTime{RoleNum: times[uint32(t)], RoleNumTotal: roleNum}
		switch {
		case t == 300:
			r.TotalTime = "0-5分钟"
		case t == 600:
			r.TotalTime = "5-10分钟"
		case t == 1800:
			r.TotalTime = "10-30分钟"
		case t == 3600:
			r.TotalTime = "30-60分钟"
		case t == 7200:
			r.TotalTime = "60-120分钟"
		case t == 3600*4:
			r.TotalTime = "2-4小时"
		case t == 3600*8:
			r.TotalTime = "4-8小时"
		case t == 3600*12:
			r.TotalTime = "8-12小时"
		case t == 3600*24:
			r.TotalTime = "12-24小时"
		case t == 3600*24+1:
			r.TotalTime = "24小时以上"
		}
		r.RoleNumRate = float32(r.RoleNum) / float32(r.RoleNumTotal)
		logs = append(logs, r)
	}

	return ctx.SendArray(logs, len(logs))
}
