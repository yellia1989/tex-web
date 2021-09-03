package stat

import (
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

func ZoneList(c echo.Context) error {
    ctx := c.(*mid.Context)

    zones := getAllZone()
    return ctx.SendResponse(zones)
}
