package game

import (
    "fmt"
	"github.com/labstack/echo/v4"
	mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/log"
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

    if err != nil {
        return ctx.SendError(-1, fmt.Sprintf("连接数据库失败: %s", err.Error()))
    }
    defer db.Close()

	sql := "SELECT roleid,sum(online_time) as online_time FROM logout "
	sql += "WHERE time BETWEEN '"+startTime+"' AND '"+endTime+"' GROUP BY roleid ORDER BY sum(online_time)"

	log.Infof("sql: %s", sql)

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
	time := []uint32{300, 600, 900, 1200, 1500, 1800, 2100, 2400, 2700, 3000, 3300, 3600, 3900, 4200, 4500, 4800, 5100, 5400, 5700, 6000, 6300, 6600, 6900, 7200, 3600 * 4, 3600 * 8, 3600 * 12, 3600 * 24, 3600*24 + 1}
	num := make([]uint32, len(time))
	// 统计数据
	for _, t := range mRole {
        i := 0
        for ; i < len(time); i += 1 {
            if t <= time[i] {
                break
            }
        }
        if i == len(time) {
            i = len(time) - 1
        }
        num[i] += 1;
	}

	for k, t := range time {
		r := onlineTime{RoleNum: num[k], RoleNumTotal: roleNum}
		switch {
		case t == 300:
			r.TotalTime = "0-5分钟"
		case t == 600:
			r.TotalTime = "5-10分钟"
		case t == 900:
			r.TotalTime = "10-15分钟"
		case t == 1200:
			r.TotalTime = "15-20分钟"
		case t == 1500:
			r.TotalTime = "20-25分钟"
		case t == 1800:
			r.TotalTime = "25-30分钟"
		case t == 2100:
			r.TotalTime = "30-35分钟"
		case t == 2400:
			r.TotalTime = "35-40分钟"
		case t == 2700:
			r.TotalTime = "40-45分钟"
		case t == 3000:
			r.TotalTime = "45-50分钟"
		case t == 3300:
			r.TotalTime = "50-55分钟"
		case t == 3600:
			r.TotalTime = "55-60分钟"
		case t == 3900:
			r.TotalTime = "60-65分钟"
		case t == 4200:
			r.TotalTime = "65-70分钟"
		case t == 4500:
			r.TotalTime = "70-75分钟"
		case t == 4800:
			r.TotalTime = "75-80分钟"
		case t == 5100:
			r.TotalTime = "80-85分钟"
		case t == 5400:
			r.TotalTime = "85-90分钟"
		case t == 5700:
			r.TotalTime = "90-95分钟"
		case t == 6000:
			r.TotalTime = "95-100分钟"
		case t == 6300:
			r.TotalTime = "100-105分钟"
		case t == 6600:
			r.TotalTime = "105-110分钟"
		case t == 6900:
			r.TotalTime = "110-115分钟"
		case t == 7200:
			r.TotalTime = "115-120分钟"
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
