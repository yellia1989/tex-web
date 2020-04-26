package common

import (
    "reflect"
)

func GetPage(data interface{}, page int, limit int) interface{} {
    v := reflect.ValueOf(data)
    l := v.Len()

    if l < limit {
        return data
    }

    from := (page-1) * limit
    if from >= l {
        from = 0
    }
    to := from + limit
    if to > l {
        to = l
    }

    return v.Slice(from, to).Interface()
}
