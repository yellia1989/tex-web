package api

import (
    "github.com/labstack/echo"
)

func RegisterHandler(group *echo.Group) {
    group.POST("/login", UserLogin)        // 登陆

    group.GET("/menu/list", MenuList)       // 菜单列表
    group.POST("/menu/update", MenuUpdate)  // 更新菜单
    group.GET("/role/list", RoleList)       // 角色列表

    /*api.GET("/user/list", api.UserList)     // 用户列表
    api.POST("/user/add", api.UserAdd)      // 用户增加
    api.POST("/user/edit", api.UserEdit)    // 用户编辑
    api.POST("/user/del", api.UserDel)      // 用户删除
    api.POST("/role/add", api.RoleAdd)      // 角色增加
    api.POST("/role/del", api.RoleDel)      // 角色删除
    api.POST("/role/edit", api.RoleEdit)    // 角色编辑
    api.GET("/permission/list", api.PermissionList)     // 权限列表
    api.POST("/permission/add", api.PermissionAdd)      // 权限增加
    api.POST("/permission/del", api.PermissionDel)      // 权限删除
    api.POST("/permission/edit", api.PermissionEdit)    // 权限编辑
    */
}
