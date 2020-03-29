package model

type User struct {
    Id int              `json:"id"`
    Name string         `json:"name"`
    Username string     `json:"username"`
    Password string     `json:"password"`
}

type Role struct {
    Id int              `json:"id"`
    Name string         `json:"name"`
}

type Permission struct {
    Id int              `json:"id"`
    Name string         `json:"name"`
    Desc string         `json:"desc"`
}
