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
	zoneid2 := zoneid

	if options.trunMerge {
		zoneid2 = common.U32toa(gm.GetZoneId(common.Atou32(zoneid)))
	}

	return sql.Open("mysql", fmt.Sprintf("%s/%slog_zone_%s", cfg.LogDbInfo, cfg.GameDbPrefix, zoneid2))
}
