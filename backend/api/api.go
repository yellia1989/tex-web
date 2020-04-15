package api

import (
    "github.com/labstack/echo"
)

func RegisterHandler(group *echo.Group) {
    group.POST("/login", UserLogin)        // 登陆

    group.GET("/menu/list", MenuList)       // 菜单列表
    group.POST("/menu/update", MenuUpdate)  // 更新菜单
    group.GET("/role/list", RoleList)       // 角色列表
    group.GET("/user/list", UserList)       // 用户列表
    group.POST("/user/add", UserAdd)        // 用户增加
    group.POST("/user/del", UserDel)        // 用户删除
    group.POST("/user/update/role", UserUpdateRole)    // 用户角色变更
    group.POST("/user/update", UserUpdate)    // 用户更新
    group.GET("/perm/list", PermList)         // 权限列表
    group.POST("/perm/add", PermAdd)          // 权限增加
    group.POST("/perm/del", PermDel)          // 权限删除
    group.POST("/perm/update", PermUpdate)    // 权限编辑

    /*api.POST("/role/add", api.RoleAdd)      // 角色增加
    api.POST("/role/del", api.RoleDel)      // 角色删除
    api.POST("/role/edit", api.RoleEdit)    // 角色编辑
    */
}
