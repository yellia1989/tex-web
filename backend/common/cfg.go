package common

import (
    "fmt"
    "net/url"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/yellia1989/tex-go/tools/util"
    tex "github.com/yellia1989/tex-go/service"
)

// 日志数据库
var logdb *sql.DB

// 游戏数据库
var gamedb *sql.DB

// 统计数据库
var statdb *sql.DB

// 配置文件
var Cfg *util.Config

// locator
var comm *tex.Communicator

// app
var app string

// logo
var logo string

func init() {
    Cfg = util.NewConfig()
}

func ParseCfg(file string) (err error) {
    defer func() {
        if err == nil {
            if err := recover(); err != nil {
                err = fmt.Errorf("load cfg file err:%s", err)
            }
        }
    }()

    Cfg.ParseFile(file)

    logdb, err = sql.Open("mysql", Cfg.GetCfg("logdb", ""))
    if err != nil {
        return
    }

    gamedb, err = sql.Open("mysql", Cfg.GetCfg("gamedb", ""))
    if err != nil {
        return
    }

    statdb, err = sql.Open("mysql", Cfg.GetCfg("statdb", ""))
    if err != nil {
        return
    }

    locator := Cfg.GetCfg("locator", "")
    if locator == "" {
        err = fmt.Errorf("invalid locator str\n")
    }
    comm = tex.NewCommunicator(locator)

    app = Cfg.GetCfg("app", "")
    if app == "" {
        err = fmt.Errorf("invalid app str\n")
    }

    logo = url.QueryEscape(Cfg.GetCfg("logo", ""))
    if logo == "" {
        err = fmt.Errorf("invalid logo str\n")
    }

    return
}

func GetLogDb() *sql.DB {
    return logdb
}

func GetDb() *sql.DB {
    return gamedb
}

func GetStatDb() *sql.DB {
    return statdb 
}

func GetLocator() *tex.Communicator {
    return comm
}

func GetApp() string {
    return app
}

func GetLogo() string {
    return logo
}
