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
    if a[i].ErrTime > a[j].ErrTime {
        return true
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
}

type resErrInfoBy []resErrInfo

func (a resErrInfoBy) Len() int      { return len(a) }
func (a resErrInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a resErrInfoBy) Less(i, j int) bool {
    if a[i].ErrTime > a[j].ErrTime {
        return true
    }

    return a[i].ZoneId < a[j].ZoneId
}

var vAction []Action
var nextUpdateTime time.Time
var mAction map[string]string

func ResControlList(c echo.Context) error {
    ctx := c.(*mid.Context)

    refreshActionList()

    db := cfg.GameGlobalDb
    sql := "SELECT res_id, action FROM t_res_control"
    rows, err := db.Query(sql)
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

    return ctx.SendArray(vResControl, len(vResControl))
}

func ActionList(c echo.Context) error {
    ctx := c.(*mid.Context)

    actionList := getAllAction()

    return ctx.SendArray(actionList, len(actionList))
}

func refreshActionList() {
    now := time.Now()
    if now.Before(nextUpdateTime) {
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
    refreshActionList()

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
    sql := "INSERT INTO t_res_control (res_id, action) VALUES(?,?)"

    rows, err := db.Query(sql, iResId, sAction)
    if err != nil {
        return err
    }
    defer rows.Close()

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
    sql := "UPDATE t_res_control SET action=? WHERE res_id=?"

    rows, err := db.Query(sql, sAction, iResId)
    if err != nil {
        return err
    }
    defer rows.Close()

    return ctx.SendResponse("编辑资源监控项成功")
}

func ActionDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    iResId, _ := strconv.Atoi(ctx.FormValue("iResId"))

    if iResId == 0 {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.GameGlobalDb
    sql := "DELETE FROM t_res_control WHERE res_id=?"

    rows, err := db.Query(sql, iResId)
    if err != nil {
        return err
    }
    defer rows.Close()

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

    refreshActionList()

    db := cfg.LogDb

    sql := "SELECT timeymd, res_id, action, count(*) as count  FROM res_prom_error "
    sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "'"
    sql += "GROUP BY timeymd, res_id, action"

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

    if sErrResId == "" || sAction == "" || sErrTime == "" {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.LogDb
    sql := "SELECT timehms, res_id, action, zoneid, roleid FROM res_prom_error "
    sql += "WHERE STR_TO_DATE(timeymd, '%Y-%m-%d') = STR_TO_DATE('" + sErrTime + "', '%Y-%m-%d')  AND res_id = '" + sErrResId + "'"
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

    return ctx.SendArray(slResErrInfo, len(slResErrInfo))
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
    sql := "UPDATE t_res_control set action=CONCAT(action,',',?) WHERE res_id=?"

    rows, err := db.Query(sql, sAction, sResId)
    if err != nil {
        return err
    }
    defer rows.Close()

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

    return ctx.SendResponse("添加可监控项: " + sActionName + " 成功！")
}
