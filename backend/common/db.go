package common

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var logdb *sql.DB
var db *sql.DB
var stat *sql.DB
func init() {
    logdb, _ = sql.Open("mysql", "dev:777777@tcp(192.168.0.16:3306)/log_zone_1")
    db, _ = sql.Open("mysql", "dev:777777@tcp(192.168.0.16:3306)/db_zone_1")
    stat, _ = sql.Open("mysql", "dev:777777@tcp(192.168.0.16:3306)/db_stat")
}

// 日志数据库
func GetLogDb() *sql.DB {
    return logdb
}

// 游戏数据库
func GetDb() *sql.DB {
    return db
}

// 统计数据库
func GetStatDb() *sql.DB {
    return stat 
}