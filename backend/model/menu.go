package model

import (
    "encoding/json"
    "github.com/yellia1989/tex-go/tools/util"
    cm "github.com/yellia1989/tex-web/backend/common"
)

var menus *cm.Map
func init() {
    bs, _ := util.LoadFromFile("data/menu.json")
    items := make([]*Menu,0)
    json.Unmarshal(bs, &items)

    items2 := make([]cm.Item,0)
    for _, item := range items {
        items2 = append(items2, item)
    }
    menus = cm.NewMap("data/menu.json", items2)
}

type Menu struct {
    Id  uint32              `json:"id"`
    Title   string          `json:"title"`
    Href    string          `json:"href"`
    FontFamily string       `json:"fontFamily"`
    Icon    string          `json:"icon"`
    Spread bool             `json:"spread"`
    IsCheck bool            `json:"isCheck"`
    Children []*Menu        `json:"children"`
}
func (m *Menu) GetId() uint32 {
    return m.Id
}
func (m *Menu) SetId(id uint32) {
    m.Id = id
}
func (m *Menu) addMenu(child *Menu) bool {
    // 同一级目录title不能相同
    if checkDuplicate(m, child) {
        return false
    }
    // 子菜单的id和它加入时子菜单个数有关
    child.Id = uint32(len(m.Children)+1)
    // 默认放在最后
    m.Children = append(m.Children, child)
    return true
}
func (m *Menu) delMenu(id uint32) bool {
    i := -1
    for index, v := range m.Children {
        if v.Id == id {
            i = index
            break
        }
    }
    if i == -1 {
        return false
    }
    m.Children = append(m.Children[:i], m.Children[i+1:]...)
    return true
}
func (m *Menu) getMenu(id uint32) *Menu {
    for _, v := range m.Children {
        if v.Id == id {
            return v
        }
    }
    return nil
}

func GetMenus()[]*Menu {
    if menus == nil {
        return nil
    }

    items := menus.GetItems(func (key, v interface{}) bool {
        return true
    })
    if len(items) == 0 {
        return nil
    }

    ms := make([]*Menu, 0)
    for _, m := range items {
        m2 := *(m.(*Menu))
        ms = append(ms, &m2)
    }
    return ms
}

func checkDuplicate(parent *Menu, m *Menu) bool {
    if parent == nil {
        // 顶级菜单
        items := menus.GetItems(func (key, v interface{}) bool {
            m2 := v.(*Menu)
            return m2.Title == m.Title
        })
        return len(items) != 0
    }
    for _, m2 := range parent.Children {
        if m2.Title == m.Title {
            return true
        }
    }
    return false
}

func AddMenu(m *Menu, pid uint32) bool {
    if menus == nil {
        return false
    }

    // 复制一份
    m2 := *m
    if pid == 0 {
        // 同一级目录title不能相同
        if checkDuplicate(nil, m) {
            return false
        }
        return menus.AddItem(&m2)
    }

    item := menus.GetItem(pid)
    if item == nil {
        // 没有对应的父元素
        return false
    }
    // 复制一份
    pm := *(item.(*Menu))
    if !pm.addMenu(&m2) {
        return false
    }
    return menus.UpdateItem(&pm)
}

func GetMenu(id uint32, pid uint32) *Menu {
    if pid == 0 {
        item := menus.GetItem(id)
        if item == nil {
            return nil
        }
        m := *(item.(*Menu))
        return &m
    }

    item := menus.GetItem(pid)
    if item == nil {
        return nil
    }
    m := *(item.(*Menu))
    return m.getMenu(id)
}

func UpdateMenu(m *Menu) bool {
    if menus == nil {
        return false
    }

    m2 := *m
    return menus.UpdateItem(&m2)
}

func DelMenu(id uint32, pid uint32) bool {
    if menus == nil {
        return false
    }

    if pid == 0 {
        item := menus.GetItem(id)
        if item == nil {
            return false
        }

        return menus.DelItem(item)
    }

    item := menus.GetItem(pid)
    if item == nil {
        return false
    }
    // 复制一份
    parent := *(item.(*Menu))
    if !parent.delMenu(id) {
        return false
    }
    return menus.UpdateItem(&parent)
}

func ClearMenu() bool {
    if menus == nil {
        return false
    }

    return menus.DelAllItem()
}
