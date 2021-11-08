package server

import (
	dsql "database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type templateData struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Parent  string `json:"parent"`
}

func TemplateList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	name := strings.TrimSpace(ctx.QueryParam("name"))
	parent := strings.TrimSpace(ctx.QueryParam("parent"))

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("USE " + cfg.GameDbPrefix + "db_tex")
	if err != nil {
		return err
	}

	sql := "SELECT name, content, parent FROM t_template"
	where := ""
	if name != "" {
		where += "name = '" + name + "'"
	}
	if parent != "" {
		if where == "" {
			where += "parent = '" + parent + "'"
		} else {
			where += " AND parent = '" + parent + "'"
		}
	}
	if where != "" {
		sql += " WHERE " + where
	}
	var total int
	err = tx.QueryRow("SELECT count(*) as total FROM (" + sql + ") a").Scan(&total)
	if err != nil {
		return err
	}

	limitstart := strconv.Itoa((page - 1) * limit)
	limitrow := strconv.Itoa(limit)
	sql += " LIMIT " + limitstart + "," + limitrow

	c.Logger().Debug(sql)

	rows, err := tx.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]templateData, 0)
	for rows.Next() {
		var r templateData
		var content, parent dsql.NullString
		if err := rows.Scan(&r.Name, &content, &parent); err != nil {
			return err
		}
		if content.Valid {
			r.Content = content.String
		}
		if parent.Valid {
			r.Parent = parent.String
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

func DelTemplate(c echo.Context) error {
	ctx := c.(*mid.Context)
	ids := ctx.FormValue("idsStr")

	if len(ids) == 0 {
		return ctx.SendError(-1, "模板不存在")
	}

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	_, err := db.Exec("USE " + cfg.GameDbPrefix + "db_tex")
	if err != nil {
		return err
	}

	sql := "DELETE FROM t_template WHERE name IN" + ids
	c.Logger().Debug(sql)

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return ctx.SendResponse("删除模板成功")
}

func TemplateDatail(c echo.Context) error {
	ctx := c.(*mid.Context)

	name := ctx.FormValue("name")
	if name == "" {
		return ctx.SendError(-1, "参数非法")
	}

	comm := cfg.Comm
	patchPrx := new(rpc.Patch)
	comm.StringToProxy("tex.mfwpatch.PatchObj", patchPrx)

    content := ""
	ret, err := patchPrx.GetTemplate(name, &content)
	if ret != 0 || err != nil {
		if err != nil {
			return fmt.Errorf("opt failed, ret:%d, err:%s", ret, err.Error())
		} else {
			return fmt.Errorf("opt failed, ret:%d", ret)
		}
	}

	return ctx.SendResponse(content)
}

func TemplateUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

	name := ctx.FormValue("name")
	parent := ctx.FormValue("parent")
	content := ctx.FormValue("content")

	if name == "" || parent == "" || content == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.GameGlobalDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	_, err := db.Exec("USE " + cfg.GameDbPrefix + "db_tex")
	if err != nil {
		return err
	}

	sql := "UPDATE t_template SET parent = " + parent + ", content = " + content
	sql += "WHERE name = " + name
	c.Logger().Debug(sql)

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return ctx.SendResponse("更新模板成功")
}
