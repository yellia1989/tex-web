package game

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-web/backend/api/gm"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
	"strconv"
	"strings"
	"time"
)

type role struct {
	AccountName   string `json:"id"`
	Actorid       uint32 `json:"actorid"`
	Name          string `json:"name"`
	VipLevel      uint32 `json:"vip_level"`
	LastLoginTime string `json:"last_login_time"`
	RegTime       string `json:"reg_time"`
	Zoneid        uint32 `json:"zoneid"`
}

func RoleList(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneid := ctx.QueryParam("zoneid")
	name := strings.TrimSpace(ctx.QueryParam("name"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	field := ctx.QueryParam("field")
	order := ctx.QueryParam("order")

	zoneid2, _ := strconv.Atoi(zoneid)
	err, conn := gm.GameDb(uint32(zoneid2))
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_zone_" + zoneid)
	if err != nil {
		return err
	}

	sql := "SELECT accountid,actorid,actorname,vip_level,lastonlinetime2,createtime,serverindex FROM actors"
	if name != "" {
		accountid, err := strconv.Atoi(name)
		if err == nil && accountid != 0 {
			sql += fmt.Sprintf(" WHERE accountid = %d or actorid = %d", accountid, accountid)
		} else {
			sql += " WHERE actorname like '%" + name + "%'"
		}
	}
	if field != "" {
		sql += " ORDER BY " + field + " " + order
	}
	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	log.Infof("sql: %s", sql)

	roles := make([]role, 0)
	for rows.Next() {
		var r role
		var tmp int64
		if err := rows.Scan(&r.AccountName, &r.Actorid, &r.Name, &r.VipLevel, &tmp, &r.RegTime, &r.Zoneid); err != nil {
			return err
		}
		r.LastLoginTime = common.FormatTimeInLocal("2006-01-02 15:04:05", time.Unix(tmp, 0))
		roles = append(roles, r)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	vPage := common.GetPage(roles, page, limit)
	return ctx.SendArray(vPage, len(roles))
}

func RoleDeatil(c echo.Context) error {
	ctx := c.(*mid.Context)
	zoneId := ctx.QueryParam("zoneId")
	roleId := ctx.QueryParam("roleId")

	if zoneId == "" || roleId == "" {
		return ctx.SendError(-1, "参数非法")
	}

	comm := cfg.Comm
	app := cfg.App

	gamePrx := new(rpc.GameService)
	comm.StringToProxy(app+".GameServer.GameServiceObj%"+app+".zone."+zoneId, gamePrx)

	result := ""
	var ret int32
	var err error
	ret, err = gamePrx.DoGmCmd("admin", "see_json "+roleId, &result)
	if ret != 0 || err != nil {
		sErr := ""
		if err != nil {
			sErr = err.Error()
		}
		result = fmt.Sprintf("ret:%d, err:%s", ret, sErr)
	}

	return ctx.SendResponse(result)
}
