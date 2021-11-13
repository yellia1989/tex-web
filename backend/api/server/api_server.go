package server

import (
	"crypto/md5"
	dsql "database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	tex "github.com/yellia1989/tex-go/service"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type nodeData struct {
	Node          string  `json:"node"`
	Obj           string  `json:"obj"`
	SettingStat   int     `json:"setting_stat"`
	CurStat       int     `json:"cur_stat"`
	HeartbeatTime string  `json:"heartbeat_time"`
	LoadAvg1      float32 `json:"loadavg1"`
	LoadAvg5      float32 `json:"loadavg5"`
	LoadAvg15     float32 `json:"loadavg15"`
}
type ServerData struct {
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

type ServiceData struct {
	Service      string `json:"service"`
	Port         int    `json:"port"`
	PortType     string `json:"port_type"`
	ThreadNum    int    `json:"thread_num"`
	Protocol     string `json:"protocol"`
	MaxConn      int    `json:"max_conn"`
	QueueCap     int    `json:"queue_cap"`
	QueueTimeout int    `json:"queue_timeout"`
}

type ServerDetailData struct {
	ServerData
	Services []ServiceData `json:"services"`
}

type patchData struct {
	Id         int    `json:"id"`
	Remark     string `json:"remark"`
	Version    string `json:"version"`
	Server     string `json:"server"`
	File       string `json:"file"`
	Md5        string `json:"md5"`
	UploadTime string `json:"upload_time"`
}

func NodeList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	var vParam []interface{}

	sql := "SELECT name, obj, setting_stat, cur_stat, heartbeat_time, load_avg1, load_avg5, load_avg15 FROM t_node_info where 1=1"

	var total int
	err := db.QueryRow("SELECT count(*) from t_node_info where 1=1", vParam...).Scan(&total)
	if err != nil {
		return err
	}

	sql += " LIMIT ?,?"

	limitstart := (page - 1) * limit
	vParam = append(vParam, limitstart)
	vParam = append(vParam, limit)

	c.Logger().Debug(sql)

	rows, err := db.Query(sql, vParam...)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]nodeData, 0)
	for rows.Next() {
		var r nodeData
		if err := rows.Scan(&r.Node, &r.Obj, &r.SettingStat, &r.CurStat, &r.HeartbeatTime, &r.LoadAvg1, &r.LoadAvg5, &r.LoadAvg15); err != nil {
			return err
		}
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func ServerList(c echo.Context) error {
	ctx := c.(*mid.Context)
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	app := strings.TrimSpace(ctx.QueryParam("app"))
	server := strings.TrimSpace(ctx.QueryParam("server"))
	division := strings.TrimSpace(ctx.QueryParam("division"))
	node := strings.TrimSpace(ctx.QueryParam("node"))

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	var vParam []interface{}

	sql := "SELECT app, server, division, node, setting_stat, cur_stat, profile_conf_template, template_name, pid FROM t_server WHERE 1=1"
	where := ""
	if app != "" {
		where += " AND app = ?"
		vParam = append(vParam, app)
	}
	if server != "" {
		where += " AND server = ?"
		vParam = append(vParam, server)
	}
	if division != "" {
		where += " AND division = ?"
		vParam = append(vParam, division)
	}
	if node != "" {
		where += " AND node = ?"
		vParam = append(vParam, node)
	}
	sql += where

	var total int
	err := db.QueryRow("SELECT count(*) FROM t_server WHERE 1=1"+where, vParam...).Scan(&total)
	if err != nil {
		return err
	}

	sql += " LIMIT ?,?"

	limitstart := (page - 1) * limit
	vParam = append(vParam, limitstart)
	vParam = append(vParam, limit)

	c.Logger().Debug(sql)

	rows, err := db.Query(sql, vParam...)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]ServerData, 0)
	for rows.Next() {
		var r ServerData
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
	req.STaskNo = uuid.NewString()
	err := json.Unmarshal([]byte(sVItem), &req.VItem)
	if err != nil {
		return err
	}
	for i := 0; i < len(req.VItem); i++ {
		v := &req.VItem[i]
		v.STaskNo = uuid.NewString()
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

func ServerDetail(c echo.Context) error {
	ctx := c.(*mid.Context)

	app := ctx.FormValue("app")
	server := ctx.FormValue("server")
	division := ctx.FormValue("division")
	node := ctx.FormValue("node")

	if app == "" || server == "" || division == "" || node == "" {
		ctx.SendError(-1, "参数非法")
	}

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT app, server, division, node, setting_stat, cur_stat, profile_conf_template, template_name, pid FROM t_server"
	where := " WHERE app = '" + app + "' AND server = '" + server + "' AND node = '" + node + "' AND division = '" + division + "'"
	sql += where

	c.Logger().Debug(sql)
	data := ServerDetailData{
		Services: make([]ServiceData, 0, 1),
	}

	row := db.QueryRow(sql)
	var profile dsql.NullString
	err := row.Scan(&data.App, &data.Server, &data.Division, &data.Node, &data.SettingStat, &data.CurStat, &profile, &data.TemplateName, &data.Pid)
	if err != nil {
		return err
	}
	if profile.Valid {
		data.ProfileConfTemplate = profile.String
	}

	sql2 := "SELECT service, endpoint, thread_num, protocol, max_conn, queue_cap, queue_timeout FROM t_service"
	sql2 += where
	rows2, err := db.Query(sql2)
	if err != nil {
		return err
	}
	defer rows2.Close()

	for rows2.Next() {
		sEndpoint := ""
		r := ServiceData{}
		if err := rows2.Scan(&r.Service, &sEndpoint, &r.ThreadNum, &r.Protocol, &r.MaxConn, &r.QueueCap, &r.QueueTimeout); err != nil {
			return err
		}

		endpoint, err := tex.NewEndpoint(sEndpoint)
		if err != nil {
			return err
		}
		r.Port = endpoint.Port
		r.PortType = endpoint.Proto
		data.Services = append(data.Services, r)
	}

	return ctx.SendResponse(data)
}

func ServerUpdate(c echo.Context) error {
	ctx := c.(*mid.Context)

	sServer := ctx.FormValue("server")
	req := ServerDetailData{}

	json.Unmarshal([]byte(sServer), &req)

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	tx, err := db.Begin()
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}
	defer tx.Rollback()

	sql := "UPDATE t_server SET setting_stat = ?, template_name = ? WHERE app = ? AND server = ? AND division = ? AND node = ?"
	c.Logger().Debug(sql)

	_, err = tx.Exec(sql, req.SettingStat, req.TemplateName, req.App, req.Server, req.Division, req.Node)
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}

	for _, v := range req.Services {
		endpoint := tex.Endpoint{
			Proto:       v.PortType,
			IP:          req.Node,
			Port:        v.Port,
			Idletimeout: time.Duration(v.QueueTimeout) * time.Millisecond,
		}
		sEndpoint := endpoint.String()
		sql2 := "INSERT INTO t_service (app, server, division, node, service, endpoint, thread_num, protocol, max_conn, queue_cap, queue_timeout) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE endpoint = ?, thread_num = ?, protocol = ?, max_conn = ?, queue_cap = ?, queue_timeout = ?"

		c.Logger().Debug(sql2)
		if _, err := tx.Exec(sql2, req.App, req.Server, req.Division, req.Node, v.Service, sEndpoint, v.ThreadNum, v.Protocol, v.MaxConn, v.QueueCap, v.QueueTimeout, sEndpoint, v.ThreadNum, v.Protocol, v.MaxConn, v.QueueCap, v.QueueTimeout); err != nil {
			return ctx.SendError(-1, err.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		return ctx.SendError(-1, err.Error())
	}

	return ctx.SendResponse("更新服务成功")
}

func ServerAdd(c echo.Context) error {
	ctx := c.(*mid.Context)

	sServer := ctx.FormValue("server")
	req := ServerDetailData{}

	json.Unmarshal([]byte(sServer), &req)

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}
	tx, err := db.Begin()
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}
	defer tx.Rollback()

	sql := "INSERT INTO t_server (app, server, division, node, setting_stat, template_name) VALUES (?,?,?,?,?,?)"
	c.Logger().Debug(sql)

	_, err = tx.Exec(sql, req.App, req.Server, req.Division, req.Node, req.SettingStat, req.TemplateName)
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}

	sql2 := "INSERT INTO t_service (app, server, division, node, service, endpoint, thread_num, protocol, max_conn, queue_cap, queue_timeout) VALUES "
	values := make([]string, 0, 1)
	params := make([]interface{}, 0, 11)
	for _, v := range req.Services {
		endpoint := tex.Endpoint{
			Proto:       v.PortType,
			IP:          req.Node,
			Port:        v.Port,
			Idletimeout: time.Duration(v.QueueTimeout) * time.Millisecond,
		}
		str := "(?,?,?,?,?,?,?,?,?,?,?)"
		values = append(values, str)
		params = append(params, req.App, req.Server, req.Division, req.Node, v.Service, endpoint.String(), v.ThreadNum, v.Protocol, v.MaxConn, v.QueueCap, v.QueueTimeout)
	}
	sql2 += strings.Join(values, ",") + ";"
	c.Logger().Debug(sql2, params)

	_, err = tx.Exec(sql2, params...)
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}
	if err := tx.Commit(); err != nil {
		return ctx.SendError(-1, err.Error())
	}

	return ctx.SendResponse("添加服务成功")
}

func ServerDel(c echo.Context) error {
	ctx := c.(*mid.Context)

	sServers := ctx.FormValue("servers")
	if sServers == "" {
		return ctx.SendError(-1, "参数非法")
	}

	servers := make([]ServerData, 0, 1)
	err := json.Unmarshal([]byte(sServers), &servers)
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	tx, err := db.Begin()
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}
	defer tx.Rollback()

	for _, v := range servers {
		sql := "DELETE FROM t_server WHERE app = ? AND server = ? AND division = ? AND node = ?"
		c.Logger().Debug(sql)
		_, err = tx.Exec(sql, v.App, v.Server, v.Division, v.Node)
		if err != nil {
			return err
		}

		sql2 := "DELETE FROM t_service WHERE app = ? AND server = ? AND division = ? AND node = ?"
		c.Logger().Debug(sql2)
		_, err = tx.Exec(sql2, v.App, v.Server, v.Division, v.Node)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return ctx.SendError(-1, err.Error())
	}

	return ctx.SendResponse("删除服务成功")
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

	md5Str := fmt.Sprintf("%x", md5hash.Sum(nil))

	row := db.QueryRow("select id from t_patch where md5 = ?", md5Str)
	if row.Scan() != dsql.ErrNoRows {
		return ctx.SendError(-1, "文件已存在")
	}

	_, err = src.Seek(0, 0)
	if err != nil {
		return err
	}

	dst, err := os.Create(path.Join(cfg.UploadPatchPrefix, file.Filename))
	defer dst.Close()

	if err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	_, err = db.Exec("insert into t_patch(server,file,md5,remark,version) values(?,?,?,?,?)", server, file.Filename, md5Str, remark, version)
	if err != nil {
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
	err := db.QueryRow("select file from t_patch where id=?", id).Scan(&fileName)
	if err != nil {
		return err
	}

	filePath := path.Join(cfg.UploadPatchPrefix, fileName)
	exist, err := common.PathExists(filePath)
	if err != nil {
		return err
	}

	if exist {
		return ctx.Attachment(filePath, fileName)
	} else {
		return ctx.SendError(-1, "文件不存在")
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
	err := db.QueryRow("select file from t_patch where id=?", id).Scan(&fileName)
	if err != nil {
		return err
	}

	filePath := path.Join(cfg.UploadPatchPrefix, fileName)

	exist, err := common.PathExists(filePath)
	if err != nil {
		return err
	}

	if exist {
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("delete from t_patch where id=?", id)
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
	err := db.QueryRow("select count(id) from t_patch where 1=1"+where, vParam...).Scan(&total)
	if err != nil {
		return err
	}

	if limit != 0 && page != 0 {
		limitstart := strconv.Itoa((page - 1) * limit)
		limitrow := strconv.Itoa(limit)
		sql += " limit ?,?"
		vParam = append(vParam, limitstart)
		vParam = append(vParam, limitrow)
	}

	c.Logger().Debug(sql)

	rows, err := db.Query(sql, vParam...)
	if err != nil {
		return err
	}
	defer rows.Close()

	logs := make([]patchData, 0)
	for rows.Next() {
		var r patchData
		if err := rows.Scan(&r.Id, &r.Server, &r.File, &r.Md5, &r.UploadTime, &r.Remark, &r.Version); err != nil {
			return err
		}
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}
