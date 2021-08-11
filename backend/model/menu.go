package model

import (
    "database/sql"
    "encoding/json"
    "github.com/yellia1989/tex-go/tools/util"
    "github.com/yellia1989/tex-go/tools/log"
    "github.com/yellia1989/tex-web/backend/cfg"
    cm "github.com/yellia1989/tex-web/backend/common"
    "strconv"
    "strings"
)

var menus *cm.Map

func InitMenu() {
    bs, err := util.LoadFromFile("data/menu.json")
    if err != nil {
        log.Error("menu init failed, %s", err.Error())
    }

    db := cfg.StatDb
    if db == nil {
        panic("menu init failed, statedb error")
    }

    items := make([]*Menu,0)
    err = json.Unmarshal(bs, &items)
    if err != nil {
        log.Error("menu init failed, %s", err.Error())
    }

    for _,v := range items {
        v.FillRole(db)
    }

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
    Checked bool            `json:"checked"`
    Children []*Menu         `json:"children"`
    Role []uint32           `json:"role"`
}

func (m *Menu) FillRole(db *sql.DB) {
    roles :=""
    db.QueryRow("select role_ids from sys_menu_role where menu_id = ?",m.Id).Scan(&roles)
    roleIds := make([]uint32,0)
    if roles != "" {
        ids := strings.Split(roles,",")
        for _,v := range ids {
            id,_ := strconv.ParseUint(v,10,32)
            roleIds = append(roleIds, uint32(id))
        }
    }
    m.Role = roleIds
    for _,v := range m.Children{
        v.FillRole(db)
    }
}

func (m *Menu) UpdateRole(db *sql.DB) error {
    var roleString string
    roleStrs := make([]string,0)
    for _,v := range m.Role {
        roleStrs = append(roleStrs, strconv.FormatUint(uint64(v),32))
    }
    roleString = strings.Join(roleStrs,",")
    _,err := db.Exec("insert into sys_menu_role(menu_id,role_ids) VALUES (?,?) on duplicate key update role_ids= ?",m.Id,roleString,roleString)
    if err!=nil{
        return err
    }
    for _,v := range m.Children{
        err = v.UpdateRole(db)
        if err!=nil {
            return err
        }
    }
    return nil
}

func (m *Menu) GetId() uint32 {
    return m.Id
}
func (m *Menu) SetId(id uint32) {
    m.Id = id
}
func (m *Menu) copy() *Menu {
    // 深拷贝
    n := &Menu{
        Id: m.Id,
        Title: m.Title,
        Href: m.Href,
        FontFamily: m.FontFamily,
        Icon: m.Icon,
        Spread: m.Spread,
        Checked: m.Checked,
    }
    if len(m.Children) != 0 {
        n.Children = make([]*Menu, len(m.Children))
        for i, v := range m.Children {
            n.Children[i] = v.copy()
        }
    }
    if len(m.Role) != 0 {
        n.Role = make([]uint32,len(m.Role))
        copy(n.Role, m.Role)
    }
    return n
}
func (m *Menu) addMenu(child *Menu) bool {
    // 同一级目录title不能相同
    if checkDuplicate(m, child) {
        return false
    }
    // id依次递增
    maxid := m.Id*100
    for _, c := range m.Children {
        if c.Id > maxid {
            maxid = c.Id
        }
    }
    child.Id = maxid + 1
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
    oid := id
    // 先找到孩子节点
    for ; id > (m.Id*100+100); id /= 100 {
    }

    var child *Menu
    for _, v := range m.Children {
        if v.Id == id {
            child = v
            break
        }
    }
    if child == nil {
        return nil
    }
    if oid == id {
        return child
    }
    return child.getMenu(oid)
}
func (m *Menu) updateMenu(update *Menu) {
    if m.Id != update.Id {
        panic("not a same menu")
    }
    m.Title = update.Title
    m.Href = update.Href
    m.FontFamily = update.FontFamily
    m.Icon = update.Icon
    m.Spread = update.Spread
    m.Checked = update.Checked
    m.Role = nil
    if len(update.Role) != 0 {
        m.Role = make([]uint32,len(update.Role))
        copy(m.Role,update.Role)
    }
    m.Children = nil
    for _, v := range update.Children {
        m.Children = append(m.Children, v.copy())
    }
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
        ms = append(ms, m.(*Menu).copy())
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
    if menus == nil || m.Id != 0 {
        return false
    }

    if pid == 0 {
        // 同一级目录title不能相同
        if checkDuplicate(nil, m) {
            return false
        }
        return menus.AddItem(m)
    }

    // 先找到顶级菜单
    top := getTopMenu(pid)
    if top == nil {
        return false
    }
    topp := top.copy()
    // 找到父节点
    parent := topp
    if pid > 100 {
        parent = topp.getMenu(pid)
    }

    if parent == nil || !parent.addMenu(m) {
        return false
    }

    return menus.UpdateItem(topp)
}

func getTopMenu(id uint32) *Menu {
    for ; id > 100; id /= 100 {
    }
    item := menus.GetItem(id)
    if item == nil {
        return nil
    }
    return item.(*Menu)
}

func GetTopMenu(id uint32) *Menu {
    top := getTopMenu(id)
    if top == nil {
        return nil
    }
    return top.copy()
}

func GetMenu(id uint32) *Menu {
    // 先找到顶级菜单
    top := getTopMenu(id)
    if top == nil {
        return nil
    }

    if id <= 100 {
        return top.copy()
    }

    m := top.getMenu(id)
    if m == nil {
        return nil
    }
    return m.copy()
}

func UpdateMenu(m *Menu) bool {
    if menus == nil {
        return false
    }

    db := cfg.StatDb
    if db == nil {
        return false
    }

    // 先找到顶级菜单
    top := getTopMenu(m.Id)
    if top == nil {
        return false
    }

    if m.Id <= 100 {
        topp := top.copy()
        topp.updateMenu(m)
        err := topp.UpdateRole(db)
        if err!=nil{
            return false
        }
        return menus.UpdateItem(topp)
    }

    topp := top.copy()
    old := topp.getMenu(m.Id)
    if old == nil {
        return false
    }
    old.updateMenu(m)
    err := old.UpdateRole(db)
    if err!=nil{
        return false
    }
    return menus.UpdateItem(topp)
}

func DelMenu(id uint32) bool {
    if menus == nil {
        return false
    }

    // 获取顶级菜单
    top := getTopMenu(id)
    if top == nil {
        return false
    }

    if id <= 100 {
        return menus.DelItem(top)
    }

    topp := top.copy()
    parent := topp
    if id/100 > 100 {
        pid := id/100*100
        parent = topp.getMenu(pid)
    }
    if !parent.delMenu(id) {
        return false
    }
    return menus.UpdateItem(topp)
}

func ClearMenu() bool {
    if menus == nil {
        return false
    }

    return menus.DelAllItem()
}
