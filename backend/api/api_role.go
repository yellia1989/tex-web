package api

import (
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/model"
)

func RoleList(c echo.Context) error {
    ctx := c.(*mid.Context)
    rs := model.GetRoles()
    return ctx.SendResponse(rs)
}