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

// RoleEditReq .
type RoleEditReq struct {
	g.Meta  `path:"/role/edit" method:"post" summary:"添加/编辑角色" tags:"角色管理"`
	ID      int64   `json:"id" dc:"角色ID"`
	MenuIds []int64 `json:"menuIds" v:"required#菜单不能为空" dc:"菜单ID"`
	Name    string  `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status  uint    `json:"status" v:"required#角色状态不能为空" dc:"角色状态：10-开启，20-关闭"`
	Sort    uint    `json:"sort" dc:"排序"`
}

// RoleEditRes .
type RoleEditRes struct {
	g.Meta `mime:"application/json"`
}

// RoleDelReq .
type RoleDelReq struct {
	g.Meta `path:"/role/del" method:"post" summary:"删除角色" tags:"角色管理"`
	ID     int64 `json:"id"`
}

// RoleDelRes .
type RoleDelRes struct {
	g.Meta `mime:"application/json"`
}

// RoleListReq .
type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" summary:"角色列表" tags:"角色管理"`
	ID     int64  `json:"id"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status uint   `json:"status" v:"required#角色状态不能为空" dc:"角色状态：10-开启，20-关闭"`
	PageBaseInfo
}

// RoleListRes .
type RoleListRes struct {
	g.Meta    `mime:"application/json"`
	Roles     []*model.RoleItem `json:"list"`
	*PageInfo `json:"pageInfo"`
}

// RoleGetReq .
type RoleGetReq struct {
	g.Meta `path:"/role/get" method:"get" summary:"角色信息" tags:"角色管理"`
	ID     int64 `json:"id"`
}

// RoleGetRes .
type RoleGetRes struct {
}

// RoleGetAllReq .
type RoleGetAllReq struct {
	g.Meta `path:"/role/getAll" method:"get" summary:"获取所有角色" tags:"角色管理"`
}

// RoleGetAllRes .
type RoleGetAllRes struct {
	g.Meta `mime:"application/json"`
	Roles  []*model.RoleItem `json:"list"`
}
