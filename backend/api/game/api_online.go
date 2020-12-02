package game

import (
    "fmt"
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type onlineTime struct {
	TotalTime    string  `json:"onlineTime_total"`
	RoleNum      uint32  `json:"onlineTime_roleNum"`
	RoleNumTotal uint32  `json:"onlineTime_roleNum_total"`
	RoleNumRate  float32 `json:"onlineTime_roleNum_rate"`
}

func OnlineTime(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")

	if zoneid == "" || startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

    db, err := zoneLogDb(zoneid)
    defer db.Close()

    if err != nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }

	sql := "SELECT roleid,sum(online_time) as online_time FROM logout "
	sql += "WHERE time BETWEEN '"+startTime+"' AND '"+endTime+"' GROUP BY roleid ORDER BY sum(online_time)"

	c.Logger().Debug(sql)

	rows, err := db.Query(sql)
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

	roleNum := uint32(len(mRole))
	logs := make([]onlineTime, 0)
    if roleNum == 0 {
	    return ctx.SendArray(logs, len(logs))
    }

	// 初始化一套时间轴
	time := []uint32{300, 600, 1800, 3600, 7200, 3600 * 4, 3600 * 8, 3600 * 12, 3600 * 24, 3600*24 + 1}
	num := make([]uint32, len(time))
	// 统计数据
	for _, t := range mRole {
		switch {
		case t <= 300:
			num[0]++
		case t <= 600:
			num[1]++
		case t <= 1800:
			num[2]++
		case t <= 3600:
			num[3]++
		case t <= 7200:
			num[4]++
		case t <= 3600*4:
			num[5]++
		case t <= 3600*8:
			num[6]++
		case t <= 3600*12:
			num[7]++
		case t <= 3600*24:
			num[8]++
		default:
			num[9]++
		}
	}

	for k, t := range time {
		r := onlineTime{RoleNum: num[k], RoleNumTotal: roleNum}
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
