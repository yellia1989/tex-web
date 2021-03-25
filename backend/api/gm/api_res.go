package gm

import (
    "time"
    "strings"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/cfg"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/tools/log"
    "strconv"
)

type resControl struct {
    ResId	uint32	`json:"iResId"`
    Action  []string `json:"sAction"`
};

type Action struct {
    Vaule string `json:"vaule"`
    Name string `json:"name"`
}

var vAction []Action
var nextUpdateTime time.Time

func ResControlList(c echo.Context) error {
    ctx := c.(*mid.Context)

    db := cfg.GameGlobalDb
    sql := "SELECT res_id, action from res_control"
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
        vResControl = append(vResControl, r)
    }

    return ctx.SendArray(vResControl, len(vResControl))
}

func ActionList(c echo.Context) error {
    ctx := c.(*mid.Context)

    actionList := getAllAction()

    return ctx.SendArray(actionList, len(actionList));
}

func refreshActionList() {
    now := time.Now()
    if now.Before(nextUpdateTime) {
        return
    }

    db := cfg.GameGlobalDb
    sql := "SELECT action from user_action"
    rows, err := db.Query(sql)
    if err != nil {
        return
    }
    defer rows.Close()

    vtmp := make([]Action, 0)
    for rows.Next() {
        var r Action
        if err := rows.Scan(&r.Name); err != nil {
            return
        }
        r.Vaule = r.Name
        vtmp = append(vtmp, r)
    }
    vAction = vtmp

    nextUpdateTime = now.Add(time.Minute * 5)
}

func getAllAction() []Action {
    refreshActionList()
    
    allAction := make([]Action, len(vAction))

    copy(allAction, vAction)

    return allAction
}

func ActionAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    resId,_ := strconv.Atoi(ctx.FormValue("iResId"))
    sAction := ctx.FormValue("sAction")

    log.Infof("resId: %d\n sAction: %s",resId, sAction)

    db := cfg.GameGlobalDb
    sql := "INSERT INTO res_control (res_id, action) VALUES(?,?)"

    rows, err := db.Query(sql, resId, sAction)
    if err != nil {
        return err
    }
    defer rows.Close()

    return ctx.SendResponse("添加资源监控项成功")
}
