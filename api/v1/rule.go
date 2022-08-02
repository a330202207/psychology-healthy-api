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

// RuleEditReq .
type RuleEditReq struct {
	g.Meta  `path:"/rule/edit" method:"post" summary:"添加/编辑角色" tags:"角色管理"`
	ID      int64   `json:"id" dc:"角色ID"`
	MenuIds []int64 `json:"menuIds" v:"required#菜单不能为空" dc:"菜单ID"`
	Name    string  `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status  uint    `json:"status" v:"required#角色状态不能为空" dc:"角色状态：10-开启，20-关闭"`
	Sort    uint    `json:"sort" dc:"排序"`
}

// RuleEditRes .
type RuleEditRes struct {
	g.Meta `mime:"application/json"`
}

// RuleDelReq .
type RuleDelReq struct {
	g.Meta `path:"/rule/del" method:"post" summary:"删除角色" tags:"角色管理"`
	ID     int64 `json:"id"`
}

// RuleDelRes .
type RuleDelRes struct {
	g.Meta `mime:"application/json"`
}

// RuleListReq .
type RuleListReq struct {
	g.Meta `path:"/rule/list" method:"get" summary:"角色列表" tags:"角色管理"`
	ID     int64  `json:"id"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status uint   `json:"status" v:"required#角色状态不能为空" dc:"角色状态：10-开启，20-关闭"`
	PageBaseInfo
}

// RuleListRes .
type RuleListRes struct {
	g.Meta   `mime:"application/json"`
	Rules    []*model.RuleItem `json:"list"`
	PageInfo *PageInfo         `json:"pageInfo"`
}

// RuleGetReq .
type RuleGetReq struct {
	g.Meta `path:"/rule/get" method:"get" summary:"角色信息" tags:"角色管理"`
	ID     int64 `json:"id"`
}

// RuleGetRes .
type RuleGetRes struct {
}

// RuleGetAllReq .
type RuleGetAllReq struct {
	g.Meta `path:"/rule/getAll" method:"get" summary:"获取所有角色" tags:"角色管理"`
}

// RuleGetAllRes .
type RuleGetAllRes struct {
	g.Meta `mime:"application/json"`
	Rules  []*model.RuleItem `json:"list"`
}
