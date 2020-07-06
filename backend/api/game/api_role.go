package game

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	tex "github.com/yellia1989/tex-go/service"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

var (
	comm = tex.NewCommunicator("tex.mfwregistry.QueryObj@tcp -h 192.168.0.16 -p 2000 -t 3600000")
)

type _role struct {
    Uid uint64  `json:"id"`
    Name string `json:"name"`
    VipLevel uint32 `json:"vip_level"`
    LastLoginTime string `json:"last_login_time"`
    RegTime string `json:"reg_time"`
}

func RoleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    zoneid := ctx.QueryParam("zoneid")
    name := ctx.QueryParam("name")
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
    field := ctx.QueryParam("field")
    order := ctx.QueryParam("order")

    db := common.GetDb()
    if db == nil {
        return ctx.SendError(-1, "连接数据库失败");
    }

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("USE db_zone_" + zoneid)
    if err != nil {
        return err
    }

    sql := "SELECT uid,name,vip_level,last_login_time,reg_time FROM t_role"
    if name != "" {
        sql += " WHERE name like '%" + name + "%'"
    }
    if field != "" {
        sql += " ORDER BY " + field + " " + order
    }
    rows, err := tx.Query(sql)
    if err != nil {
        return err
    }
    defer rows.Close()

    roles := make([]_role, 0)
    for rows.Next() {
        var r _role
        if err := rows.Scan(&r.Uid, &r.Name, &r.VipLevel, &r.LastLoginTime, &r.RegTime); err != nil {
            return err
        }
        roles = append(roles, r)
    }
    if err := rows.Err(); err != nil {
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }

    vPage := common.GetPage(roles, page, limit)
    return ctx.SendArray(vPage, len(roles))
}

type _heroData struct {
	HeroID uint32 `json:"iHeroId"`
	Level uint32 `json:"level"`
	Star uint32 `json:"star"`
}

func getRoleDetail(zone, role string) []byte {
	gamePrx := new(rpc.GameService)
	comm.StringToProxy("aqua.GameServer.GameServiceObj%aqua.zone."+zone, gamePrx)

	result := ""
	buff := bytes.Buffer{}
	var ret int32
	var err error
	ret, err = gamePrx.DoGmCmd("admin", "see_json "+role, &result)
	if ret != 0 || err != nil {
		sErr := ""
		if err != nil {
			sErr = err.Error()
		}
		result = fmt.Sprintf("ret:%d, err:%s", ret, sErr)
	}

	buff.WriteString(result+"\n")
	return buff.Bytes()
}

func RoleDeatil(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneId := ctx.QueryParam("zoneId")
	roleId := ctx.QueryParam("roleId")

	if zoneId == "" || roleId == "" {
		return ctx.SendError(-1, "参数非法")
	}

	result := getRoleDetail(zoneId, roleId)

	return ctx.SendResponse(result)
}
