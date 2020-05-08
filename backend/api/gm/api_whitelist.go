package gm

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

func WhiteList(c echo.Context) error {
	ctx := c.(*mid.Context)

	db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16)/db_loginserver")
	defer db.Close()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT * FROM t_whitelist;")
	if err != nil {
		return err
	}
	defer rows.Close()

    var strs []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
        strs = append(strs, id)
	}

	return ctx.SendResponse(strings.Join(strs, ";"))
}

func WhiteAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		notice := rpc.NoticeDataInfo{}
		if err := ctx.Bind(&notice); err != nil {
			return err
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		ret, err := bulletinPrx.AddNotice(notice)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	*/

	return ctx.SendResponse("添加白名单用户成功")
}

func WhiteDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		ids := strings.Split(ctx.FormValue("idsStr"), ",")

		if len(ids) == 0 {
			return ctx.SendError(-1, "白名单用户不存在")
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		for _, id := range ids {
			id, _ := strconv.ParseUint(id, 10, 32)
			ret, err := bulletinPrx.DelNotice(uint32(id))
			if err := checkRet(ret, err); err != nil {
				return err
			}
		}
	*/

	return ctx.SendResponse("删除白名单用户成功")
}

func WhiteReplace(c echo.Context) error {
	ctx := c.(*mid.Context)

	/*
		notice := rpc.NoticeDataInfo{}
		if err := ctx.Bind(&notice); err != nil {
			return err
		}

		bulletinPrx := new(rpc.BulletinService)
		comm.StringToProxy("aqua.BulletinServer.BulletinServiceObj", bulletinPrx)

		ret, err := bulletinPrx.ModifyNotice(notice)
		if err := checkRet(ret, err); err != nil {
			return err
		}
	*/

	return ctx.SendResponse("覆盖白名单用户成功")
}
