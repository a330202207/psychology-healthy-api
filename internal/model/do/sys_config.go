// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta    `orm:"table:sys_config, do:true"`
	Id        interface{} // 配置ID
	Name      interface{} // 参数名称
	Key       interface{} // 参数键名
	Value     interface{} // 参数键值
	Status    interface{} // 状态：10-开启，20-关闭
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
