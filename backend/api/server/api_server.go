package server

import (
	"github.com/labstack/echo/v4"
	"github.com/yellia1989/tex-web/backend/api/gm/rpc"
)

func ServerList(ctx echo.Context) error {
    _ = rpc.PatchTaskItemReq{}
    return nil
}

func ServerStart(ctx echo.Context) error {
    return nil
}

func ServerStop(ctx echo.Context) error {
    return nil
}

func ServerRestart(ctx echo.Context) error {
    return nil
}
