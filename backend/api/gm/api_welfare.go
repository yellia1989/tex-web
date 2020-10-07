package gm

import (
    "strings"
    "strconv"
    dsql "database/sql"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _WelfareTask struct {
    ID  uint32  `json:"iId"`
    Name string `json:"sName"`
    Cmds string `json:"sCmds"`
    Roles string `json:"sRoles"`
    BeginTime string `json:"sBeginTime"`
    EndTime string `json:"sEndTime"`
    CmdBeginTime string `json:"sCmdBeginTime"`
    CmdEndTime string `json:"sCmdEndTime"`
    Status  uint32  `json:"iStatus"`
    Step uint32  `json:"iStep"`
}

func WelfareTaskList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.StatDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    sql := "SELECT id,name,roles,cmds,begin_time,end_time,cmd_time,status,step FROM welfare_task WHERE status != 2 order by id desc"
    var total int
    err := db.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

	c.Logger().Debug(sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

    tasks := make([]_WelfareTask,0)
    var cmdTime string
    for rows.Next() {
        var task _WelfareTask
        if err := rows.Scan(&task.ID, &task.Name, &task.Roles, &task.Cmds, &task.BeginTime, &task.EndTime, &cmdTime, &task.Status, &task.Step); err != nil {
            return err
        }
        vCmdTime := strings.SplitN(cmdTime, "-", 2)
        task.CmdBeginTime = vCmdTime[0]
        task.CmdEndTime = vCmdTime[1]
        tasks = append(tasks, task)
    }

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(tasks, total)
}

func WelfareTaskPause(c echo.Context) error {
	ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "福利id不存在")
    }

    db := cfg.StatDb
    if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
    }

    sql := "UPDATE welfare_task SET status=0 WHERE id in (?)"
    _, err := db.Exec(sql, strings.Join(ids, ","))
    if err != nil {
        return err
    }

    return ctx.SendResponse("暂停福利成功")
}

func WelfareTaskResume(c echo.Context) error {
	ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "福利id不存在")
    }

    db := cfg.StatDb
    if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
    }

    sql := "UPDATE welfare_task SET status=1 WHERE id in (?)"
    _, err := db.Exec(sql, strings.Join(ids, ","))
    if err != nil {
        return err
    }

    return ctx.SendResponse("恢复福利成功")
}

func WelfareTaskUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)
    sId := ctx.FormValue("id")
    sName := ctx.FormValue("sName")
    sBeginTime := ctx.FormValue("sBeginTime")
    sEndTime := ctx.FormValue("sEndTime")
    sCmdBeginTime := ctx.FormValue("sCmdBeginTime")
    sCmdEndTime := ctx.FormValue("sCmdEndTime")
    sRoles := ctx.FormValue("sRoles")
    sCmds := ctx.FormValue("sCmds")
    step,_ := strconv.ParseUint(ctx.FormValue("iStep"), 10, 32)

    if sId == "" || sName == "" || sBeginTime == "" || sEndTime == "" || sCmdBeginTime == "" || sCmdEndTime == "" || sRoles == "" || sCmds == "" || step == 0 {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb
    if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
    }

    sql := "UPDATE welfare_task SET name=?, roles=?, cmds=? ,cmd_time=? ,begin_time=?, end_time=?, step=? WHERE id=?"
    _, err := db.Exec(sql, sName, sRoles, sCmds, sCmdBeginTime+"-"+sCmdEndTime, sBeginTime, sEndTime, step, sId)
    if err != nil {
        return err
    }

    return ctx.SendResponse("更新福利成功")
}

func WelfareTaskAdd(c echo.Context) error {
	ctx := c.(*mid.Context)
    sName := ctx.FormValue("sName")
    sBeginTime := ctx.FormValue("sBeginTime")
    sEndTime := ctx.FormValue("sEndTime")
    sCmdBeginTime := ctx.FormValue("sCmdBeginTime")
    sCmdEndTime := ctx.FormValue("sCmdEndTime")
    sRoles := ctx.FormValue("sRoles")
    sCmds := ctx.FormValue("sCmds")
    step,_ := strconv.ParseUint(ctx.FormValue("iStep"), 10, 32)

    if sName == "" || sBeginTime == "" || sEndTime == "" || sCmdBeginTime == "" || sCmdEndTime == "" || sRoles == "" || sCmds == "" || step == 0 {
        return ctx.SendError(-1, "参数非法")
    }

    db := cfg.StatDb
    if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
    }

    sql := "INSERT INTO welfare_task(name,roles,cmds,cmd_time,status,begin_time,end_time,step) VALUES(?,?,?,?,?,?,?,?)"
    _, err := db.Exec(sql, sName, sRoles, sCmds, sCmdBeginTime+"-"+sCmdEndTime, 1, sBeginTime, sEndTime, step)
    if err != nil {
        return err
    }

    return ctx.SendResponse("添加福利成功")
}

func WelfareTaskDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "福利id不存在")
    }

    db := cfg.StatDb
    if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
    }

    sql := "UPDATE welfare_task SET status=2 WHERE id in(?)"
    _, err := db.Exec(sql, strings.Join(ids, ","))
    if err != nil {
        return err
    }

    return ctx.SendResponse("删除福利成功")
}

type _WelfareRole struct {
    TaskName    string `json:"taskname"`
    Zoneid  int `json:"zoneid"`
    Roleid  int `json:"roleid"`
    Cmd string `json:"cmd"`
    Time    string `json:"time"`
    ExecTime    string `json:"exec_time"`
    ExecResult  string `json:"exec_result"`
}

func WelfareRoleList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    taskid := ctx.QueryParam("taskid")
    roleid := ctx.QueryParam("roleid")
    zoneid := ctx.QueryParam("zoneid")
    begin_time := ctx.QueryParam("startTime")
    end_time := ctx.QueryParam("endTime")

	db := cfg.StatDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

    sql := "SELECT welfare_task.name as taskname,a.roleid,a.zoneid,a.cmd,a.time,a.exec_time,a.exec_result FROM (SELECT roleid,zoneid,cmd,time,exec_time,exec_result,taskid_pk FROM welfare_roles WHERE time BETWEEN '"+begin_time+"' AND '"+end_time+"' and zoneid="+zoneid+" and taskid_pk="+taskid
    if roleid != "" {
        sql += " and roleid="+roleid
    }
    sql += ") a LEFT JOIN welfare_task on welfare_task.id = a.taskid_pk"
    var total int
    err := db.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

    roles := make([]_WelfareRole,0)
    for rows.Next() {
        var r _WelfareRole
        var exec_time dsql.NullString
        var exec_result dsql.NullString
        if err := rows.Scan(&r.TaskName, &r.Roleid, &r.Zoneid, &r.Cmd, &r.Time, &exec_time, &exec_result); err != nil {
            return err
        }
        r.ExecTime = exec_time.String
        r.ExecResult = exec_result.String
        roles = append(roles, r)
    }

	return ctx.SendArray(roles, total)
}
