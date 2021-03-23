package gm

import (
    "time"
	"strconv"
    "strings"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type resControl struct {
    ResId	uint32	`json:"iResId"`
    Access  []string `json:"sAccess"`
};

var vAccess []*string
var nextUpdateTime time.Time

func ResControlList(c echo.Context) error {
    ctx := c.(*mid.Context)

    db := cfg.GameGlobalDb
    sql := "SELECT res_id, access from res_control"
    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    vResControl := make([]resControl, 0)
    for row.Next() {
        var r resControl
        var sAccess string
        if err := rows.Scan(&r.ResId, &sAccess) != nil {
            return err
        }

        r.Access = strings.Split(sAccess, ",")
        vResControl = append(vResControl, r)
    }

    return ctx.SendArray(vResControl, len(vResControl))
}

func AccessList(c echo.Context) error {
    ctx := c.(*mid.Context)

    accessList := getAllAccess()
    return ctx.SendResponse(accessList)
}

func refreshAccessList() {
    now := time.Now()
    if now.Before(nextUpdateTime) {
        return
    }

    db := cfg.GameGlobalDb
    sql := "SELECT action from user_action"
    rows, err := db.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    vtmp := make([]*string, 0)
    for rows.Next() {
        var r string
        if err := rows.Scan(&r) != nil {
            return err
        }
        vtmp = append(vtmp, &r)
    }
    vAccess = vtmp

    nextUpdateTime = now.Add(time.Minute * 5)
}

func getAllAccess() []*string {
    refreshAccessList()
    
    allAccess := make([]*string, len(vAccess))

    copy(allAccess, vAccess)

    return allAccess
}

func AccessUpdate(c echo.Context) error {   
}