package model

import (
    "testing"
)

func TestAddUser(t *testing.T) {
    u := AddUser("yellia", "pwd", 0)
    if u == nil || u.Id != 1 {
        t.Fatalf("add user failed, should success, %v", u)
    }

    u2 := GetUserByUserName(u.UserName)
    if u2.Id != u.Id {
        t.Fatal("GetUserByUserName failed")
    }

    u3 := GetUser(u.Id)
    if u3.Id != u.Id {
        t.Fatal("GetUser failed")
    }

    if ret := AddUser("yellia", "pwd", 0); ret != nil {
        t.Fatal("add user success, should failed")
    }

    if ret := DelUser(u); ret == false {
        t.Fatal("DelUser failed, should success")
    }

    u = AddUser("yellia", "pwd", 0)
    if u == nil || u.Id != 2 {
        t.Fatal("add user failed, should success")
    }

    u.Id = 10
    if ret := UpdateUser(u); ret == true {
        t.Fatal("UpdateUser success, should failed")
    }

    if ret := DelAllUser(); ret == false {
        t.Fatal("DelAllUser failed, should success")
    }
}
