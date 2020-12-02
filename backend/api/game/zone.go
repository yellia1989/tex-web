package game

import (
    "fmt"
    "database/sql"
    "github.com/yellia1989/tex-web/backend/cfg"
)

func zoneLogDb(zoneid string) (*sql.DB,error) {
    db := cfg.StatDb

    var dbhost sql.NullString
    if err := db.QueryRow("SELECT logdbhost FROM zone WHERE zoneid = " + zoneid).Scan(&dbhost); err != nil {
        return nil,err
    }

    return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/log_zone_%s", cfg.LogDbUser, cfg.LogDbPwd, dbhost.String, zoneid))
}
