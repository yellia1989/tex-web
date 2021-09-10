package stat

import (
    "time"
    "github.com/labstack/echo/v4"
    mid "github.com/yellia1989/tex-web/backend/middleware"
)

func MarkList(c echo.Context) error {
    ctx := c.(*mid.Context)

    datas := make(map[string]interface{})

    mmark := make(map[string]string)

    marks := getMarkDates()
    for _, t := range marks {
        mmark[t.time.Format("2006-01-02")] = t.Desc
    }

    max := getMaxDate()
    now := getDateByTime(time.Now())
    if max.ID > now.ID {
        max = now
    }

    datas["mark"] = mmark
    datas["min"] = getMinDate().time.Format("2006-01-02")
    datas["max"] = max.time.Format("2006-01-02")

    return ctx.SendResponse(datas)
}
