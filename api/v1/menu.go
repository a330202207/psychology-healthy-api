// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/18 14:53
// @Package v1

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MenuEditReq .
type MenuEditReq struct {
	g.Meta    `path:"/menu/edit" method:"post" tag:"menuService" summary:"添加/编辑菜单" tags:"添加/编辑菜单"`
	ID        int64  `v:"required#菜单ID不能为空" json:"id" dc:"菜单ID"`
	Name      string `v:"required#菜单名称不能为空" json:"name" dc:"菜单名称"`
	Icon      string `json:"icon" dc:"菜单图标"`
	Code      string `v:"required#菜单编码不能为空" json:"code" dc:"菜单编码"`
	Status    uint   `v:"required#菜单状态不能为空" json:"status" dc:"菜单状态：10-开启，20-关闭"`
	Sort      uint   `json:"sort" dc:"排序"`
	PID       int64  `v:"required#菜单名称不能为空" json:"pid" dc:"父级菜单ID"`
	Type      uint   `v:"required#菜单名称不能为空" json:"type" dc:"菜单类型：1-目录，2-菜单，3-按钮"`
	Perms     string `json:"perms" dc:"权限标识"`
	Path      string `v:"required#路由地址不能为空" json:"path"`
	Component string `json:"component"`
	Query     string `json:"query"`
	IsCache   bool   `v:"required#是否缓存不能为空" json:"isCache"`
	IsLeaf    bool   `v:"required#是否隐藏不能为空" json:"isLeaf"`
	Level     uint   `v:"required#菜单级别不能为空" json:"level"`
	Tree      string `json:"tree"`
}

// MenuEditRes .
type MenuEditRes struct {
}

// MenuDelReq .
type MenuDelReq struct {
	g.Meta `path:"/menu/del" method:"post" tag:"menuService" summary:"删除菜单" tags:"删除菜单"`
	ID     int64 `v:"required#菜单ID不能为空" json:"id"`
}

// MenuDelRes .
type MenuDelRes struct {
}

// MenuListReq .
type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tag:"menuService" summary:"菜单列表" tags:"菜单列表"`
}

// MenuListRes .
type MenuListRes struct {
}
