package common

import (
    "os"
    "time"
    "strings"
    "strconv"
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

func Atou32(s string) uint32 {
    ui, _ := strconv.ParseUint(s, 10, 32)
    return uint32(ui)
}

func U32toa(u32 uint32) string {
    return strconv.FormatUint(uint64(u32), 10)
}

func Atoi(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func Atou32v(s string, sep string) []uint32 {
    if len(s) == 0 {
        return nil
    }

    vv := strings.Split(s, sep)
    vu32 := make([]uint32, 0)

    for _, v := range vv {
        vu32 = append(vu32, Atou32(v))
    }
    return vu32
}

func U32vtoa(vv []uint32, sep string) string {
    vs := make([]string,0)
    for _, v := range vv {
        vs = append(vs, U32toa(v))
    }
    return strings.Join(vs, sep)
}

func ParseTimeInLocal(layout string, v string) time.Time {
    t,_ := time.ParseInLocation(layout, v, time.Local)
    return t
}

func FormatTimeInLocal(layout string, t time.Time) string {
    return t.Local().Format(layout)
}

func MaxInt(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func MaxInt64(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
