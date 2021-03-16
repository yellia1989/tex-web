package game

import (
//    "fmt"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
//    "github.com/yellia1989/tex-go/tools/log"
)

type errInfo struct {
    ZoneId      uint32 `json:"zond_id"`
    ErrMessage  string `json:"err_info"`
    ErrTime     string `json:"err_time"`
}

func ErrInfo(c echo.Context) error {
    ctx := c.(*mid.Context)
    logs := make([]errInfo, 0)
    return ctx.SendArray(logs, len(logs))
}
