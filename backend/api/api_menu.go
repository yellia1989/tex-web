package api

import (
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func MenuList(c echo.Context) error {
    ctx := c.(*mid.Context)
    ms := model.GetMenus()
    return ctx.SendResponse(ms)
}
