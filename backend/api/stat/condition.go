package stat

import (
    "strings"
    "github.com/yellia1989/tex-web/backend/cfg"
)

type condition struct {
    sql string
}

func (d *condition) String() string {
    return d.sql
}

func newCondition(sql string) *condition {
    return &condition{sql: sql}
}

// 只统计and_gm,ios_gm渠道的账号
var accountCond *condition
var roleCond *condition

func InitCondition() {
    if len(cfg.StatChannels) == 0 {
        return
    }

    vtmp := make([]string,0)
    for _, v := range cfg.StatChannels {
        vtmp = append(vtmp, "'"+v+"'")
    }
    channel := strings.Join(vtmp, ",")
    accountCond = newCondition(" channel IN ("+channel+") ")
    roleCond = newCondition(" EXISTS (Select 0 FROM account WHERE accountid_fk = account.id AND channel IN ("+channel+"))")
}
