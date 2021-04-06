package gm

import (
    "errors"
    "sort"
    "strconv"
    "strings"
    "time"
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/common"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

type resControl struct {
    ResId      uint32   `json:"iResId"`
    Action     []string `json:"sAction"`
    ActionName []string `json:"sActionName"`
}

type Action struct {
    Vaule string `json:"vaule"`
    Name  string `json:"name"`
}

type resErrSimpleInfo struct {
    ErrTime       string `json:"err_time"`
    ErrResId      string `json:"err_res_id"`
    ErrTimes      uint32 `json:"err_times"`
    ErrAction     string `json:"err_action"`
    ErrActionName string `json:"err_action_name"`
}

type resErrSimpleInfoBy []resErrSimpleInfo

func (a resErrSimpleInfoBy) Len() int      { return len(a) }
func (a resErrSimpleInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a resErrSimpleInfoBy) Less(i, j int) bool {
    TmpTimeI := common.ParseTimeInLocal("2006-01-02", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("2006-01-02", a[j].ErrTime)

    if !TmpTimeI.Equal(TmpTimeJ) {
        return TmpTimeI.After(TmpTimeJ)
    }

    return a[i].ErrTimes > a[j].ErrTimes 
}

type resErrInfo struct {
    ErrTime       string `json:"err_time"`
    ErrResId      string `json:"err_res_id"`
    ErrAction     string `json:"err_action"`
    ErrActionName string `json:"err_action_name"`
    ZoneId        uint32 `json:"zone_id"`
    RoleId        uint32 `json:"role_id"`
    ErrParam      uint32 `json:"err_param"`
}

type resErrInfoBy []resErrInfo

func (a resErrInfoBy) Len() int      { return len(a) }
func (a resErrInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a resErrInfoBy) Less(i, j int) bool {
    TmpTimeI := common.ParseTimeInLocal("15:04:05", a[i].ErrTime)
    TmpTimeJ := common.ParseTimeInLocal("15:04:05", a[j].ErrTime)

    if !TmpTimeI.Equal(TmpTimeJ) {
        return TmpTimeI.After(TmpTimeJ)
    }

    return a[i].ZoneId < a[j].ZoneId
}

var vAction []Action
var nextUpdateTime time.Time
var mAction map[string]string

func ResControlList(c echo.Context) error {
    ctx := c.(*mid.Context)

    refreshActionList(true)

    db := cfg.GameGlobalDb
    if db == nil {
        ctx.SendError(-1, "数据库空")
    }

	tx, err := db.Begin()
	if err != nil {
        return err
	}
    defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
        return err
	}

    sql := "SELECT res_id, action FROM t_res_control"
    rows, err := tx.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    vResControl := make([]resControl, 0)
    for rows.Next() {
        var r resControl
        var sAction string
        if err := rows.Scan(&r.ResId, &sAction); err != nil {
            return err
        }

        r.Action = strings.Split(sAction, ",")
        r.ActionName, err = getActionName(r.Action)
        if err != nil {
            return err
        }
        vResControl = append(vResControl, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendArray(vResControl, len(vResControl))
}

func ActionList(c echo.Context) error {
    ctx := c.(*mid.Context)

    actionList := getAllAction()

    return ctx.SendArray(actionList, len(actionList))
}

func refreshActionList(bForce bool) {
    now := time.Now()
    if now.Before(nextUpdateTime) && !bForce{
        return
    }

    db := cfg.StatDb
    sql := "SELECT action, action_name from user_action"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }
    defer rows.Close()

    vtmp := make([]Action, 0)
    mtmp := make(map[string]string)
    for rows.Next() {
        var r Action
        if err := rows.Scan(&r.Vaule, &r.Name); err != nil {
            return
        }
        vtmp = append(vtmp, r)
        mtmp[r.Vaule] = r.Name
    }
    vAction = vtmp
    mAction = mtmp

    nextUpdateTime = now.Add(time.Minute * 5)
}

func getActionName(action []string) ([]string, error) {
    vAction := make([]string, 0)
    for _, v := range action {
        val, ok := mAction[v]
        if ok {
            vAction = append(vAction, val)
        } else {
            return nil, errors.New("key : " + v + " is nil Value!")
        }
    }
    return vAction, nil
}

func getAllAction() []Action {
    refreshActionList(false)

    allAction := make([]Action, len(vAction))

    copy(allAction, vAction)

    return allAction
}

func ActionAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))
    sAction := ctx.FormValue("sAction")

    if iResId == 0 || sAction == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.GameGlobalDb
    if db == nil {
        ctx.SendError(-1, "数据库空")
    }

	tx, err := db.Begin()
	if err != nil {
        return err
	}
    defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
        return err
	}

    sql := "INSERT INTO t_res_control (res_id, action) VALUES(?,?)"

    rows, err := tx.Query(sql, iResId, sAction)
    if err != nil {
        return err
    }
    defer rows.Close()

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("添加资源监控项成功")
}

func ActionEdit(c echo.Context) error {
    ctx := c.(*mid.Context)
    iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))
    sAction := ctx.FormValue("sAction")

    if iResId == 0 || sAction == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.GameGlobalDb
    if db == nil {
        ctx.SendError(-1, "数据库空")
    }

	tx, err := db.Begin()
	if err != nil {
        return err
	}
    defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
        return err
	}

    sql := "UPDATE t_res_control SET action=? WHERE res_id=?"

    rows, err := tx.Query(sql, sAction, iResId)
    if err != nil {
        return err
    }
    defer rows.Close()

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("编辑资源监控项成功")
}

func ActionDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))

    if iResId == 0 {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.GameGlobalDb
    if db == nil {
        ctx.SendError(-1, "数据库空")
    }

	tx, err := db.Begin()
	if err != nil {
        return err
	}
    defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
        return err
	}

    sql := "DELETE FROM t_res_control WHERE res_id=?"

    rows, err := tx.Query(sql, iResId)
    if err != nil {
        return err
    }
    defer rows.Close()

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("删除资源监控项成功")
}

func ResErrInfo(c echo.Context) error {
    ctx := c.(*mid.Context)
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    refreshActionList(true)

    db := cfg.LogDb

    sql := "SELECT logymd, res_id, action, count(*) as count FROM res_action_prom_error "
    sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "' "
    sql += "GROUP BY logymd, res_id, action"

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slSimpleResErrInfo := make([]resErrSimpleInfo, 0)
    for rows.Next() {
        var r resErrSimpleInfo
        if err := rows.Scan(&r.ErrTime, &r.ErrResId, &r.ErrAction, &r.ErrTimes); err != nil {
            return err
        }

        var ok bool
        r.ErrActionName, ok = mAction[r.ErrAction]
        if !ok {
            r.ErrActionName = r.ErrAction
        }

        slSimpleResErrInfo = append(slSimpleResErrInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(resErrSimpleInfoBy(slSimpleResErrInfo))

    vSimpleResErrInfo := common.GetPage(slSimpleResErrInfo, page, limit)

    return ctx.SendArray(vSimpleResErrInfo, len(slSimpleResErrInfo))
}

func ResErrDetail(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrTime := ctx.QueryParam("ErrTime")
    sErrResId := ctx.QueryParam("ErrResId")
    sAction := ctx.QueryParam("ErrAction")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if sErrResId == "" || sAction == "" || sErrTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb
    sql := "SELECT loghms, res_id, action, zoneid, roleid FROM res_action_prom_error "
    sql += "WHERE STR_TO_DATE(logymd, '%Y-%m-%d') = STR_TO_DATE('" + sErrTime + "', '%Y-%m-%d')  AND res_id = '" + sErrResId + "'"
    sql += "AND action = '" + sAction + "'"

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slResErrInfo := make([]resErrInfo, 0)

    for rows.Next() {
        var r resErrInfo
        if err := rows.Scan(&r.ErrTime, &r.ErrResId, &r.ErrAction, &r.ZoneId, &r.RoleId); err != nil {
            return err
        }

        var ok bool
        r.ErrActionName, ok = mAction[r.ErrAction]
        if !ok {
            r.ErrActionName = r.ErrAction
        }

        slResErrInfo = append(slResErrInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(resErrInfoBy(slResErrInfo))
    vResErrInfo := common.GetPage(slResErrInfo, page, limit)

    return ctx.SendArray(vResErrInfo, len(slResErrInfo))
}

func ResAppendResControl(c echo.Context) error {
    ctx := c.(*mid.Context)
    sResId := ctx.FormValue("iResId")
    sAction := ctx.FormValue("sAction")

    if sResId == "" || sAction == "" {
        return ctx.SendError(-1, "参数非法")
    }

    sActionName, ok := mAction[sAction]
    if !ok {
        return ctx.SendError(-2, "未找到监控项: " + sAction + " 请点击 '添加可监控项' 按钮 添加到可监控项中")
    }

    db := cfg.GameGlobalDb
    if db == nil {
        ctx.SendError(-1, "数据库空")
    }

    db.SetConnMaxLifetime(time.Second * 30)
	tx, err := db.Begin()
	if err != nil {
        return err
	}
    defer tx.Rollback()

	_, err = tx.Exec("USE "+cfg.GameDbPrefix+"db_zone_global")
	if err != nil {
        return err
	}

    sql1 := "SELECT res_id FROM t_res_control WHERE res_id=?"
    sql2 := "INSERT INTO t_res_control set res_id = ?, action=?"
    sql3 := "UPDATE t_res_control set action=CONCAT(action,',',?) WHERE res_id=?"
    sql4 := "SELECT action FROM t_res_control WHERE res_id=?"

    rows1, err1 := tx.Query(sql1, sResId)
    if err1 != nil {
        return err1
    }

    // 没有则插入
    if !rows1.Next() && rows1.Err() == nil {
        rows2, err2 := tx.Query(sql2, sResId, sAction);
        if err2 != nil {
            return err2
        }
        defer rows2.Close()

	    if err = tx.Commit(); err != nil {
		    return err
	    }

        sSuccess := "在资源ID: " + sResId + " 中添加监控: " + sActionName + " 成功"
        return ctx.SendResponse(sSuccess)
    }
    rows1.Close()

    rows4, err4 := tx.Query(sql4, sResId)
    if err4 != nil {
        return err4
    }

    for rows4.Next() {
        var sAllAction string
        if err := rows4.Scan(&sAllAction); err != nil {
            return err
        }
        vAllAction := strings.Split(sAllAction, ",")
        for _,v := range vAllAction {
            if v == sAction {
	            defer tx.Commit()
                return ctx.SendError(-3, "本监控项已被监控")
            }
        }
    } 

    if err = rows4.Err(); err != nil {
        return err
    }
    rows4.Close()

    // 有则追加
    rows3, err3 := tx.Query(sql3, sAction, sResId)
    if err3 != nil {
        return err3
    }
    defer rows3.Close()

	if err = tx.Commit(); err != nil {
		return err
	}

    sSuccess := "在资源ID: " + sResId + " 中添加监控: " + sActionName + " 成功"
    return ctx.SendResponse(sSuccess)
}

func ResAppendAction(c echo.Context) error {
    ctx := c.(*mid.Context)
    sAction := ctx.FormValue("sAction")
    sActionName := ctx.FormValue("sActionName")

    _, ok := mAction[sAction]
    if ok {
        return ctx.SendError(-2, "监控项已存在")
    }

    if sAction == "" || sActionName == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb
    sql := "INSERT INTO user_action (action, action_name) VALUES(?,?)"

    rows, err := db.Query(sql, sAction, sActionName)
    if err != nil {
        return err
    }
    defer rows.Close()

    refreshActionList(true)

    return ctx.SendResponse("添加可监控项: " + sActionName + " 成功！")
}

func ResNumErrInfo(c echo.Context) error {
    ctx := c.(*mid.Context)
    startTime := ctx.QueryParam("startTime")
    endTime := ctx.QueryParam("endTime")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if startTime == "" || endTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb

    sql := "SELECT logymd, res_id, count(*) as count  FROM res_add_prom_error "
    sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "' "
    sql += "GROUP BY logymd, res_id"

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slSimpleResErrInfo := make([]resErrSimpleInfo, 0)
    for rows.Next() {
        var r resErrSimpleInfo
        if err := rows.Scan(&r.ErrTime, &r.ErrResId, &r.ErrTimes); err != nil {
            return err
        }

        slSimpleResErrInfo = append(slSimpleResErrInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(resErrSimpleInfoBy(slSimpleResErrInfo))

    vSimpleResErrInfo := common.GetPage(slSimpleResErrInfo, page, limit)

    return ctx.SendArray(vSimpleResErrInfo, len(slSimpleResErrInfo))
}

func ResNumErrDetail(c echo.Context) error {
    ctx := c.(*mid.Context)
    sErrTime := ctx.QueryParam("ErrTime")
    sErrResId := ctx.QueryParam("ErrResId")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    if sErrResId == "" || sErrTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb
    sql := "SELECT loghms, zoneid, param1, roleid FROM res_add_prom_error "
    sql += "WHERE STR_TO_DATE(logymd, '%Y-%m-%d') = STR_TO_DATE('" + sErrTime + "', '%Y-%m-%d')  AND res_id = '" + sErrResId + "'"

    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    slResErrInfo := make([]resErrInfo, 0)

    for rows.Next() {
        var r resErrInfo
        if err := rows.Scan(&r.ErrTime, &r.ZoneId, &r.ErrParam, &r.RoleId); err != nil {
            return err
        }

        slResErrInfo = append(slResErrInfo, r)
    }

    if err := rows.Err(); err != nil {
        return err
    }

    sort.Sort(resErrInfoBy(slResErrInfo))
    vResErrInfo := common.GetPage(slResErrInfo, page, limit)

    return ctx.SendArray(vResErrInfo, len(slResErrInfo))
}
