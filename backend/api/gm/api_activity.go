package gm

import (
    "fmt"
	"strconv"
    "encoding/json"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-go/tools/util"
	"github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/api/gm/rpc"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _activityData struct {
	ActivityID    uint32 `json:"activity_id"`
	ActivityType  uint32 `json:"activity_type"`
	ApplyZone     string `json:"apply_zone"`
	ApplyMap      string `json:"apply_map"`
	ConfigureData string `json:"configure_data"`
	ConfigureDesc string `json:"configure_desc"`
    Locked  bool    `json:"locked"`
    Slg uint32  `json:"slg"`
}

func ActivityList(c echo.Context) error {
	ctx := c.(*mid.Context)
	stype := ctx.QueryParam("activityType")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
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

	sql := "SELECT activity_id,activity_type,apply_zone,apply_map,slg,configure_data,configure_desc,locked FROM t_activity"
	where := ""
    if stype != "" {
		where += " activity_type IN (" + stype + ")"
	}
    if where != "" {
        sql += " WHERE " + where
    }
    sql += " ORDER BY activity_id desc"
    var total int
    err = tx.QueryRow("SELECT count(*) as total FROM ("+sql+") a").Scan(&total)
    if err != nil {
        return err
    }

    limitstart := strconv.Itoa((page-1)*limit)
    limitrow := strconv.Itoa(limit)
    sql += " LIMIT "+limitstart+","+limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]_activityData, 0)
	for rows.Next() {
		var r _activityData
		if err := rows.Scan(&r.ActivityID, &r.ActivityType, &r.ApplyZone, &r.ApplyMap, &r.Slg, &r.ConfigureData, &r.ConfigureDesc, &r.Locked); err != nil {
			return err
		}
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func ActivityAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    activity_id := ctx.FormValue("iActivityId")
    apply_zone := ctx.FormValue("apply_zone")
    apply_map := ctx.FormValue("apply_map")
    slg := ctx.FormValue("slg")
    config_desc := ctx.FormValue("configure_desc")
    config_data := ctx.FormValue("configure_data")

    json_data := make(map[string]interface{})
    err := json.Unmarshal([]byte(config_data), &json_data)
    if err != nil {
        return err
    }

    activity_type := json_data["type"]

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
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

	_, err = tx.Exec("insert into t_activity(activity_id,activity_type,apply_zone,apply_map,slg,configure_data,configure_desc) value(?,?,?,?,?,?,?)", activity_id, activity_type, apply_zone, apply_map, slg, config_data, config_desc)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("添加活动成功")
}

func ActivityEdit(c echo.Context) error {
    ctx := c.(*mid.Context)
    activity_id := ctx.FormValue("iActivityId")
    apply_zone := ctx.FormValue("apply_zone")
    apply_map := ctx.FormValue("apply_map")
    config_desc := ctx.FormValue("configure_desc")
    config_data := ctx.FormValue("configure_data")

    json_data := make(map[string]interface{})
    err := json.Unmarshal([]byte(config_data), &json_data)
    if err != nil {
        return err
    }

    activity_type := json_data["type"]

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
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

	result, err := tx.Exec("update t_activity set activity_type=?,apply_zone=?,apply_map=?,configure_data=?,configure_desc=? WHERE activity_id=? and locked=0",activity_type, apply_zone, apply_map, config_data, config_desc, activity_id)
	if err != nil {
		return err
	}

    updateRows, err := result.RowsAffected()
    if err != nil {
        return err
    }

	if err := tx.Commit(); err != nil {
		return err
	}

    if updateRows == 0 {
        return ctx.SendResponse("活动已锁定不能编辑");
    }

    return ctx.SendResponse("更新活动成功")
}

func ActivityDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := ctx.FormValue("idsStr")

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
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

	result, err := tx.Exec("Delete FROM t_activity WHERE activity_id IN ("+ids+") and locked=0")
	if err != nil {
		return err
	}

    updateRows, err := result.RowsAffected()
    if err != nil {
        return err
    }

	if err := tx.Commit(); err != nil {
		return err
	}

    if updateRows == 0 {
        return ctx.SendResponse("活动已锁定不能删除");
    }

    return ctx.SendResponse("删除活动成功")
}

type _importAct struct {
    Id int `json:id`
    Type int `json:type`
    Data string `json:data`
    Desc string `json:desc`
}

func ActivityImport(c echo.Context) error {
    ctx := c.(*mid.Context)
    apply_zone := ctx.FormValue("apply_zone")
    slg := ctx.FormValue("slg")
    apply_map := ctx.FormValue("apply_map")
    filename := ctx.FormValue("filepath")

    content, err := util.LoadFromFile(filename)
    if err != nil {
        return err
    }

    var acts []_importAct
    if err := json.Unmarshal(content, &acts); err != nil {
        return err
    }
    if len(acts) == 0 {
        return ctx.SendError(-1, "导入活动为空")
    }

    sql := "insert into t_activity(activity_id,activity_type,apply_zone,apply_map,slg,configure_data,configure_desc) values"
    for k, v := range acts {
        if k != 0 {
            sql += ","
        }
        sql += fmt.Sprintf("(%d,%d,'%s','%s', %s, '%s','%s')", v.Id, v.Type, apply_zone, apply_map, slg, v.Data, v.Desc)
    }

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
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

	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("批量导入活动成功")
}

func ActivityOnlineZone(c echo.Context) error {
	ctx := c.(*mid.Context)
	activityId := ctx.QueryParam("activity_id")

    comm := cfg.Comm
    app := cfg.App
    u := ctx.GetUser()

    dirPrx := new(rpc.DirService)
    comm.StringToProxy(app+".DirServer.DirServiceObj", dirPrx)

    var zones []rpc.ZoneInfo
    ret, err := dirPrx.GetAllZone(&zones)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    var onlinezones []uint32

    var result string
    for _, z := range zones {
        cmd := "is_activity_online " + activityId
        zoneid := fmt.Sprintf("%d", z.IZoneId)
        gamePrx := new(rpc.GameService)
        comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+zoneid, gamePrx)
        ret, err = gamePrx.DoGmCmd(u.UserName, cmd, &result)
        if ret == 0 && err == nil && result == "on" {
            onlinezones = append(onlinezones, z.IZoneId)
        }
    }

    return ctx.SendResponse(onlinezones)
}

func ActivityLock(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := ctx.FormValue("idsStr")
    locked,_ := strconv.Atoi(ctx.FormValue("locked"))

    if len(ids) == 0 {
        return ctx.SendError(-1, "参数非法")
    }

	db := cfg.GameDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
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

    sql := "UPDATE t_activity SET locked=? WHERE activity_id IN("+ids+")"
	_, err = tx.Exec(sql, locked)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("操作成功")
}
