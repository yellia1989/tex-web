package model

import (
    "testing"
)

func TestAddMenu(t *testing.T) {
    m := &Menu{Title: "index"}
    if AddMenu(m, 0) == false {
        t.Fatal("AddMenu falied, should success")
    }
    if AddMenu(m, 0) == true {
        t.Fatal("AddMenu success, should failed")
    }
    m.Title = "gm"
    if AddMenu(m, 0) == false {
        t.Fatal("AddMenu falied, should success")
    }
    m2 := &Menu{Title: "adddir", Href: "dir-add.html"}
    if AddMenu(m2, 2) == false {
        t.Fatal("AddMenu falied, should success")
    }
    if DelMenu(1, 2) == false {
        t.Fatal("DelMenu falied, should success")
    }
    if DelMenu(2, 0) == false {
        t.Fatal("DelMenu falied, should success")
    }

    m3 := GetMenu(1, 0)
    if m3 == nil {
        t.Fatal("GetMenu failed, should success")
    }
    m3.Href = "index.html"
    if UpdateMenu(m3) == false {
        t.Fatal("UpdateMenu failed, should success")
    }
    m4 := GetMenu(1, 0)
    if m4 == nil || m4.Href != "index.html" {
        t.Fatal("UpdateMenu failed")
    }

    if DelMenu(1, 0) == false {
        t.Fatal("DelMenu falied, should success")
    }

    if ClearMenu() == false {
        t.Fatal("ClearMenu failed, should success")
    }
}
