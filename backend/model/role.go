package model

import (
    "encoding/json"
    "github.com/yellia1989/tex-go/tools/util"
    cm "github.com/yellia1989/tex-web/backend/common"
)

var roles *cm.Map
func init() {
    bs, _ := util.LoadFromFile("data/roles.json")
    items := make([]*Role,0)
    json.Unmarshal(bs, &items)

    items2 := make([]cm.Item,0)
    for _, item := range items {
        items2 = append(items2, item)
    }
    roles = cm.NewMap("data/roles.json", items2)
}

type Role struct {
    Id  uint32              `json:"value"`
    Name string             `json:"name"`
    Perms []uint32          `json:"perms"`
}
func (r *Role) GetId() uint32 {
    return r.Id
}
func (r *Role) SetId(id uint32) {
    r.Id = id
}
func (r *Role) copy() *Role {
    r2 := &Role{
        Id: r.Id,
        Name: r.Name,
        Perms: make([]uint32,len(r.Perms)),
    }
    copy(r2.Perms, r.Perms)
    return r2
}

func GetRoles() []*Role {
    if roles == nil {
        return nil
    }

    items := roles.GetItems(func (key, v interface{})bool{
        return true
    })

    if len(items) == 0 {
        return nil
    }

    rs := make([]*Role,0)
    for _, item := range items {
        // 复制一份防止原始值被修改
        r := item.(*Role).copy()
        rs = append(rs, r)
    }
    return rs
}

func GetRole(id uint32) *Role {
    if roles == nil {
        return nil
    }

    item := roles.GetItem(id)
    if item == nil {
        return nil
    }

    return item.(*Role).copy()
}

func AddRole(name string, perms []uint32) *Role {
    if roles == nil {
        return nil
    }

    items := roles.GetItems(func (key, v interface{})bool{
        r := v.(*Role)
        return r.Name == name
    })

    if len(items) != 0 {
        return nil
    }

    r := &Role{
        Name: name,
        Perms: make([]uint32,len(perms)),
    }
    copy(r.Perms,perms)
    if !roles.AddItem(r) {
        return nil
    }

    return r.copy()
}

func DelRole(r *Role) bool {
    if roles == nil {
        return false
    }

    return roles.DelItem(r)
}

func UpdateRole(r *Role) bool {
    if roles == nil {
        return false
    }

    r2 := r.copy()
    return roles.UpdateItem(r2)
}
