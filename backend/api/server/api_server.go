package server

import (
	"crypto/md5"
	dsql "database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	_ "golang.org/x/crypto/ssh/terminal"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	tex "github.com/yellia1989/tex-go/service"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
    frpc "github.com/yellia1989/tex-go/sdp/rpc"
)

type nodeData struct {
	Node          string  `json:"node"`
	Obj           string  `json:"obj"`
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
	CurStat             int    `json:"cur_stat"`
	AutoStart     int     `json:"auto_start"`
	ProfileConfTemplate string `json:"profile_conf_template"`
	TemplateName        string `json:"template_name"`
	Pid                 int    `json:"pid"`
	PublishVersion      string `json:"publish_version"`
	PublishUserName     string `json:"publish_username"`
	PublishTime         string `json:"publish_time"`
	PromPort            int    `json:"prom_port"`
	ManualStop    int     `json:"manual_stop"`
    MfwServer   int `json:"mfw_server"`
    StartScript string `json:"start_script"`
    MonitorScript string `json:"monitor_script"`
    StopScript string `json:"stop_script"`
    StartTime   int `json:"start_time"`
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
    Timeout      int    `json:"timeout"`
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
	Default    int    `json:"default"`
}

type wsWrapper struct {
	*websocket.Conn
}

func (wsw *wsWrapper) Write(p []byte) (n int, err error) {
	writer, err := wsw.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	defer writer.Close()
	return writer.Write(p)
}

func (wsw *wsWrapper) Read(p []byte) (n int, err error) {
	for {
		msgType, reader, err := wsw.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}
}

var upgrader = websocket.Upgrader{}

func websocketHandle(con *websocket.Conn, ip string, cols string, rows string, user string, key string) {
	rw := io.ReadWriter(&wsWrapper{con})
	webprintln := func(data string) {
		rw.Write([]byte(data + "\r\n"))
	}
	con.SetCloseHandler(func(code int, text string) error {
		con.Close()
		return nil
	})

	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		webprintln(err.Error())
		return
	}

	sshConfig := &ssh.ClientConfig{
		User:    user,
		Auth:    []ssh.AuthMethod{ssh.PublicKeys(signer)},
		Timeout: 10 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	ip = ip + ":22"
	client, err := ssh.Dial("tcp", ip, sshConfig)
	if err != nil {
		webprintln(err.Error())
		return
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		webprintln(err.Error())
		return
	}
	defer session.Close()
	session.Stdout = rw
	session.Stderr = rw
	session.Stdin = rw
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	termWidth, _ := strconv.Atoi(cols)
	termHeight, _ := strconv.Atoi(rows)
	err = session.RequestPty("xterm", termHeight, termWidth, modes)
	if err != nil {
		webprintln(err.Error())
		return
	}
	err = session.Shell()
	if err != nil {
		webprintln(err.Error())
		return
	}
	err = session.Wait()
	if err != nil {
		webprintln(err.Error())
		return
	}
}

func ShellWs(c echo.Context) error {
	ctx := c.(*mid.Context)
	node := ctx.QueryParam("node")
	cols := ctx.QueryParam("cols")
	rows := ctx.QueryParam("rows")

	u := ctx.GetUser()
	if u == nil {
		return ctx.SendError(-1, "账号不存在")
	}

	if u.TerminalUser == "" || u.TerminalKey == "" {
		return ctx.SendError(-1, "ssh账号不存在")
	}

	con, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		c.Logger().Error("Not a websocket connection")
		return ctx.SendError(-1, "Not a websocket handshake")
	} else if err != nil {
		return ctx.SendError(-1, err)
	}
	go websocketHandle(con, node, cols, rows, u.TerminalUser, u.TerminalKey)
	return ctx.SendResponse("")
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

	sql := "SELECT name, obj, cur_stat, heartbeat_time, load_avg1, load_avg5, load_avg15 FROM t_node_info where 1=1"

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
		if err := rows.Scan(&r.Node, &r.Obj, &r.CurStat, &r.HeartbeatTime, &r.LoadAvg1, &r.LoadAvg5, &r.LoadAvg15); err != nil {
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

	sql := "SELECT app, server, division, node, auto_start, cur_stat, profile_conf_template, template_name, pid, publish_version, publish_username, publish_time, prom_port,manual_stop,mfw_server,start_script,monitor_script,stop_script,start_time FROM t_server WHERE 1=1"
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
		var profile, version, username, publishTime dsql.NullString
		if err := rows.Scan(&r.App, &r.Server, &r.Division, &r.Node, &r.AutoStart, &r.CurStat, &profile, &r.TemplateName, &r.Pid, &version, &username, &publishTime, &r.PromPort, &r.ManualStop, &r.MfwServer, &r.StartScript, &r.MonitorScript, &r.StopScript, &r.StartTime); err != nil {
			return err
		}
		if profile.Valid {
			r.ProfileConfTemplate = profile.String
		}
		if version.Valid {
			r.PublishVersion = version.String
		}
		if username.Valid {
			r.PublishUserName = username.String
		}
		if publishTime.Valid {
			r.PublishTime = publishTime.String
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

	user := ctx.GetUser()

	req := rpc.PatchTaskReq{}
	req.STaskNo = uuid.NewString()
	req.SUsername = user.UserName
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

	if app == "" || server == "" || node == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	sql := "SELECT app, server, division, node, auto_start, cur_stat, profile_conf_template, template_name, pid, mfw_server, start_script, monitor_script, stop_script FROM t_server"
	where := " WHERE app = '" + app + "' AND server = '" + server + "' AND node = '" + node + "' AND division = '" + division + "'"
	sql += where

	c.Logger().Debug(sql)

	data := ServerDetailData{
		Services: make([]ServiceData, 0, 1),
	}

	row := db.QueryRow(sql)
	var profile dsql.NullString
	err := row.Scan(&data.App, &data.Server, &data.Division, &data.Node, &data.AutoStart, &data.CurStat, &profile, &data.TemplateName, &data.Pid, &data.MfwServer, &data.StartScript, &data.MonitorScript, &data.StopScript)
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
        r.Timeout = int(endpoint.Idletimeout.Milliseconds())
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

	sql := "UPDATE t_server SET auto_start = ?, template_name = ?, profile_conf_template = ?, mfw_server = ?, prom_port = ?, start_script = ?, monitor_script = ?, stop_script = ? WHERE app = ? AND server = ? AND division = ? AND node = ?"
	c.Logger().Debug(sql)

	_, err = tx.Exec(sql, req.AutoStart, req.TemplateName, req.ProfileConfTemplate, req.MfwServer, req.PromPort, req.StartScript, req.MonitorScript, req.StopScript, req.App, req.Server, req.Division, req.Node)
	if err != nil {
		return ctx.SendError(-1, err.Error())
	}

	for _, v := range req.Services {
		endpoint := tex.Endpoint{
			Proto:       v.PortType,
			IP:          req.Node,
			Port:        v.Port,
			Idletimeout: time.Duration(v.Timeout) * time.Millisecond,
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

	sql := "INSERT INTO t_server (app, server, division, node, auto_start, template_name, profile_conf_template, mfw_server, start_script, monitor_script, stop_script) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	c.Logger().Debug(sql)

	_, err = tx.Exec(sql, req.App, req.Server, req.Division, req.Node, req.AutoStart, req.TemplateName, req.ProfileConfTemplate, req.MfwServer, req.StartScript, req.MonitorScript, req.StopScript)
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
			Idletimeout: time.Duration(v.Timeout) * time.Millisecond,
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
	sql := "select id,server,file,md5,upload_time,remark,version,def from t_patch where 1=1"
	where := ""
	if server != "" {
		where += " and server = ?"
		vParam = append(vParam, server)
	}
	sql += where
    sql += " order by upload_time desc"
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
		if err := rows.Scan(&r.Id, &r.Server, &r.File, &r.Md5, &r.UploadTime, &r.Remark, &r.Version, &r.Default); err != nil {
			return err
		}
		logs = append(logs, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.SendArray(logs, total)
}

func RegistryIp() (map[uint32]string,error) {
    comm := cfg.Comm

    queryPrx := new(frpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    var vObj []frpc.ObjEndpoint
    ret, err := queryPrx.GetAllEndpoints(&vObj)
    if ret != 0 || err != nil {
        return nil, fmt.Errorf("ret: %d, err: %s", ret, err.Error())
    }

    m := make(map[uint32]string)
    for _, o := range vObj {
        if len(o.SDivision) == 0 {
            continue
        }
        vzone := strings.Split(o.SDivision, ".")
        if len(vzone) != 3 || vzone[1] != "zone" {
            continue
        }
        vep := strings.Split(o.SEp, " ")
        if len(vep) != 7 {
            continue
        }
        m[common.Atou32(vzone[2])] = vep[2]
    }

    return m, nil
}

func AllocPromPort(c echo.Context) error {
	ctx := c.(*mid.Context)
	node := ctx.QueryParam("node")

    if node == "" {
        return ctx.SendError(-1, "参数非法");
    }

	db := cfg.TexDb
	if db == nil {
		return ctx.SendError(-1, "连接数据库失败")
	}

	rows, err := db.Query("select prom_port from t_server where node=?", node)
	if err != nil {
		return err
	}
	defer rows.Close()

    ports := make(map[int]bool)
	for rows.Next() {
        tmp := 0
        if err := rows.Scan(&tmp); err != nil {
            return err
        }
        ports[tmp] = true
    }

    // 15001 ~ 16000
    prom_port := 15001
    for prom_port < 16000 {
        if ok,_ := ports[prom_port]; !ok {
            break;
        }
        prom_port += 1;
    }
    return ctx.SendResponse(prom_port)
}
