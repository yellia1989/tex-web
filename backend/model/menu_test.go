package model

import (
    "testing"
    "encoding/json"
)

func TestAddMenu(t *testing.T) {
    m := &Menu{Title: "index"}
    if AddMenu(m, 0) == false || m.Id == 0 {
        t.Fatal("AddMenu falied, should success")
    }
    m2 := m.copy()
    // id != 0的不能添加
    if AddMenu(m2, 0) == true {
        t.Fatal("AddMenu success, should failed")
    }
    // 同级title相同的不能添加
    m2.Id = 0
    if AddMenu(m2, 0) == true {
        t.Fatal("AddMenu success, should failed")
    }
    m2.Title = "gm"
    if AddMenu(m2, 0) == false {
        t.Fatal("AddMenu falied, should success")
    }
    // 添加子菜单
    m3 := &Menu{Title: "adddir", Href: "dir-add.html"}
    if AddMenu(m3, 2) == false || m3.Id != 201 {
        t.Fatal("AddMenu falied, should success")
    }
    m4 := m3.copy()
    m4.Title = "分区管理"
    if UpdateMenu(m4) == false {
        t.Fatal("UpdateMenu failed, should success")
    }

    ms := GetMenus()
    str, _ := json.MarshalIndent(ms, "", "  ")
    t.Logf("%s", string(str))

    if DelMenu(201) == false {
        t.Fatal("DelMenu falied, should success")
    }
    if DelMenu(2) == false {
        t.Fatal("DelMenu falied, should success")
    }

    m5 := GetMenu(1)
    if m5 == nil {
        t.Fatal("GetMenu failed, should success")
    }
    m5.Href = "index.html"
    if UpdateMenu(m5) == false {
        t.Fatal("UpdateMenu failed, should success")
    }
    m6 := GetMenu(1)
    if m6 == nil || m6.Href != "index.html" {
        t.Fatal("UpdateMenu failed")
    }

    if DelMenu(1) == false {
        t.Fatal("DelMenu falied, should success")
    }

    if ClearMenu() == false {
        t.Fatal("ClearMenu failed, should success")
    }
}
