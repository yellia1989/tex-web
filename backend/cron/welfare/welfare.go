package welfare

import (
    "fmt"
    "time"
    "strings"
    "bytes"
    "sort"
    "strconv"
    "math/rand"
    "context"
    dsql "database/sql"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	"github.com/yellia1989/tex-web/backend/api/gm"
	"github.com/yellia1989/tex-go/tools/log"
)

type wfRole struct {
    id int
    zoneid int
    roleid int
    mapid int
    t time.Time
    cmd string
}

type wfTask struct {
    conn *dsql.Conn
    id int
    sroles string
    scmds string
    cmd_time string
    scur_time string
    roles map[int]*wfRole
    status int
    begin_time time.Time
    end_time time.Time
    step int
    cur_time time.Time
    slg int
}

type wfRoleSorter []*wfRole
// Len is part of sort.Interface.
func (s wfRoleSorter) Len() int {
    return len(s)
}
// Swap is part of sort.Interface.
func (s wfRoleSorter) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
// Less is part of sort.Interface.
func (s wfRoleSorter) Less(i, j int) bool {
    return s[i].t.Before(s[j].t)
}

func (task *wfTask) run(now time.Time) {
    // 生成今日福利
    if task.status == 1 && len(task.roles) == 0 && task.scur_time != now.Format("2006-01-02") {
        if now.After(task.end_time) {
            // 已经过了福利期截止日期
            return
        }

        if now.Before(task.cur_time.Add(time.Duration(task.step*24) * time.Hour)) {
            // 距离上次发奖日期没到
            return
        }

        tx, err := task.conn.BeginTx(ctx, nil)
	    if err != nil {
            panic(fmt.Sprintf("cron [welfare] 开始事务失败, err: %s, taskid: %d", err.Error(), task.id))
	    }
	    defer tx.Rollback()

        var buff bytes.Buffer
        if !task.generate(now, &buff) {
            return
        }

        _, err = tx.Exec(buff.String())
	    if err != nil {
            log.Errorf("cron [welfare] task exec, err: %s, sql: %s, taskid: %d", err.Error(), buff.String(), task.id)
            return
	    }
        _, err = tx.Exec("UPDATE welfare_task SET cur_time='"+now.Format("2006-01-02")+"' WHERE id="+strconv.Itoa(task.id))
	    if err != nil {
            log.Errorf("cron [welfare] task exec, err: %s, taskid: %d", err.Error(), task.id)
            return
	    }

        if err := tx.Commit(); err != nil {
            log.Errorf("cron [welfare] commit err: %s, taskid: %d", err.Error(), task.id)
            return
        }

        task.scur_time = now.Format("2006-01-02")
        task.cur_time = now
        log.Infof("cron [welfare] generate complete: %s, taskid: %d", now.Format("2006-01-02"), task.id)
        return
    }

    // 读取今日福利
    if len(task.roles) == 0 {
        sql := "SELECT id,zoneid,roleid,mapid,time,cmd FROM welfare_roles WHERE status = 0 and taskid_pk = ? order by time asc limit 100"
	    rows, err := task.conn.QueryContext(ctx, sql, task.id)
	    if err != nil {
            log.Errorf("cron [welfare] role query, err: %s, taskid: %d", err.Error(), task.id)
            return
	    }
	    defer rows.Close()

        var t string
        for rows.Next() {
            var r wfRole
            if err := rows.Scan(&r.id, &r.zoneid, &r.roleid, &r.mapid, &t, &r.cmd); err != nil {
                log.Errorf("cron [welfare] role scan, err: %s, taskid: %d", err.Error(), task.id)
            } else {
                r.t = common.ParseTimeInLocal("2006-01-02 15:04:05", t)
                task.roles[r.id] = &r
            }
        }
    } else {
        // 调用gm发奖
        roles := make([]*wfRole, 0)
        for _, r := range task.roles {
            if r.t.After(now) {
                continue
            }
            roles = append(roles, r)
        }
        if len(roles) != 0 {
            var result string
            for _, r := range roles {
                if err := gm.Cmd("welfare", strconv.Itoa(r.zoneid), strconv.Itoa(r.mapid), r.cmd, &result); err != nil {
                    log.Errorf("cron [welfare] gm failed, err: %s, result: %s", err.Error(), result)
                }

                delete(task.roles, r.id)

 	            _, err := task.conn.ExecContext(ctx, "UPDATE welfare_roles SET status=1,exec_time='"+now.Format("2006-01-02 15:04:05")+"',exec_result='"+result+"' WHERE id="+strconv.Itoa(r.id))
	            if err != nil {
                    log.Errorf("cron [welfare] update role status, err: %s, id: %d", err.Error(), r.id)
	            }
            }
        }
    }
}

func (task *wfTask) generate(now time.Time, buff *bytes.Buffer) bool {
    vcmdtime := strings.Split(task.cmd_time, "-")
    if len(vcmdtime) != 2 {
        log.Errorf("cron [welfare] generate invalid cmd_time format: %s, taskid: %d", task.cmd_time, task.id)
        return false
    }

    cmd_begin_time := common.ParseTimeInLocal("2006-01-02 15:04:05", now.Format("2006-01-02") + " " + vcmdtime[0])
    cmd_end_time := common.ParseTimeInLocal("2006-01-02 15:04:05", now.Format("2006-01-02") + " " + vcmdtime[1])
    if cmd_begin_time.After(cmd_end_time) {
        cmd_begin_time, cmd_end_time = cmd_end_time, cmd_begin_time
    }

    d := int64(cmd_end_time.Sub(cmd_begin_time).Seconds())

    vcmds := strings.Split(strings.Replace(task.scmds, "\r\n", "\n", -1), "\n")
    if len(vcmds) == 0 {
        log.Errorf("cron [welfare] generate invalid cmds format: %s, taskid: %d", task.scmds, task.id)
        return false
    }

    vroles := strings.Split(strings.Replace(task.sroles, "\r\n", "\n", -1), "\n")
    if len(vroles) == 0 {
        log.Errorf("cron [welfare] generate invalid roles format: %s, taskid: %d", task.sroles, task.id)
        return false
    }

    var err error
    role2cmd := make([]*wfRole, 0)
    for _, r := range vroles {
        vr := strings.Split(r, ",")
        if task.slg == 0 && len(vr) != 2 {
            log.Errorf("cron [welfare] generate invalid role format: %s, taskid: %d", r, task.id)
            continue
        }
        if task.slg == 1 && len(vr) != 3 {
            log.Errorf("cron [welfare] generate invalid role format: %s, taskid: %d", r, task.id)
            continue
        }
        var zoneid int
        var roleid int
        var mapid int
        if zoneid, err = strconv.Atoi(vr[0]); err != nil {
            log.Errorf("cron [welfare] generate invalid role format: %s, taskid: %d", r, task.id)
            continue
        }
        if roleid, err = strconv.Atoi(vr[1]); err != nil {
            log.Errorf("cron [welfare] generate invalid role format: %s, taskid: %d", r, task.id)
            continue
        }
        if task.slg == 1 {
            if mapid, err = strconv.Atoi(vr[2]); err != nil {
                log.Errorf("cron [welfare] generate invalid role format: %s, taskid: %d", r, task.id)
                continue
            }
        }
        for _, cmd := range vcmds {
            if len(cmd)== 0 {
                continue
            }
            var r wfRole
            r.zoneid = zoneid
            r.roleid = roleid
            r.mapid = mapid
            r.t = cmd_begin_time.Add(time.Duration(rand.Int63n(d))*time.Second)
            r.cmd = strings.ReplaceAll(strings.TrimSpace(cmd), "\t", " ")
            r.cmd = strings.ReplaceAll(r.cmd, "{roleid}", vr[1])
            r.cmd = strings.ReplaceAll(r.cmd, "{zoneid}", vr[0])
            role2cmd = append(role2cmd, &r)
        }
    }

    if len(role2cmd) == 0 {
        return false
    }
    buff.WriteString("INSERT INTO welfare_roles(zoneid,roleid,mapid,time,cmd,status,taskid_pk) VALUES")
    sort.Sort(wfRoleSorter(role2cmd))
    for i, r := range role2cmd {
        if i != 0 {
            buff.WriteString(",")
        }
        buff.WriteString("("+strconv.Itoa(r.zoneid)+","+strconv.Itoa(r.roleid)+","+strconv.Itoa(r.mapid)+",'"+r.t.Format("2006-01-02 15:04:05")+"','"+r.cmd+"',0,"+strconv.Itoa(task.id)+")")
    }
    return true
}

var ctx context.Context
var conn *dsql.Conn
var tasks map[int]*wfTask

func init() {
    ctx = context.Background()
    tasks = make(map[int]*wfTask,0)
}

func checkConn() {
    var err error
    if conn != nil {
        err = conn.PingContext(ctx)
        if err != nil {
            conn.Close()
            conn = nil
        } else {
            return
        }
    }

    if conn == nil {
        conn, err = cfg.StatDb.Conn(ctx)
        if err != nil {
            panic(fmt.Sprintf("cron [welfare] create conn err: %s", err.Error()))
        }
    }
}

func Cron(now time.Time) {
    defer func() {
        if err := recover(); err != nil {
            log.Errorf("%v", err)
            if len(tasks) != 0 {
                tasks = make(map[int]*wfTask,0)
            }
        }
    }()

    checkConn()

    sql := "SELECT id,roles,cmds,cmd_time,cur_time,status,begin_time,end_time,step,slg FROM welfare_task WHERE ? between begin_time and end_time"
	rows, err := conn.QueryContext(ctx, sql, now.Format("2006-01-02"))
	if err != nil {
        log.Errorf("cron [welfare] query err: %s", err.Error())
        return
	}
	defer rows.Close()

    for rows.Next() {
        var task wfTask
        task.conn = conn
        task.roles = make(map[int]*wfRole,0)
        var cur_time dsql.NullString
        var beginTime string
        var endTime string
        if err := rows.Scan(&task.id, &task.sroles, &task.scmds, &task.cmd_time, &cur_time, &task.status, &beginTime, &endTime, &task.step, &task.slg); err != nil {
            log.Errorf("cron [welfare] scan err: %s, taskid: %d", err.Error(), task.id)
        } else {
            task.scur_time = cur_time.String
            if task.scur_time != "" {
                task.cur_time = common.ParseTimeInLocal("2006-01-02", task.scur_time)
            }
            task.begin_time = common.ParseTimeInLocal("2006-01-02", beginTime)
            task.end_time = common.ParseTimeInLocal("2006-01-02", endTime)

            if oldtask,ok := tasks[task.id]; !ok {
                tasks[task.id] = &task
            } else {
                oldtask.sroles = task.sroles
                oldtask.scmds = task.scmds
                oldtask.cmd_time = task.cmd_time
                oldtask.status = task.status
                oldtask.begin_time = task.begin_time
                oldtask.end_time = task.end_time
                oldtask.step = task.step
                oldtask.slg = task.slg
            }
        }
    }

    for _, task := range tasks {
        task.run(now)
    }
}
