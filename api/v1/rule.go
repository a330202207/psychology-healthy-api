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

// RuleEditReq .
type RuleEditReq struct {
	g.Meta  `path:"/rule/edit" method:"post" tag:"ruleService" summary:"添加/编辑菜单" tags:"添加/编辑菜单"`
	ID      int64   `json:"id"`
	MenuIds []int64 `json:"menuIds" v:"required#菜单不能为空" dc:"菜单ID"`
	Name    string  `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status  uint    `json:"status" v:"required#角色状态不能为空" dc:"角色状态：10-开启，20-关闭"`
	Sort    uint    `json:"sort" dc:"排序"`
}

// RuleEditRes .
type RuleEditRes struct {
}

// RuleDelReq .
type RuleDelReq struct {
	g.Meta `path:"/rule/del" method:"post" tag:"ruleService" summary:"删除角色" tags:"删除角色"`
	ID     int64 `json:"id"`
}

// RuleDelRes .
type RuleDelRes struct {
}

// RuleListReq .
type RuleListReq struct {
	g.Meta `path:"/rule/list" method:"get" tag:"ruleService" summary:"角色列表" tags:"角色列表"`
}

// RuleListRes .
type RuleListRes struct {
}
