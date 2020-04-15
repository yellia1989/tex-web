package model

import (
    "testing"
)

func TestAddPermission(t *testing.T) {
    p := AddPerm("测试权限1",[]string{"GET /api/get","POST  /api/post "})
    if p == nil {
        t.Fatal("添加权限失败")
    }

    p2 := AddPerm("测试权限1",[]string{"GET /api/get","POST  /api/post "})
    if p2 != nil {
        t.Fatal("重复添加权限成功,应该失败")
    }

    if DelPerm(p) == false {
        t.Fatal("删除权限失败")
    }

    p3 := AddPerm("测试权限1",[]string{"GET /api/get","POST  /api/post "})
    if p3 == nil {
        t.Fatal("添加权限失败")
    }

    p3.Name = "测试权限2"
    if UpdatePerm(p3) == false {
        t.Fatal("测试更新权限失败")
    }

    if DelPerm(p3) == false {
        t.Fatal("删除权限失败")
    }
}
