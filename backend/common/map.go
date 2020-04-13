// 实现一个简单地内存数据结构由于保存数据量小的
// 有落地需求的结构

package common

import (
    "sync"
    "sync/atomic"
    "sort"
    "encoding/json"
    "github.com/yellia1989/tex-go/tools/util"
)

type Map struct {
    items sync.Map
    path string
    uid uint32
}

type Item interface {
    GetId() uint32
    SetId(id uint32)
}

func (m *Map) GetItem(key interface{}) Item {
    v, ok := m.items.Load(key)
    if !ok {
        return nil
    }
    return v.(Item)
}

type itemBy []Item
func (a itemBy) Len() int           { return len(a) }
func (a itemBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a itemBy) Less(i, j int) bool { return a[i].GetId() < a[j].GetId() }

func (m *Map) GetItems(filter func(key, v interface{}) bool) []Item {
    items := make([]Item,0)
    m.items.Range(func (key, v interface{}) bool {
        if filter(key, v) {
            items = append(items, v.(Item))
        }
        return true
    })

    sort.Sort(itemBy(items))
    return items
}

func (m *Map) AddItem(u Item) bool {
    // id必须有内部分配
    if u.GetId() != 0 {
        return false
    }
    u.SetId(atomic.AddUint32(&m.uid,1))

    m.items.Store(u.GetId(), u)
    return m.save() == nil
}

func (m *Map) DelItem(u Item) bool {
    _, ok := m.items.Load(u.GetId())
    if !ok {
        return false
    }
    m.items.Delete(u.GetId())
    return m.save() == nil
}

func (m *Map) DelAllItem() bool {
    ids := make([]uint32,0)
    m.items.Range(func (key, v interface{}) bool {
        ids = append(ids, key.(uint32))
        return true
    })

    for _, id := range ids {
        m.items.Delete(id)
    }
    return m.save() == nil
}

func (m *Map) UpdateItem(u Item) bool {
    if m.GetItem(u.GetId()) == nil {
        return false
    }
    m.items.Store(u.GetId(), u)
    return m.save() == nil
}

func (m *Map) save() error {
    us := m.GetItems(func (key,v interface{})bool{return true})

    bs, err := json.MarshalIndent(us, "", "  ")
    if err != nil {
        return err
    }
    return util.SaveToFile(m.path, bs, false)
}

func NewMap(path string, items []Item) *Map {
    m := &Map{path: path}
    for _, u := range items {
        m.items.Store(u.GetId(), u)
        if u.GetId() > m.uid {
            m.uid = u.GetId()
        }
    }
    return m
}
