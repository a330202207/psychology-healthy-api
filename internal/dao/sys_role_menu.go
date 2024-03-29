// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/a330202207/psychology-healthy-api/internal/dao/internal"
)

// internalSysRoleMenuDao is internal type for wrapping internal DAO implements.
type internalSysRoleMenuDao = *internal.SysRoleMenuDao

// sysRoleMenuDao is the data access object for table sys_role_menu.
// You can define custom methods on it to extend its functionality as you wish.
type sysRoleMenuDao struct {
	internalSysRoleMenuDao
}

var (
	// SysRoleMenu is globally public accessible object for table sys_role_menu operations.
	SysRoleMenu = sysRoleMenuDao{
		internal.NewSysRoleMenuDao(),
	}
)

// Fill with you ideas below.

// UpdateRoleMenuByIds 更新角色菜单
func (d *sysRoleMenuDao) UpdateRoleMenuByIds(ctx context.Context, roleId int64, menusIds []int64, tx *gdb.TX) (err error) {
	if _, err = d.Ctx(ctx).TX(tx).Where("role_id", roleId).Delete(); err != nil {
		return
	}

	var menuData []map[string]interface{}

	for _, val := range menusIds {
		var data = make(map[string]interface{})
		data["role_id"] = roleId
		data["menu_id"] = val

		menuData = append(menuData, data)
	}

	if _, err = d.Ctx(ctx).TX(tx).Insert(menuData); err != nil {
		return
	}

	return
}
