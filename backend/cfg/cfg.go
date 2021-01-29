package cfg

import (
    "fmt"
    "time"
    "strings"
    "net/url"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/yellia1989/tex-go/tools/util"
    tex "github.com/yellia1989/tex-go/service"
    "github.com/yellia1989/tex-web/backend/common"
)

// 配置
var Config *util.Config

// 是否开启调试模式
var Debug bool

// 是否开启框架日志
var FrameworkDebug bool

// 监听端口号
var Listen string

// locator
var Comm *tex.Communicator

// app
var App string

// logo
var Logo string

// 日志数据库
var LogDb *sql.DB

// 日志数据库连接用户名
var LogDbUser string

// 日志数据库连接密码
var LogDbPwd string

// 游戏全局数据库
var GameGlobalDb *sql.DB

// 统计数据库
var StatDb *sql.DB

// gamedb prefix
var GameDbPrefix string

// 最小日期
var LogDateMin time.Time

// 日志同步间隔
var LogSyncInterval time.Duration

// 日志统计间隔
var LogStatInterval time.Duration

// 大R充值金额分
var StatRmoney uint32

// 只统计相关渠道信息
var StatChannels []string

func ParseCfg(file string) (err error) {
    if Config == nil {
        Config = util.NewConfig()
    }
    Config.ParseFile(file)
    cfg := Config

    Debug = cfg.GetBool("debug", false)
    FrameworkDebug = cfg.GetBool("framework-debug", false)
    Listen = cfg.GetCfg("listen", ":8008")

    locator := cfg.GetCfg("locator", "")
    if locator == "" {
        panic("locator required")
    }
    Comm = tex.NewCommunicator(locator)

    App = cfg.GetCfg("app", "")
    if App == "" {
        panic("app required")
    }

    Logo = url.QueryEscape(cfg.GetCfg("logo", ""))
    if Logo == "" {
        panic("logo required")
    }

    logdb := cfg.GetCfg("logdb", "")
    if len(logdb) == 0 {
        panic("logdb required")
    }
    vtmp := strings.SplitN(logdb, "@", 2)
    vtmp2 := strings.Split(vtmp[0], ":")
    if len(vtmp2) != 2 {
        panic("invalid logdb format")
    }
    LogDbUser = vtmp2[0]
    LogDbPwd = vtmp2[1]
    LogDb, err = sql.Open("mysql", logdb)
    if err != nil {
        panic(fmt.Sprintf("create log db err: %s", err.Error()))
    }

    gameglobaldb := cfg.GetCfg("gameglobaldb", "")
    if len(gameglobaldb) == 0 {
        panic("gameglobaldb required")
    }
    GameGlobalDb, err = sql.Open("mysql", gameglobaldb)
    if err != nil {
        panic(fmt.Sprintf("create game global db err: %s", err.Error()))
    }

    statdb := cfg.GetCfg("statdb", "")
    if len(statdb) == 0 {
        panic("statdb required")
    }
    StatDb, err = sql.Open("mysql", statdb)
    if err != nil {
        panic(fmt.Sprintf("create stat db err: %s", err.Error()))
    }

    GameDbPrefix = cfg.GetCfg("gamedb-prefix", "")

    clog := cfg.GetSubCfg("log")
    if clog == nil {
        panic("<log> conf required")
    }
    LogDateMin = common.ParseTimeInLocal("2006-01-02", clog.GetCfg("dateMin", ""))
    LogSyncInterval = clog.GetDuration("syncInterval", "5m")
    LogStatInterval = clog.GetDuration("statInterval", "5m")

    cstat := cfg.GetSubCfg("stat")
    if cstat == nil {
        panic("<stat> conf required")
    }
    StatRmoney = uint32(cstat.GetInt("rmoney", 500000))

    tmp := cstat.GetCfg("channel","")
    StatChannels = strings.Split(tmp, ",")

    return
}
