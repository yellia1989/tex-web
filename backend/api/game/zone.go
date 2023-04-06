package game

import (
    "fmt"
    "database/sql"
    "github.com/yellia1989/tex-web/backend/cfg"
    "github.com/yellia1989/tex-web/backend/api/gm"
    "github.com/yellia1989/tex-web/backend/common"
)

func zoneLogDb(zoneid string) (*sql.DB,error) {
    db := cfg.StatDb

    zoneid2 := common.U32toa(gm.GetZoneId(common.Atou32(zoneid)))

    var dbhost sql.NullString
    if err := db.QueryRow("SELECT logdbhost FROM zone WHERE zoneid = " + zoneid2).Scan(&dbhost); err != nil {
        return nil,err
    }

    return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/log_zone_%s", cfg.LogDbUser, cfg.LogDbPwd, dbhost.String, zoneid2))
}
