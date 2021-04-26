package model

import (
    "database/sql"
    "fmt"
    "github.com/yellia1989/tex-web/backend/cfg"
    "strconv"
    "strings"
)

type Role struct {
    Id  uint32              `json:"id"`
    Name string             `json:"name"`
    Perms []uint32          `json:"perms"`
    PermString string
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

func (p *Role) InitPath()  {
    p.Perms = make([]uint32,0)
    if len(p.PermString)>0{
        tempStrArr := strings.Split(p.PermString,",")
        for _,v:= range tempStrArr{
            if id,err:=strconv.ParseUint(v,10,32);err==nil{
                p.Perms = append(p.Perms, uint32(id))
            }
        }
    }
}

func (p *Role) StringifyPath()  {
    if len(p.Perms)>0{
        tempStrArr := make([]string,0)
        for _,v := range p.Perms{
            tempStrArr = append(tempStrArr, strconv.FormatUint(uint64(v),10))
        }
        p.PermString = strings.Join(tempStrArr,",")
    }
}

func GetRoles() []*Role {
    db := cfg.StatDb
    if db == nil {
        return nil
    }

    rows, err := db.Query("select id,name,(select group_concat(perm_id) from role_perm where role_id = r.id) as perms from system_role r")
    if err!=nil{
        return nil
    }
    rs := make([]*Role,0)
    var nullStr sql.NullString
    for rows.Next() {
        var role Role
        if err := rows.Scan(&role.Id,&role.Name,&nullStr);err!=nil{
            fmt.Println(err)
            return nil
        }
        if nullStr.Valid{
            role.PermString = nullStr.String
        }
        role.InitPath()
        rs = append(rs,&role)
    }
    return rs
}

func GetRole(id uint32) *Role {
    db := cfg.StatDb
    if db == nil {
        return nil
    }

    role := &Role{}
    var nullStr sql.NullString
    if err:=db.QueryRow("select id,name,(select group_concat(perm_id) from role_perm where role_id = r.id) as perms from system_role r where id = ?",id).Scan(&role.Id,&role.Name,&nullStr);err!=nil{
        return nil
    }
    if nullStr.Valid{
        role.PermString = nullStr.String
    }
    role.InitPath()
    return role
}

func AddRole(name string, perms []uint32) *Role {
    db := cfg.StatDb
    if db == nil {
        return nil
    }

    role := &Role{}
    if err:=db.QueryRow("select id from system_role where name = ?",name).Scan(&role.Id);err==nil{
        return nil
    }

    r := &Role{
        Name: name,
        Perms: perms,
    }

    tx,err:= db.Begin()
    if err!= nil{
        return nil
    }

    result,err := tx.Exec("insert into system_role(name) values(?)", r.Name)
    if err!=nil{
        tx.Rollback()
        return nil
    }
    insertId,err:= result.LastInsertId()
    if err!=nil{
        tx.Rollback()
        return nil
    }
    r.Id = uint32(insertId)
    for _,v:= range r.Perms {
        _,err := tx.Exec("insert into role_perm(role_id,perm_id) values(?,?)", r.Id,v)
        if err!=nil{
            tx.Rollback()
            return nil
        }
    }
    err = tx.Commit()
    if err!=nil{
        tx.Rollback()
        return nil
    }
    return r
}

func DelRole(r *Role) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    tx,err:= db.Begin()
    if err!= nil{
        return false
    }
    _,err = tx.Exec("delete from system_role where id = ?",r.Id)
    if err!=nil{
        tx.Rollback()
        return false
    }
    _,err = tx.Exec("delete from role_perm where role_id = ?",r.Id)
    if err!=nil{
        tx.Rollback()
        return false
    }
    err = tx.Commit()
    if err!=nil{
        tx.Rollback()
        return false
    }
    return true
}

func UpdateRole(r *Role) bool {
    db := cfg.StatDb
    if db == nil {
        return false
    }
    tx,err:= db.Begin()
    if err!= nil{
        return false
    }
    _,err = tx.Exec("update system_role set name = ? where id = ?",r.Name,r.Id)
    if err!=nil{
        tx.Rollback()
        return false
    }
    _,err = tx.Exec("delete from role_perm where role_id = ?",r.Id)
    if err!=nil{
        tx.Rollback()
        return false
    }
    for _,v:= range r.Perms {
        _,err := tx.Exec("insert into role_perm(role_id,perm_id) values(?,?)", r.Id,v)
        if err!=nil{
            tx.Rollback()
            return false
        }
    }
    err = tx.Commit()
    if err!=nil{
        tx.Rollback()
        return false
    }
    return true
}

func DelRolePerm(perms []uint32) {
    db := cfg.StatDb
    if db == nil {
        return
    }
    permStr := make([]string,len(perms))
    for _,v := range perms{
        permStr = append(permStr,strconv.FormatUint(uint64(v),10))
    }
    sql := "delete role_perm where perm_id in  (%s)"
    ids := strings.Join(permStr,",")
    sql = fmt.Sprintf(sql,ids)
    db.Exec(sql)
}
