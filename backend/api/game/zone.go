package game

import (
	"database/sql"
	"fmt"
	"github.com/yellia1989/tex-web/backend/api/gm"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
)

func zoneLogDb(zoneid string) (*sql.DB, error) {
	return zoneLogDbWithOptions(zoneid, ZoneLogOptions{trunMerge: true})
}

type ZoneLogOptions struct {
	trunMerge bool
}

func zoneLogDbWithOptions(zoneid string, options ZoneLogOptions) (*sql.DB, error) {
	db := cfg.StatDb

	zoneid2 := zoneid

	if options.trunMerge {
		zoneid2 = common.U32toa(gm.GetZoneId(common.Atou32(zoneid)))
	}

	var dbhost sql.NullString
	if err := db.QueryRow("SELECT logdbhost FROM zone WHERE zoneid = " + zoneid2).Scan(&dbhost); err != nil {
		return nil, err
	}

	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/log_zone_%s", cfg.LogDbUser, cfg.LogDbPwd, dbhost.String, zoneid2))
}
