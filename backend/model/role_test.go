package model

import (
    "testing"
)

func TestAddRole(t *testing.T) {
    r := AddRole("测试角色1",[]uint32{1,2})
    if r == nil {
        t.Fatal("添加角色失败")
    }

    r2 := AddRole("测试角色1",[]uint32{1,2})
    if r2 != nil {
        t.Fatal("重复添加角色成功,应该失败")
    }

    if DelRole(r) == false {
        t.Fatal("删除角色失败")
    }

    r3 := AddRole("测试角色1",[]uint32{1,2})
    if r3 == nil {
        t.Fatal("添加角色失败")
    }

    r3.Name = "测试权限2"
    if UpdateRole(r3) == false {
        t.Fatal("测试更新角色失败")
    }

    if DelRole(r3) == false {
        t.Fatal("删除角色失败")
    }
}
