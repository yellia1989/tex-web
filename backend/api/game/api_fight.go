package game

import (
	"encoding/base64"
	"sort"
	"strconv"
	"strings"
	"github.com/labstack/echo"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-web/backend/cfg"
	"github.com/yellia1989/tex-web/backend/common"
	mid "github.com/yellia1989/tex-web/backend/middleware"
)

type fightVerifyErrInfo struct {
	ErrTime         string `json:"err_time"`
	ClientVerify    string `json:"client_verify"`
	ServerVerify    string `json:"server_verify"`
}

type fightVerifyErrInfoBy []fightVerifyErrInfo

func (a fightVerifyErrInfoBy) Len() int      { return len(a) }
func (a fightVerifyErrInfoBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a fightVerifyErrInfoBy) Less(i, j int) bool {
	TmpTimeI := common.ParseTimeInLocal("2006-01-02", a[i].ErrTime)
	TmpTimeJ := common.ParseTimeInLocal("2006-01-02", a[j].ErrTime)

	if !TmpTimeI.Equal(TmpTimeJ) {
		return TmpTimeI.After(TmpTimeJ)
	}

	return a[i].ErrTime > a[j].ErrTime
}

func FightErrInfo(c echo.Context) error {
	ctx := c.(*mid.Context)
	startTime := ctx.QueryParam("startTime")
	endTime := ctx.QueryParam("endTime")
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	if startTime == "" || endTime == "" {
		return ctx.SendError(-1, "参数非法")
	}

	db := cfg.LogDb

	sql := "SELECT timeymd, client_verify, server_verify FROM fight_verify "
	sql += "WHERE time BETWEEN '" + startTime + "' AND '" + endTime + "'"

	log.Infof("sql: %s", sql)

	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	slFightVerifyInfo := make([]fightVerifyErrInfo, 0)
	for rows.Next() {
		var r fightVerifyErrInfo
		if err := rows.Scan(&r.ErrTime, &r.ClientVerify, &r.ServerVerify); err != nil {
			return err
		}

		decodeBytes, _ := base64.StdEncoding.DecodeString(r.ClientVerify)
		r.ClientVerify = string(decodeBytes)
		r.ClientVerify = strings.Replace(r.ClientVerify, "\n", "<br>", -1)

		decodeBytes1, _ := base64.StdEncoding.DecodeString(r.ServerVerify)
		r.ServerVerify = string(decodeBytes1)
		r.ServerVerify = strings.Replace(r.ServerVerify, "\n", "<br>", -1)

		slFightVerifyInfo = append(slFightVerifyInfo, r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	sort.Sort(fightVerifyErrInfoBy(slFightVerifyInfo))

	vSimpleErrInfo := common.GetPage(slFightVerifyInfo, page, limit)

	return ctx.SendArray(vSimpleErrInfo, len(slFightVerifyInfo))
}
