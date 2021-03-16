package game

import (
	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
    common "github.com/yellia1989/tex-web/backend/common"
)

type errSimpleInfo struct {
	ErrTime    string `json:"err_time"`
	ErrMessage string `json:"err_info"`
	ErrTimes   uint32 `json:"err_times"`
}

type errSimpleInfoBy []errSimpleInfo

func (a errSimpleInfoBy) Len() int      { return len(a) }
func (a errSimpleInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a errSimpleInfoBy) Less(i, j int) bool {

	return a[i].ErrTime > a[j].ErrTime

	return a[i].ErrTimes > a[j].ErrTimes

	return a[i].ErrMessage < a[j].ErrMessage
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

	return a[i].ErrTime > a[j].ErrTime

	return a[i].ErrMessage < a[j].ErrMessage

    return a[i].ZoneId < a[j].ZoneId
}

func ErrInfo(c echo.Context) error {
	ctx := c.(*mid.Context)
    startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")

    if startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

    db, err := zoneLogDb(zoneid)

    sql := "select time, message FROM client_err "
    sql += "WHERE time BETWEEN '"+startTime"' AND '"+endTime"'"
    
    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()


	simpleErrInfo := make([]errSimpleInfo, 0)
    for rows.Next() {
        var r errSimpleInfo
        var string sErrtime
        if err := row.Scan(&ErrTime, &r.ErrMessage); err != nil {
            return err
        }
        // 日期按天排序
        timeFormat := "2006-01-02 15:04:05"
        tErrtime := common.ParseTimeInLocal(timeFormat, sErrTime)
        r.ErrTime = tErrtime.Format("2006-01-02")    
        
        ++r.ErrTimes
        for k,v := range simpleErrInfo {
            if (v.ErrTime == r.ErrTime && v.ErrMessage == r.ErrMessage) {
                ++simpleErrInfo[k].Errtimes
            }
        }
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(errSimpleInfo(simpleErrInfo))
    
	return ctx.SendArray(simpleErrInfo, len(simpleErrInfo))
}

func ErrDetail(c echo.Context) error {
    ctx := c.(*mid.Context)
    logs := make([]errSimpleInfo, 0)

    return ctx.SendArray(logs, len(logs))
}