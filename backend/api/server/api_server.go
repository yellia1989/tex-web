package server

import (
	"crypto/md5"
	dsql "database/sql"
	"encoding/json"
	"fmt"
	"github.com/yellia1989/tex-web/backend/common"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type serverData struct {
	App                 string `json:"app"`
	Server              string `json:"server"`
	Division            string `json:"division"`
	Node                string `json:"node"`
	SettingStat         int    `json:"setting_stat"`
	CurStat             int    `json:"cur_stat"`
	ProfileConfTemplate string `json:"profile_conf_template"`
	TemplateName        string `json:"template_name"`
	Pid                 int    `json:"pid"`
}

type patchData struct {
	Id 					int	   `json:"id"`
	Remark				string `json:"remark"`
	Version				string `json:"version"`
	Server              string `json:"server"`
	File	            string `json:"file"`
	Md5                 string `json:"md5"`
	UploadTime          string `json:"upload_time"`
}

func ServerList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	app := strings.TrimSpace(ctx.QueryParam("app"))
	server := strings.TrimSpace(ctx.QueryParam("server"))

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	var vParam []interface{}

	sql := "SELECT app, server, division, node, setting_stat, cur_stat, profile_conf_template, template_name, pid FROM t_server where 1=1"
	where := ""
	if app != "" {
		where += " and app = ?"
		vParam = append(vParam, app)
	}
	if server != "" {
		where += " and server = ?"
		vParam = append(vParam,server)
	}
	if where != "" {
		sql += where
	}
	var total int
	err := db.QueryRow("SELECT count(*) from t_server where 1=1" + where,vParam...).Scan(&total)
	if err != nil {
		return err
	}

	sql += " LIMIT ?,?"

	limitstart := (page - 1) * limit
	vParam = append(vParam, limitstart)
	vParam = append(vParam, limit)

	c.Logger().Debug(sql)

	rows, err := db.Query(sql,vParam...)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]serverData, 0)
	for rows.Next() {
		var r serverData
		var profile dsql.NullString
		if err := rows.Scan(&r.App, &r.Server, &r.Division, &r.Node, &r.SettingStat, &r.CurStat, &profile, &r.TemplateName, &r.Pid); err != nil {
			return err
		}
		if profile.Valid {
			r.ProfileConfTemplate = profile.String
		}
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func ServerOperator(c echo.Context) error {
	ctx := c.(*mid.Context)

	sVItem := ctx.FormValue("vItem")

	if sVItem == "" {
		return ctx.SendError(-1, "参数为空")
	}

    req := rpc.PatchTaskReq{}
	req.STaskNo = uuid.NewString();
    err := json.Unmarshal([]byte(sVItem), &req.VItem)
    if err != nil {
        return err
    }
    for i := 0; i < len(req.VItem); i++ {
        v := &req.VItem[i]
        v.STaskNo = uuid.NewString()
        v.SNodeName = "192.168.0.16"
    }
    fmt.Printf("%v", req)

	comm := cfg.Comm
	patchPrx := new(rpc.Patch)
	comm.StringToProxy("tex.mfwpatch.PatchObj", patchPrx)

	ret, err := patchPrx.AddTask(req)
	if ret != 0 || err != nil {
        if err != nil {
            return fmt.Errorf("opt failed, ret:%d, err:%s", ret, err.Error())
        } else {
            return fmt.Errorf("opt failed, ret:%d", ret)
        }
	}

	return ctx.SendResponse(req.STaskNo)
}

func GetTask(c echo.Context) error {
	ctx := c.(*mid.Context)

	taskNo := ctx.FormValue("taskNo")
	if taskNo == "" {
		return ctx.SendError(-1, "参数非法")
	}

	comm := cfg.Comm
	patchPrx := new(rpc.Patch)
	comm.StringToProxy("tex.mfwpatch.PatchObj", patchPrx)

	taskRsp := rpc.NewPatchTaskRsp()
	ret, err := patchPrx.GetTask(taskNo, taskRsp)
	if ret != 0 || err != nil {
        if err != nil {
            return fmt.Errorf("opt failed, ret:%d, err:%s", ret, err.Error())
        } else {
            return fmt.Errorf("opt failed, ret:%d", ret)
        }
	}

	return ctx.SendResponse(taskRsp)
}

func UploadPatch(c echo.Context) error {
	ctx := c.(*mid.Context)
	server := ctx.FormValue("server")
	remark := ctx.FormValue("remark")
	version := ctx.FormValue("version")
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	if server == "" || version == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	//打开用户上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	md5hash := md5.New()
	if _, err = io.Copy(md5hash, src); err != nil {
		return err
	}

	md5Str := fmt.Sprintf("%x",md5hash.Sum(nil))

	row := db.QueryRow("select id from t_patch where md5 = ?",md5Str)
	if row.Scan() != dsql.ErrNoRows {
		return ctx.SendError(-1, "文件已存在")
	}

	_,err = src.Seek(0,0)
	if err != nil {
		return err
	}

	dst,err := os.Create(path.Join(cfg.UploadPatchPrefix,file.Filename))
	defer dst.Close()

	if err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	_,err = db.Exec("insert into t_patch(server,file,md5,remark,version) values(?,?,?,?,?)",server,file.Filename,md5Str,remark,version)
	if err != nil{
		return ctx.SendError(-1, err.Error())
	}

	return ctx.SendResponse("上传成功")
}

func DownloadPatch(c echo.Context) error {
	ctx := c.(*mid.Context)
	id := ctx.QueryParam("id")
	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	var fileName string
	err := db.QueryRow("select file from t_patch where id=?",id).Scan(&fileName)
	if err != nil {
		return err
	}

	filePath := path.Join(cfg.UploadPatchPrefix,fileName)
	exist,err := common.PathExists(filePath)
	if err != nil {
		return err
	}

	if exist{
		return ctx.Attachment(filePath,fileName)
	} else {
		return ctx.SendError(-1,"文件不存在")
	}
}

func DeletePatch(c echo.Context) error {
	ctx := c.(*mid.Context)
	id := ctx.QueryParam("id")
	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	var fileName string
	err := db.QueryRow("select file from t_patch where id=?",id).Scan(&fileName)
	if err != nil {
		return err
	}

	filePath:= path.Join(cfg.UploadPatchPrefix,fileName)

	exist,err := common.PathExists(filePath)
	if err != nil {
		return err
	}

	if exist {
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	_,err = db.Exec("delete from t_patch where id=?",id)
	if err != nil {
		return err
	}

	return ctx.SendResponse("删除成功")
}

func PatchList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	server := strings.TrimSpace(ctx.QueryParam("server"))

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	var vParam []interface{}
	sql := "select id,server, file, md5, upload_time,remark,version from t_patch where 1=1"
	where := ""
	if server != "" {
		where += " and server = ?"
		vParam = append(vParam, server)
	}
	sql += where
	var total int
	err := db.QueryRow("select count(id) from t_patch where 1=1" + where,vParam...).Scan(&total)
	if err!= nil{
		return err
	}

	if limit !=0 && page !=0 {
		limitstart := strconv.Itoa((page - 1) * limit)
		limitrow := strconv.Itoa(limit)
		sql += " limit ?,?"
		vParam = append(vParam, limitstart)
		vParam = append(vParam, limitrow)
	}

	c.Logger().Debug(sql)

	rows, err := db.Query(sql,vParam...)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]patchData, 0)
	for rows.Next() {
		var r patchData
		if err := rows.Scan(&r.Id, &r.Server, &r.File, &r.Md5, &r.UploadTime,&r.Remark,&r.Version); err != nil {
			return err
		}
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}