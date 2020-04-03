package model

import (
    "fmt"
    "encoding/json"
    "github.com/yellia1989/tex-go/tools/util"
    cm "github.com/yellia1989/tex-web/backend/common"
)

var roles *cm.Map
func init() {
    bs, err := util.LoadFromFile("data/roles.json")
    if err != nil {
        fmt.Printf("roles init failed, %s", err.Error())
    }
    items := make([]*Role,0)
    err = json.Unmarshal(bs, &items)
    if err != nil {
        fmt.Printf("roles init failed, %s", err.Error())
    }

    items2 := make([]cm.Item,0)
    for _, item := range items {
        items2 = append(items2, item)
    }
    roles = cm.NewMap("data/roles.json", items2)
}

type Role struct {
    Id  uint32              `json:"value"`
    Name string             `json:"name"`
    Desc string             `json:"desc"`
}
func (r *Role) GetId() uint32 {
    return r.Id
}
func (r *Role) SetId(id uint32) {
    r.Id = id
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
        r := *(item.(*Role))
        rs = append(rs, &r)
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

    r := *(item.(*Role))
    return &r
}
