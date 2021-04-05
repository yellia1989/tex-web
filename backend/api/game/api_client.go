package game

import (
    "sort"
    "strings"
    "strconv"
    "encoding/base64"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
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
    TmpTimeI := common.ParseTimeInLocal("2006-01-02", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("2006-01-02", a[j].ErrTime)

    if TmpTimeI.After(TmpTimeJ) {
        return true
    }

    if TmpTimeI.Equal(TmpTimeJ) && a[i].ClientVersion > a[j].ClientVersion {
        return true
    }

    if TmpTimeI.Equal(TmpTimeJ) && a[i].ClientVersion == a[j].ClientVersion && a[i].ErrTimes > a[j].ErrTimes {
        return true
    }

    return false
}

type errInfo struct {
    ErrTime    string `json:"err_time"`
    ZoneId     uint32 `json:"zone_id"`
    RoleId     uint32 `json:"role_id"`
}

type errInfoBy []errInfo

func (a errInfoBy) Len() int      { return len(a) }
func (a errInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a errInfoBy) Less(i, j int) bool {
    TmpTimeI := common.ParseTimeInLocal("15:04:05", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("15:04:05", a[j].ErrTime)

    if TmpTimeI.After(TmpTimeJ) {
        return true
    }

    if TmpTimeI.Equal(TmpTimeI) && a[i].ZoneId < a[j].ZoneId {
        return true
    }

    return false
}

func ErrInfo(c echo.Context) error {
    ctx := c.(*mid.Context)
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb

    sql := "SELECT timeymd, client_version, stack, stackmd5, count(*) as count  FROM client_error "
    sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "'"
    sql += "GROUP BY timeymd, client_version, stackmd5"

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slSimpleErrInfo := make([]errSimpleInfo, 0)
    for rows.Next() {
        var r errSimpleInfo
        if err := rows.Scan(&r.ErrTime, &r.ClientVersion, &r.ErrMessage, &r.ErrMessageMd5, &r.ErrTimes); err != nil {
            return err
        }

        decodeBytes, _ := base64.StdEncoding.DecodeString(r.ErrMessage)
        r.ErrMessage = string(decodeBytes)
        r.ErrMessage = strings.Replace(r.ErrMessage, "\n", "<br>", -1)

        slSimpleErrInfo = append(slSimpleErrInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(errSimpleInfoBy(slSimpleErrInfo))

    vSimpleErrInfo := common.GetPage(slSimpleErrInfo, page, limit)

    return ctx.SendArray(vSimpleErrInfo, len(slSimpleErrInfo))
}

func ErrDetail(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrInfoMd5 := ctx.QueryParam("ErrInfo")
    sErrTime := ctx.QueryParam("ErrTime")
    sClientVersion := ctx.QueryParam("ClientVersion")

    if sErrInfoMd5 == "" || sErrTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb
    sql := "SELECT timehms, zoneid, roleid FROM client_error "
    sql += "WHERE STR_TO_DATE(timeymd, '%Y-%m-%d') = STR_TO_DATE('" + sErrTime + "', '%Y-%m-%d')  AND stackmd5 = '" + sErrInfoMd5 + "'" 
    sql += "AND client_version = '"+ sClientVersion + "'"

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slErrInfo := make([]errInfo, 0)

    for rows.Next() {
        var r errInfo
        if err := rows.Scan(&r.ErrTime, &r.ZoneId, &r.RoleId); err != nil {
            return err
        }

        slErrInfo = append(slErrInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(errInfoBy(slErrInfo))

    return ctx.SendArray(slErrInfo, len(slErrInfo))
}
