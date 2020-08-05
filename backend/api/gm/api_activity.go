package gm

import (
    "fmt"
	"strconv"
    "encoding/json"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-go/tools/util"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type _activityData struct {
	ActivityID    uint32 `json:"activity_id"`
	ActivityType  uint32 `json:"activity_type"`
	ApplyZone     string `json:"apply_zone"`
	ApplyUser     string `json:"apply_user"`
	ConfigureData string `json:"configure_data"`
	ConfigureDesc string `json:"configure_desc"`
}

func ActivityList(c echo.Context) error {
	ctx := c.(*mid.Context)
	stype := ctx.QueryParam("activityType")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_zone_global")
	if err != nil {
		return err
	}

	sql := "SELECT activity_id,activity_type,apply_zone,apply_user,configure_data,configure_desc FROM t_activity"
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
		if err := rows.Scan(&r.ActivityID, &r.ActivityType, &r.ApplyZone, &r.ApplyUser, &r.ConfigureData, &r.ConfigureDesc); err != nil {
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
    apply_user := ctx.FormValue("apply_user")
    config_desc := ctx.FormValue("configure_desc")
    config_data := ctx.FormValue("configure_data")

    json_data := make(map[string]interface{})
    err := json.Unmarshal([]byte(config_data), &json_data)
    if err != nil {
        return err
    }

    activity_type := json_data["type"]

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_zone_global")
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into t_activity(activity_id,activity_type,apply_zone,apply_user,configure_data,configure_desc) value(?,?,?,?,?,?)", activity_id, activity_type, apply_zone, apply_user, config_data, config_desc)
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
    apply_user := ctx.FormValue("apply_user")
    config_desc := ctx.FormValue("configure_desc")
    config_data := ctx.FormValue("configure_data")

    json_data := make(map[string]interface{})
    err := json.Unmarshal([]byte(config_data), &json_data)
    if err != nil {
        return err
    }

    activity_type := json_data["type"]

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_zone_global")
	if err != nil {
		return err
	}

	_, err = tx.Exec("update t_activity set activity_type=?,apply_zone=?,apply_user=?,configure_data=?,configure_desc=? WHERE activity_id=?",activity_type, apply_zone, apply_user, config_data, config_desc, activity_id)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

    return ctx.SendResponse("更新活动成功")
}

func ActivityDel(c echo.Context) error {
    ctx := c.(*mid.Context)
    ids := ctx.FormValue("idsStr")

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_zone_global")
	if err != nil {
		return err
	}

	_, err = tx.Exec("Delete FROM t_activity WHERE activity_id IN ("+ids+")")
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
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
    apply_user := ctx.FormValue("apply_user")
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
        return ctx.SendError(-1, "参数非法")
    }

    sql := "insert into t_activity(activity_id,activity_type,apply_zone,apply_user,configure_data,configure_desc) values"
    for k, v := range acts {
        if k != 0 {
            sql += ","
        }
        sql += fmt.Sprintf("(%d,%d,'%s','%s','%s','%s')", v.Id, v.Type, apply_zone, apply_user, v.Data, v.Desc)
    }

	db := common.GetDb()
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE "+common.GetDbPrefix()+"db_zone_global")
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
