package gm

import (
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

type Item struct {
    ID uint32 `json:"value"`
    Name string `json:"name"`
};

func ItemList(c echo.Context) error {
    ctx := c.(*mid.Context)

    items := make([]Item,0)
    items = append(items,Item{ID: 301, Name: "改名卡"})
    items = append(items,Item{ID: 2100, Name: "天命碎片"})
    items = append(items,Item{ID: 2101, Name: "曹操碎片"})
    items = append(items,Item{ID: 2102, Name: "甑姬碎片"})
    items = append(items,Item{ID: 2103, Name: "许褚碎片"})
    items = append(items,Item{ID: 2104, Name: "夏侯渊碎片"})

    return ctx.SendResponse(items)
}
