package model

import (
    "path"
    "strings"
    "encoding/json"
    "github.com/yellia1989/tex-go/tools/util"
    cm "github.com/yellia1989/tex-web/backend/common"
)

var perms *cm.Map
func init() {
    bs, _ := util.LoadFromFile("data/perms.json")
    items := make([]*Permission,0)
    json.Unmarshal(bs, &items)

    items2 := make([]cm.Item,0)
    for _, item := range items {
        items2 = append(items2, item)
    }
    perms = cm.NewMap("data/perms.json", items2)
}

type Permission struct {
    Id  uint32              `json:"id"`
    Name string             `json:"name"`
    Paths []string          `json:"paths"`
}
func (p *Permission) GetId() uint32 {
    return p.Id
}
func (p *Permission) SetId(id uint32) {
    p.Id = id
}
func (p *Permission) copy() *Permission {
    p2 := &Permission{
        Id: p.Id,
        Name: p.Name,
        Paths: make([]string,len(p.Paths)),
    }
    copy(p2.Paths, p.Paths)
    return p2
}
func (perm *Permission) checkPermission(method string, spath string) bool {
    // /index.html特殊处理
    if spath == "/index.html" || spath == "/" {
        return true
    }

    for _, p := range perm.Paths {
        pp := strings.Split(p, " ")
        if pp[0] != "ALL" && pp[0] != method {
            continue
        }

        if matched, _ := path.Match(pp[1], spath); matched {
            return true
        }
    }
    return false
}

func GetPerms() []*Permission {
    if perms == nil {
        return nil
    }

    items := perms.GetItems(func (key, v interface{})bool{
        return true
    })

    if len(items) == 0 {
        return nil
    }

    ps := make([]*Permission,0)
    for _, item := range items {
        // 复制一份防止原始值被修改
        p := item.(*Permission).copy()
        ps = append(ps, p)
    }
    return ps
}

func GetPerm(id uint32) *Permission {
    if perms == nil {
        return nil
    }

    p := perms.GetItem(id)
    if p == nil {
        return nil
    }
    // 复制一份防止原始值被修改
    return p.(*Permission).copy()
}

func AddPerm(name string, paths []string) *Permission {
    if perms == nil {
        return nil
    }

    if name == "" || len(paths) == 0 {
        return nil
    }

    // 对path进行格式化
    for i,_ := range paths {
        tmp := strings.Fields(paths[i])
        if len(tmp) != 2 {
            return nil
        }
        paths[i] = strings.Join(tmp, " ")
    }

    // name不能重复
    items := perms.GetItems(func (key, v interface{})bool{
        p := v.(*Permission)
        return p.Name == name
    })
    if len(items) != 0 {
        return nil
    }

    p := &Permission{
        Name: name,
        Paths: make([]string, len(paths)),
    }
    copy(p.Paths, paths)

    if !perms.AddItem(p) {
        return nil
    }

    // 复制一份防止原始值被修改
    return p.copy()
}

func DelPerm(p *Permission) bool {
    if perms == nil {
        return false
    }

    return perms.DelItem(p)
}

func UpdatePerm(p *Permission) bool {
    if users == nil {
        return false
    }

    p2 := p.copy()
    return perms.UpdateItem(p2)
}
