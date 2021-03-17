package game

import (
	"sort"

	"strings"

	"github.com/labstack/echo"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-web/backend/cfg"
	common "github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type errSimpleInfo struct {
	ErrTime       string `json:"err_time"`
	ErrMessage    string `json:"err_info"`
	ErrTimes      uint32 `json:"err_times"`
	ClientVersion string `json:"client_version"`
	ErrMessageMd5 string `json:"err_info_md5"`
}

type errSimpleInfoBy []errSimpleInfo

func (a errSimpleInfoBy) Len() int      { return len(a) }
func (a errSimpleInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a errSimpleInfoBy) Less(i, j int) bool {
	if a[i].ErrTime > a[j].ErrTime {
		return true
	}

	if a[i].ClientVersion > a[j].ClientVersion {
		return true
	}

	return a[i].ErrTimes > a[j].ErrTimes
}

type errInfo struct {
	ErrTime    string `json:"err_time"`
	ErrMessage string `json:"err_info"`
	ZoneId     uint32 `json:"zone_id"`
	RoleId     uint32 `json:"role_id"`
}

type errInfoBy []errInfo

func (a errInfoBy) Len() int      { return len(a) }
func (a errInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a errInfoBy) Less(i, j int) bool {
	if a[i].ErrTime > a[j].ErrTime {
		return true
	}

	if a[i].ErrMessage > a[j].ErrMessage {
		return true
	}

	return a[i].ZoneId < a[j].ZoneId
}

func ErrInfo(c echo.Context) error {
	ctx := c.(*mid.Context)
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")

	if startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.LogDb

	timeFormatDay := "2006-01-02"
	tStartTimeDay := common.ParseTimeInLocal(timeFormatDay, startTime)
	tEndTime := common.ParseTimeInLocal(timeFormatDay, endTime)

	sql := "SELECT timeymd, client_version, stack, stackmd5 FROM client_err "
	sql += "WHERE time BETWEEN '" + tStartTime.Format(timeFormatDay) + "' AND '" + tEndTime.Format(timeFormatDay) + "'"

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	slSimpleErrInfo := make([]errSimpleInfo, 0)
	for rows.Next() {
		var r errSimpleInfo
		if err := rows.Scan(&r.ErrTime, &r.ClientVersion, &r.ErrMessage, &r.ErrMessageMd5); err != nil {
			return err
		}

		r.ErrMessage = strings.Replace(r.ErrMessage, "\n", "<br>", -1)

		r.ErrTimes += 1

		bFind := false
		for k, v := range slSimpleErrInfo {
			if len(slSimpleErrInfo) != 0 && v.ErrTime == r.ErrTime && v.ErrMessageMd5 == r.ErrMessageMd5 {
				slSimpleErrInfo[k].ErrTimes += 1
				bFind = true
			}
		}

		if !bFind {
			slSimpleErrInfo = append(slSimpleErrInfo, r)
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	sort.Sort(errSimpleInfoBy(slSimpleErrInfo))

	return ctx.SendArray(slSimpleErrInfo, len(slSimpleErrInfo))
}

func ErrDetail(c echo.Context) error {
	ctx := c.(*mid.Context)
	sErrInfoMd5 := ctx.QueryParam("ErrInfo")
	sErrTime := ctx.QueryParam("ErrTime")

	if sErrInfo == "" || sErrTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.LogDb
	sql := "SELECT timehms, message, zoneid, roleid FROM client_err "
	sql += "WHERE str_to_date(timeymd, '%Y-%m-%d') = str_to_date('" + sErrTime + "')  AND stackmd5 = '" + sErrInfoMd5 + "'"

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	slErrInfo := make([]errInfo, 0)

	for rows.Next() {
		var r errInfo
		if err := rows.Scan(&r.ErrTime, &r.ErrMessage, &r.ZoneId, &r.RoleId); err != nil {
			return err
		}
		r.ErrMessage = strings.Replace(r.ErrMessage, "\n", "<br>", -1)

		slErrInfo = append(slErrInfo, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	sort.Sort(errInfoBy(slErrInfo))

	return ctx.SendArray(slErrInfo, len(slErrInfo))
}
