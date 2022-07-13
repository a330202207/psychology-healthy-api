// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for table sys_role.
type SysRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleColumns defines and stores column names for table sys_role.
type SysRoleColumns struct {
	Id                string // 角色ID
	Name              string // 角色名称
	Key               string // 角色权限字符串
	MenuCheckStrictly string // 菜单树选择项是否关联显示：10-显示，20-隐藏
	Sort              string // 排序
	Status            string // 角色状态：10-开启，20-关闭
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

//  sysRoleColumns holds the columns for table sys_role.
var sysRoleColumns = SysRoleColumns{
	Id:                "id",
	Name:              "name",
	Key:               "key",
	MenuCheckStrictly: "menu_check_strictly",
	Sort:              "sort",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		group:   "-t",
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}