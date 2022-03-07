package game

import (
    "encoding/base64"
    "sort"
    "strconv"
    "strings"

    "github.com/labstack/echo/v4"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/common"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

type errSimpleInfo struct {
    ErrTime       string `json:"err_time"`
    ErrMessage    string `json:"err_info"`
    ErrTimes      uint32 `json:"err_times"`
    ClientVersion string `json:"client_version"`
    ErrMessageMd5 string `json:"err_info_md5"`
    Status        uint32 `json:"status"`
    ErrId         uint32 `json:"err_id"`
}

type errSimpleInfoBy []errSimpleInfo

func (a errSimpleInfoBy) Len() int      { return len(a) }
func (a errSimpleInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a errSimpleInfoBy) Less(i, j int) bool {
    TmpTimeI := common.ParseTimeInLocal("2006-01-02", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("2006-01-02", a[j].ErrTime)

    if a[i].ClientVersion != a[j].ClientVersion {
        return a[i].ClientVersion > a[j].ClientVersion
    }

    if !TmpTimeI.Equal(TmpTimeJ) {
        return TmpTimeI.After(TmpTimeJ)
    }

    return a[i].ErrTimes > a[j].ErrTimes
}

type errInfo struct {
    ErrTime string `json:"err_time"`
    ZoneId  uint32 `json:"zone_id"`
    RoleId  uint32 `json:"role_id"`
}

type errInfoBy []errInfo

func (a errInfoBy) Len() int      { return len(a) }
func (a errInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a errInfoBy) Less(i, j int) bool {
    TmpTimeI := common.ParseTimeInLocal("15:04:05", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("15:04:05", a[j].ErrTime)

    if !TmpTimeI.Equal(TmpTimeJ) {
        return TmpTimeI.After(TmpTimeJ)
    }

    return a[i].ZoneId < a[j].ZoneId
}

type disposeInfo struct {
    ErrMessage    string `json:"err_info"`
    ErrMessageMd5 string `json:"err_info_md5"`
    ClientVersion string `json:"client_version"`
    Status        uint32 `json:"status"`
    Note          string `json:"err_note"`
}

type disposeInfoBy []disposeInfo

func (a disposeInfoBy) Len() int      { return len(a) }
func (a disposeInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a disposeInfoBy) Less(i, j int) bool {
    if a[i].Status !=  a[j].Status {
        return a[i].Status < a[j].Status
    } 
    return a[i].ClientVersion > a[j].ClientVersion
}

func ErrInfo(c echo.Context) error {
    ctx := c.(*mid.Context)
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    clientVersion := ctx.QueryParam("client_version")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb

    sql := "SELECT _rid, timeymd, client_version, stack, stackmd5, count(*) as count  FROM client_error "
    where := "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "'"
    if clientVersion != "" {
        where += " AND client_version IN ('" + clientVersion + "')"
    }
    sql += where
    sql += "GROUP BY client_version, stackmd5"

    var total int
    row2, err2 := db.Query("SELECT count(*) as total FROM ("+sql+") a")
    if err2 != nil {
        return err2
    }
    for row2.Next() {
        if err := row2.Scan(&total); err != nil {
            return err
        }
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slSimpleErrInfo := make([]errSimpleInfo, 0)
    for rows.Next() {
        var r errSimpleInfo
        if err := rows.Scan(&r.ErrId, &r.ErrTime, &r.ClientVersion, &r.ErrMessage, &r.ErrMessageMd5, &r.ErrTimes); err != nil {
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

    db = cfg.StatDb
    sql1 := "SELECT status, client_version, stackmd5 FROM client_dispose "
    rows1, err1 := db.Query(sql1)
    if err1 != nil {
        return err1
    }
    defer rows1.Close()

    slDisposeInfo := make([]disposeInfo, 0)
    for rows1.Next() {
        var r disposeInfo
        if err2 := rows1.Scan(&r.Status, &r.ClientVersion, &r.ErrMessageMd5); err != nil {
            return err2
        }

        slDisposeInfo = append(slDisposeInfo, r)
    }

    k := 0
    vErrId := make([]uint32, 0)
    for _, errInfo := range slSimpleErrInfo {
        bFind := false
        for _, disInfo := range slDisposeInfo {
            if errInfo.ErrMessageMd5 == disInfo.ErrMessageMd5 && errInfo.ClientVersion == disInfo.ClientVersion {
                vErrId = append(vErrId, errInfo.ErrId)
                total--
                bFind = true               
            }
        }
        if !bFind {
            slSimpleErrInfo[k] = errInfo
            k++
        }
    }
    slSimpleErrInfo = slSimpleErrInfo[:k]

    sErrId := ""
    for k, v := range vErrId {
        sErrId = sErrId + "'" + strconv.Itoa(int(v)) + "'"
        if k != (len(vErrId) - 1) {
            sErrId += ", "
        }
    }

    if  sErrId != "" {
        delClientError(sErrId)
    }

    sort.Sort(errSimpleInfoBy(slSimpleErrInfo))

    return ctx.SendArray(slSimpleErrInfo, total)
}

func ErrDetail(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrInfoMd5 := ctx.QueryParam("ErrInfo")
    sClientVersion := ctx.QueryParam("ClientVersion")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if sErrInfoMd5 == "" || sClientVersion == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb
    sql := "SELECT time, zoneid, roleid FROM client_error "
    sql += "WHERE stackmd5 = '" + sErrInfoMd5 + "'"
    sql += "AND client_version = '" + sClientVersion + "'"

    var total int
    row2, err2 := db.Query("SELECT count(*) as total FROM ("+sql+") a")
    if err2 != nil {
        return err2
    }
    for row2.Next() {
        if err := row2.Scan(&total); err != nil {
            return err
        }
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

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

    return ctx.SendArray(slErrInfo, total)
}

func ErrDispose(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrInfo := ctx.FormValue("sErrInfo")
    sErrInfoMd5 := ctx.FormValue("sErrInfoMd5")
    sClientVersion := ctx.FormValue("sClientVersion")
    if sErrInfoMd5 == "" || sClientVersion == "" || sErrInfo == "" {
        return ctx.SendError(-1, "参数非法")
    }
    db := cfg.StatDb

    sql := "INSERT INTO client_dispose (stack, stackmd5, client_version, status, note) VALUES (?,?,?,?,?)"
    rows, err := db.Query(sql, sErrInfo, sErrInfoMd5, sClientVersion, 2, "")
    if err != nil {
        return err
    }
    defer rows.Close()

    return ctx.SendResponse("添加处理成功")
}

func FinishDispose(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrInfoMd5 := ctx.FormValue("sErrInfoMd5")
    sClientVersion := ctx.FormValue("sClientVersion")
    if sErrInfoMd5 == "" || sClientVersion == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb

    sql := "UPDATE client_dispose SET status=3 WHERE client_version=? AND stackmd5=? "
    rows, err := db.Query(sql, sClientVersion, sErrInfoMd5)
    if err != nil {
        return err
    }
    defer rows.Close()

    return ctx.SendResponse("更新处理状态成功")
}

func DisposeList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    

    db := cfg.StatDb

    
    sql := "SELECT client_version, stack, stackmd5, status, note FROM client_dispose "
    var total int
    row2, err2 := db.Query("SELECT count(*) as total FROM ("+sql+") a")
    if err2 != nil {
        return err2
    }
    for row2.Next() {
        if err := row2.Scan(&total); err != nil {
            return err
        }
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slDisposeInfo := make([]disposeInfo, 0)
    for rows.Next() {
        var r disposeInfo
        if err := rows.Scan(&r.ClientVersion, &r.ErrMessage, &r.ErrMessageMd5, &r.Status, &r.Note); err != nil {
            return err
        }

        slDisposeInfo = append(slDisposeInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(disposeInfoBy(slDisposeInfo))

    return ctx.SendArray(slDisposeInfo, total)
}

func AddDisposeNote(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrInfoMd5 := ctx.FormValue("sErrInfoMd5")
    sClientVersion := ctx.FormValue("sClientVersion")
    sNote := ctx.FormValue("sNote")
    if sErrInfoMd5 == "" || sClientVersion == "" || sNote == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb

    sql := "UPDATE  client_dispose SET note=? WHERE client_version=? AND stackmd5=? "
    rows, err := db.Query(sql, sNote, sClientVersion, sErrInfoMd5)
    if err != nil {
        return err
    }
    defer rows.Close()

    return ctx.SendResponse("添加备注成功")
}

func delClientError(sDelId string) int {   
    db := cfg.LogDb
	if db == nil {
		return -1
	}
	tx, err := db.Begin()
	if err != nil {
		return -1
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"log_global")
	if err != nil {
		return -1
	}

    sql := "SELECT client_version, stackmd5 FROM client_error "
    where := "WHERE _rid IN ("+sDelId+")"
    sql += where
    sql += "GROUP BY client_version, stackmd5"
    rows, err := tx.Query(sql)
    if err != nil {
        log.Infof("查询语句: %s, 错误原因: %s", sql, err)
        return -1
    }

    clientVersions := make([]string, 0)
    stackmd5s := make([]string, 0)
    for rows.Next() {
       var clientVersion string
       var stackmd5 string
       if err := rows.Scan(&clientVersion, &stackmd5); err != nil {
           return -1
       }
       clientVersions = append(clientVersions, clientVersion)
       stackmd5s = append(stackmd5s, stackmd5)
    }

    var sClientVersion string
    var sStackmd5 string
    for k, v := range clientVersions {
        sClientVersion = sClientVersion + "'" + v + "'"
        if k != (len(clientVersions) - 1) {
            sClientVersion += ", "
        }
    }

    for k, v := range stackmd5s {
        sStackmd5 = sStackmd5 + "'" + v + "'"
        if k != (len(stackmd5s) - 1) {
            sStackmd5 += ", "
        }
    }

	_, err = tx.Exec("Delete FROM client_error WHERE client_version IN ("+sClientVersion+")" + " AND stackmd5 IN (" + sStackmd5 + ")" )
	if err != nil {
		return -1
	}

	if err := tx.Commit(); err != nil {
		return -1
	}

    return 0
}

func DelClientErr(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := ctx.FormValue("idsStr")

   iRet := delClientError(ids)
   if iRet != 0 {
       return ctx.SendError(iRet, "删除错误失败")
   }

    return ctx.SendResponse("删除错误成功")
    
}

func Clientlist(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    db := cfg.LogDb

    sql := "SELECT _rid, timeymd, client_version, stack, stackmd5, count(*) as count  FROM client_error "
    where := "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "'"
    sql += where
    sql += "GROUP BY client_version"

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slSimpleErrInfo := make([]errSimpleInfo, 0)
    for rows.Next() {
        var r errSimpleInfo
        if err := rows.Scan(&r.ErrId, &r.ErrTime, &r.ClientVersion, &r.ErrMessage, &r.ErrMessageMd5, &r.ErrTimes); err != nil {
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