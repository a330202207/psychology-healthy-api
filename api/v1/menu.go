// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/18 14:53
// @Package v1

package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/a330202207/psychology-healthy-api/internal/model"
)

// MenuEditReq .
type MenuEditReq struct {
	g.Meta    `path:"/menu/edit" method:"post" summary:"添加/编辑菜单" tags:"菜单管理"`
	ID        int64  `json:"id" dc:"菜单ID"`
	Name      string `v:"required#菜单名称不能为空" json:"name" dc:"菜单名称"`
	Icon      string `json:"icon" dc:"菜单图标"`
	Code      string `v:"required#菜单编码不能为空" json:"code" dc:"菜单编码"`
	Status    uint   `v:"required#菜单状态不能为空" json:"status" dc:"菜单状态：10-开启，20-关闭"`
	Sort      uint   `json:"sort" dc:"排序"`
	PID       int64  `v:"required#菜单名称不能为空" json:"pid" dc:"父级菜单ID"`
	Type      uint   `v:"required#菜单名称不能为空" json:"type" dc:"菜单类型：1-目录，2-菜单，3-按钮"`
	Perms     string `json:"perms" dc:"权限标识"`
	Path      string `v:"required#路由地址不能为空" json:"path" dc:"路由地址"`
	Component string `json:"component" dc:"组件路径"`
	Query     string `json:"query" dc:"路由参数"`
	IsCache   bool   `v:"required#是否缓存不能为空" json:"isCache" dc:"是否缓存"`
	IsLeaf    bool   `v:"required#是否隐藏不能为空" json:"isLeaf" dc:"是否隐藏"`
	Level     uint   `v:"required#菜单级别不能为空" json:"level" dc:"菜单级别"`
	Tree      string `json:"tree" dc:"菜单树"`
}

// MenuEditRes .
type MenuEditRes struct {
	g.Meta `mime:"application/json"`
}

// MenuDelReq .
type MenuDelReq struct {
	g.Meta `path:"/menu/del" method:"post" summary:"删除菜单" tags:"菜单管理"`
	ID     int64 `v:"required#菜单ID不能为空" json:"id"`
}

// MenuDelRes .
type MenuDelRes struct {
	g.Meta `mime:"application/json"`
}

// MenuListReq .
type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" summary:"菜单列表" tags:"菜单管理"`
	Name   string `json:"name" dc:"菜单名称"`
	Status uint   `json:"status" dc:"菜单状态：10-开启，20-关闭"`
	Type   uint   `json:"type" dc:"菜单类型：1-目录，2-菜单，3-按钮"`
	PageBaseInfo
}

// MenuListRes .
type MenuListRes struct {
	g.Meta   `mime:"application/json"`
	Menus    []*model.MenuItem `json:"list"`
	PageInfo *PageInfo         `json:"pageInfo"`
}

// MenuGetReq .
type MenuGetReq struct {
	g.Meta `path:"/menu/get" method:"get" summary:"获取菜单信息" tags:"菜单管理"`
	ID     int64 `v:"required#菜单ID不能为空" json:"id"`
}

// MenuGetRes .
type MenuGetRes struct {
	g.Meta    `mime:"application/json"`
	ID        int64  `json:"id" dc:"菜单ID"`
	Name      string `json:"name" dc:"菜单名称"`
	Icon      string `json:"icon" dc:"菜单图标"`
	Code      string `json:"code" dc:"菜单编码"`
	Status    uint   `json:"status" dc:"菜单状态：10-开启，20-关闭"`
	Sort      uint   `json:"sort" dc:"排序"`
	PID       int64  `json:"pid" dc:"父级菜单ID"`
	Type      uint   `json:"type" dc:"菜单类型：1-目录，2-菜单，3-按钮"`
	Perms     string `json:"perms" dc:"权限标识"`
	Path      string `json:"path" dc:"路由地址"`
	Component string `json:"component" dc:"组件路径"`
	Query     string `json:"query" dc:"路由参数"`
	IsCache   bool   `json:"isCache" dc:"是否缓存"`
	IsLeaf    bool   `json:"isLeaf" dc:"是否隐藏"`
	Level     uint   `json:"level" dc:"菜单级别"`
	Tree      string `json:"tree" dc:"菜单树"`
}

// MenuGetAllReq .
type MenuGetAllReq struct {
	g.Meta `path:"/menu/getAll" method:"get" summary:"获取所有菜单" tags:"菜单管理"`
}

// MenuGetAllRes .
type MenuGetAllRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.MenuTree `json:"list"`
}
