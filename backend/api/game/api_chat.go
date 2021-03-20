package game

import (
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    ssql "database/sql"
    "strconv"
    "strings"
    "sync"
)

var rwMutex sync.RWMutex
var maskWord = make([]string,0)
var maskReload bool

type chatLog struct {
    Id     uint32 `json:"id"`
    Time   string `json:"time"`
    Type   uint32  `json:"type"`
    FromMapId uint32 `json:"fromMapId"`
    FromZoneId uint32 `json:"fromZoneId"`
    FromRoleId uint32 `json:"fromRoleId"`
    FromRoleName string `json:"fromRoleName"`
    ToZoneId uint32 `json:"toZoneId"`
    ToRoleId uint32 `json:"toRoleId"`
    ToRoleName string `json:"toRoleName"`
    AllianceId uint64 `json:"allianceId"`
    AllianceName string `json:"allianceName"`
    Content string `json:"content"`
    DirtyWord string `json:"dirtyWord"`
}

func GetMaskWord() []string  {
    rwMutex.RLock()
    if !maskReload {
        rwMutex.RUnlock()
        reloadMaskWord()
    } else {
        defer rwMutex.RUnlock()
    }

    return maskWord
}

func reloadMaskWord() {
    rwMutex.Lock()
    defer rwMutex.Unlock()

    if maskReload {
        return
    }

    db := cfg.StatDb
    if db == nil {
        log.Errorf("logDb is Null")
        return
    }

    var wordStr ssql.NullString
    err := db.QueryRow("select words from chat_dirty_word limit 1").Scan(&wordStr)
    if err != nil && err != ssql.ErrNoRows {
        log.Errorf("reloadMaskWord err: %s", err.Error())
        return
    }

    if wordStr.String != "" {
        maskWord = strings.Split(wordStr.String,";")
    }
    maskReload = true
}

func ChatGetNewest(c echo.Context) error {
    ctx := c.(*mid.Context)
    fromid, _ := strconv.Atoi(ctx.QueryParam("maxid"))

    db := cfg.LogDb
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    var limit = 140
    if fromid == 0 {
        var maxid ssql.NullInt32
        if err := db.QueryRow("select max(_rid) from chat").Scan(&maxid); err != nil {
            return err
        }
        if int(maxid.Int32) > limit {
            fromid = int(maxid.Int32) - limit
        }
    }

    sql := "select _rid,time,zoneid,mapid,type,fromroleid,fromrolename,tozoneid,toroleid,torolename,allianceid,alliancename,content from chat where _rid > ? limit ?"
    rows, err := db.Query(sql, fromid, limit)
    if err != nil {
        return err
    }

    defer rows.Close()

    logs := make([]chatLog, 0)
    for rows.Next() {
        var r chatLog
        if err := rows.Scan(&r.Id, &r.Time, &r.FromZoneId, &r.FromMapId, &r.Type, &r.FromRoleId, &r.FromRoleName, &r.ToZoneId, &r.ToRoleId, &r.ToRoleName, &r.AllianceId, &r.AllianceName, &r.Content); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, limit)
}

func ChatGetHistory(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    roleid := ctx.QueryParam("roleid")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    if zoneid == "" || startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    sqlcount := "select count(_rid) as total FROM chat"
    sqlcount += " WHERE zoneid=" + zoneid + " AND time between '" + startTime + "' AND '" + endTime + "'"
    if roleid != "" {
        sqlcount += " AND fromroleid=" + roleid
    }

    var total int
    err := db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page - 1) * limit)
    limitrow := strconv.Itoa(limit)
    sql := "select _rid,time,zoneid,mapid,type,fromroleid,fromrolename,tozoneid,toroleid,torolename,allianceid,alliancename,content from chat"
    sql += " WHERE zoneid=" + zoneid + " AND time between '" + startTime + "' AND '" + endTime + "'"
    if roleid != "" {
        sql += " AND fromroleid=" + roleid
    }
    sql += " ORDER BY time desc, _rid desc"
    sql += " LIMIT " + limitstart + "," + limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]chatLog, 0)
    for rows.Next() {
        var r chatLog
        if err := rows.Scan(&r.Id, &r.Time, &r.FromZoneId, &r.FromMapId, &r.Type, &r.FromRoleId, &r.FromRoleName, &r.ToZoneId, &r.ToRoleId, &r.ToRoleName, &r.AllianceId, &r.AllianceName, &r.Content); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, total)
}

func ChatGetMaskNewest(c echo.Context) error {
    ctx := c.(*mid.Context)
    fromid, _ := strconv.Atoi(ctx.QueryParam("maxid"))

    db := cfg.StatDb
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    var limit = 140
    if fromid == 0 {
        var maxid ssql.NullInt32
        if err := db.QueryRow("select max(_rid) from chat_dirty_history").Scan(&maxid); err != nil {
            return err
        }
        if int(maxid.Int32) > limit {
            fromid = int(maxid.Int32) - limit
        }
    }
    sql := "select _rid,time,zoneid,mapid,type,fromroleid,fromrolename,tozoneid,toroleid,torolename,allianceid,alliancename,content,dirtyword from chat_dirty_history where _rid > ? limit ?"
    rows, err := db.Query(sql, fromid, limit)
    if err != nil {
        return err
    }

    defer rows.Close()

    logs := make([]chatLog, 0)
    for rows.Next() {
        var r chatLog
        if err := rows.Scan(&r.Id, &r.Time, &r.FromZoneId, &r.FromMapId, &r.Type, &r.FromRoleId, &r.FromRoleName, &r.ToZoneId, &r.ToRoleId, &r.ToRoleName, &r.AllianceId, &r.AllianceName, &r.Content, &r.DirtyWord); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, limit)
}

func ChatGetMaskLogs(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    roleid := ctx.QueryParam("roleid")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")

    if zoneid == "" || startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    sqlcount := "SELECT count(_rid) as total FROM chat_dirty_history "
    sqlcount += " WHERE zoneid=" + zoneid + " AND time between '" + startTime + "' AND '" + endTime + "'"
    if roleid != "" {
        sqlcount += " AND fromroleid=" + roleid
    }
    var total int
    err := db.QueryRow(sqlcount).Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page - 1) * limit)
    limitrow := strconv.Itoa(limit)
    sql := "select _rid,time,zoneid,mapid,type,fromroleid,fromrolename,tozoneid,toroleid,torolename,allianceid,alliancename,content,dirtyword from chat_dirty_history"
    sql += " WHERE zoneid=" + zoneid + " AND time between '" + startTime + "' AND '" + endTime + "'"
    if roleid != "" {
        sql += " AND fromroleid=" + roleid
    }
    sql += " ORDER BY time desc, _rid desc"
    sql += " LIMIT " + limitstart + "," + limitrow

    log.Infof("sql: %s", sql)

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    logs := make([]chatLog, 0)
    for rows.Next() {
        var r chatLog
        if err := rows.Scan(&r.Id, &r.Time, &r.FromZoneId, &r.FromMapId, &r.Type, &r.FromRoleId, &r.FromRoleName, &r.ToZoneId, &r.ToRoleId, &r.ToRoleName, &r.AllianceId, &r.AllianceName, &r.Content, &r.DirtyWord); err != nil {
            return err
        }
        logs = append(logs, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    return ctx.SendArray(logs, limit)
}

func ChatGetMaskWord(c echo.Context) error {
    ctx := c.(*mid.Context)
    maskWord := GetMaskWord()

    return ctx.SendResponse(strings.Join(maskWord,"\n"))
}

func ChatSetMaskWord(c echo.Context) error {
    ctx := c.(*mid.Context)
    wordStr := ctx.FormValue("input")
    stringArr:= strings.Split(wordStr,"\n")

    db := cfg.StatDb
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败")
    }

    tempStr := strings.Join(stringArr,";")
    _, err := db.Exec("insert into chat_dirty_word (id,words) values (1,?) on duplicate key update words=?", tempStr, tempStr)
    if err != nil {
        return err
    }

    rwMutex.Lock()
    maskReload = false
    rwMutex.Unlock()

    return ctx.SendResponse("设置屏蔽字成功")
}
