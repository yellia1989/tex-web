package game

import (
    "sort"
    "strings"
    "encoding/base64"
    "strconv"
    "fmt"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/common"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
    "github.com/yellia1989/tex-web/backend/api/sys"
    "bytes"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

type fightVerifyErrInfo struct {
    ErrTime         string `json:"err_time"`
    ReportId        uint64 `json:"report_id"`
    StageId         uint32 `json:"stage_id"`
    RoleId          uint32 `json:"role_id"`
    ZoneId          uint32 `json:"zone_id"`
    FightType       uint32 `json:"fight_type"`
    LogMd5          string `json:"log_md5"`
    MapId           uint32 `json:"map_id"`
}

type fightVerifyErrInfoBy []fightVerifyErrInfo

func (a fightVerifyErrInfoBy) Len() int      { return len(a) }
func (a fightVerifyErrInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a fightVerifyErrInfoBy) Less(i, j int) bool {
    TmpTimeI := common.ParseTimeInLocal("2006-01-02", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("2006-01-02", a[j].ErrTime)

    if !TmpTimeI.Equal(TmpTimeJ) {
        return TmpTimeI.After(TmpTimeJ)
    }

    return a[i].ErrTime > a[j].ErrTime
}

func FightErrInfo(c echo.Context) error {
    ctx := c.(*mid.Context)
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb

    sql := "SELECT time, report_id, stage_id, roleid, zoneid, fight_type, log_md5, mapid FROM fight_verify_error "
    sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "' "
    sql += "AND is_server=1"

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slFightVerifyInfo := make([]fightVerifyErrInfo, 0)
    for rows.Next() {
        var r fightVerifyErrInfo
        if err := rows.Scan(&r.ErrTime, &r.ReportId, &r.StageId, &r.RoleId, &r.ZoneId, &r.FightType, &r.LogMd5, &r.MapId); err != nil {
            return err
        }

        slFightVerifyInfo = append(slFightVerifyInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sql = "SELECT time, report_id, stage_id, roleid, zoneid, fight_type FROM chapter_verify_error "
    sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "' "
    sql += "AND is_server=1"

    rows1, err1 := db.Query(sql)
    if err1 != nil {
        return err1
    }
    defer rows1.Close()

    for rows1.Next() {
        var r fightVerifyErrInfo
        if err := rows1.Scan(&r.ErrTime, &r.ReportId, &r.StageId, &r.RoleId, &r.ZoneId, &r.FightType); err != nil {
            return err
        }

        slFightVerifyInfo = append(slFightVerifyInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }
    sort.Sort(fightVerifyErrInfoBy(slFightVerifyInfo))

    vSimpleErrInfo := common.GetPage(slFightVerifyInfo, page, limit)

    return ctx.SendArray(vSimpleErrInfo, len(slFightVerifyInfo))
}

func FightExport(c echo.Context) error {
    ctx := c.(*mid.Context)
    szoneid := ctx.FormValue("zoneids")
    cmd := strings.ReplaceAll(strings.TrimSpace(ctx.FormValue("cmd")), "\t", " ")
    reportid := ctx.FormValue("reportid")
    logmd5 := ctx.FormValue("logmd5")
    fightType, _ := strconv.Atoi(ctx.FormValue("fighttype"))
    smapid := ctx.FormValue("mapids")

    if szoneid == "" || cmd == "" {
        return ctx.SendError(-1, "参数非法")
    }

    if (fightType == 8) {
        szoneid = smapid;
    }

    buff := bytes.Buffer{}
    u := ctx.GetUser()

    comm := cfg.Comm
    app := cfg.App

    izoneid,_ := strconv.Atoi(szoneid)
    gamePrx := new(rpc.GameService)
    gfPrx := new(rpc.GFService)
    mapPrx := new(rpc.MapService)
    if izoneid != 0 {
        if izoneid != 8888 && izoneid != 9999 && izoneid > 1000 {
            comm.StringToProxy(app+".MapServer.MapServiceObj%"+app+".map."+ szoneid, mapPrx)
        } else {
            comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+ szoneid, gamePrx)
        }
    } else {
        comm.StringToProxy(app+".GFServer.GFServiceObj", gfPrx)
    }

    cmd = strings.Trim(strings.ReplaceAll(cmd, "   ", ""), " ")

    result := ""
    var ret int32
    var err error

    buff.WriteString("zone["+szoneid + "] > " + cmd + "\n")

    if izoneid != 0 {
        if izoneid == 8888 || izoneid == 9999 || izoneid <= 1000 {
            ret, err = gamePrx.DoGmCmd(u.UserName, cmd, &result)
        } else {
            ret, err = mapPrx.DoGmCmd(u.UserName, cmd, &result)
        }
    } else {
        ret, err = gfPrx.DoGmCmd(u.UserName, cmd, &result)
    }
    if ret != 0 || err != nil {
        serr := ""
        if err != nil {
            serr = err.Error()
        }
        result = fmt.Sprintf("ret:%s, err:%s\n", rpc.ErrorCode(ret), serr)
    }
    buff.WriteString(result+"\n")

    sys.LogAdd(u.UserName, "gm", "[" + szoneid + "]>" + cmd)

    splitLine := "\n\n==========================================================================\n\n"

    vString := make([]string, 0)
    vString = append(vString, " ========\n|| 服务器战报 ||\n ========\n\n")
    vString = append(vString, buff.String() + splitLine)

    db := cfg.LogDb
    var sql string

    if (fightType != 11) {
      sql = "SELECT log, is_server FROM fight_verify_error WHERE log_md5 = '" + logmd5 + "' "
    } else {
      sql = "SELECT log, is_server FROM chapter_verify_error WHERE report_id = '" + reportid + "' "
    }

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    var clientLog string
    var serverLog string
    for rows.Next() {
        var log string
        var bServer uint32
        if err := rows.Scan(&log, &bServer); err != nil {
            return err
        }

        if (fightType != 11) {
            decodeBytes, _ := base64.StdEncoding.DecodeString(log)
            log = string(decodeBytes)
        }

        if bServer == 1 {
            serverLog = log;
        } else {
            clientLog = log;
        }
    }
    vString = append(vString, " ========\n|| 客户端日志 ||\n ========\n\n")
    vString = append(vString, clientLog + splitLine)
    vString = append(vString, " ========\n|| 服务器日志 ||\n ========\n\n")
    vString = append(vString, serverLog + splitLine)

    return ctx.SendResponse(vString)
}
