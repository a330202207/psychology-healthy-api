// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for table sys_menu.
type SysMenuColumns struct {
	Id        string // 菜单ID
	Pid       string // 父菜单ID
	Name      string // 菜单名称
	Code      string // 菜单编码
	Icon      string // 菜单图标
	Type      string // 菜单类型（M目录 C菜单 F按钮）
	Perms     string // 权限标识
	Path      string // 路由地址
	Component string // 组件路径
	Query     string // 路由参数
	IsCache   string // 是否缓存：10-是，20-否
	IsVisible string // 是否隐藏：10-是，20-否
	Level     string // 级别
	Tree      string // 树
	Sort      string // 排序
	Status    string // 菜单状态：10-开启，20-关闭
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

//  sysMenuColumns holds the columns for table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Code:      "code",
	Icon:      "icon",
	Type:      "type",
	Perms:     "perms",
	Path:      "path",
	Component: "component",
	Query:     "query",
	IsCache:   "is_cache",
	IsVisible: "is_visible",
	Level:     "level",
	Tree:      "tree",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "-t",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}