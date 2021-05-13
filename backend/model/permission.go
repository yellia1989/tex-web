package model

import (
    "github.com/yellia1989/tex-web/backend/cfg"
    "path"
    "strings"
)

type Permission struct {
    Id  uint32              `json:"id"`
    Name string             `json:"name"`
    Paths []string          `json:"paths"`
    PathString string
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

func (p *Permission) InitPath()  {
    if len(p.PathString)>0 {
        p.Paths = strings.Split(p.PathString,";")
    }else {
        p.Paths = make([]string,0)
    }
}

func (p *Permission) StringifyPath()  {
    if len(p.Paths)>0 {
        p.PathString = strings.Join(p.Paths,";")
    }
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
    db := cfg.StatDb
    if db == nil {
        return nil
    }

    rows, err := db.Query("select id,name,paths from system_perms")
    if err!=nil {
        return nil
    }
    ps := make([]*Permission,0)
    for rows.Next() {
        var perm Permission
        if err := rows.Scan(&perm.Id,&perm.Name,&perm.PathString);err!=nil{
            return nil
        }
        perm.InitPath()
        ps = append(ps,&perm)
    }
    return ps
}

func GetPerm(id uint32) *Permission {
    db := cfg.StatDb
    if db == nil {
        return nil
    }
    perm := &Permission{}
    if err:=db.QueryRow("select id,name,paths from system_perms where id = ?",id).Scan(&perm.Id,&perm.Name,&perm.PathString);err!=nil{
        return nil
    }
    perm.InitPath()
    return perm
}

func AddPerm(name string, paths []string) *Permission {
    db := cfg.StatDb
    if db == nil {
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

    perm := &Permission{}
    if err:=db.QueryRow("select id from system_perms where name = ?",name).Scan(&perm.Id);err==nil {
        return nil
    }

    p := &Permission{
        Name: name,
        Paths: paths,
    }
    p.StringifyPath()
    _,err := db.Exec("insert into system_perms(name,paths) values(?,?)",p.Name,p.PathString)
    if err!=nil{
        return nil
    }
    return p
}

func DelPerm(p *Permission) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    _,err := db.Exec("delete from system_perms where id = ?",p.Id)
    if err!=nil {
        return false
    }
    return true
}

func UpdatePerm(p *Permission) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    p.StringifyPath()
    _,err := db.Exec("update system_perms set name = ?,paths = ? where id = ?",p.Name,p.PathString,p.Id)
    if err!=nil{
        return false
    }
    return true
}
