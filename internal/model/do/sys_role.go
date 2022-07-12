// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta            `orm:"table:sys_role, do:true"`
	Id                interface{} // 角色ID
	Name              interface{} // 角色名称
	Key               interface{} // 角色权限字符串
	MenuCheckStrictly interface{} // 菜单树选择项是否关联显示：10-显示，20-隐藏
	Sort              interface{} // 排序
	Status            interface{} // 角色状态：10-开启，20-关闭
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
