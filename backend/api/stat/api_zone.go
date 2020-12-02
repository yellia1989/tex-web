package stat

import (
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

func ZoneList(c echo.Context) error {
    ctx := c.(*mid.Context)

    zones := getAllZone()
    return ctx.SendResponse(zones)
}
