// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/18 15:05
// @Package model

package model

import (
	"github.com/a330202207/psychology-healthy-api/internal/model/entity"
)

// RuleEditInput .
type RuleEditInput struct {
	ID      int64
	Name    string  // 菜单名称
	MenuIds []int64 // 菜单ID
	Status  uint    // 状态:10-正常，20-关闭，30-待验证
	Sort    uint    // 排序
}

// RuleBaseInput .
type RuleBaseInput struct {
	ID int64
}

// RuleListInput .
type RuleListInput struct {
	ID       int64
	Name     string // 角色名称
	Status   uint   // 状态:10-正常，20-关闭，30-待验证
	Page     int    `json:"page"`     // 分页码
	PageSize int    `json:"pageSize"` // 分页数量
}

// RuleItem .
type RuleItem struct {
	*entity.SysRole
}

// RuleListOut .
type RuleListOut struct {
	List     []*RuleItem
	Page     int `json:"page"`     // 分页码
	PageSize int `json:"pageSize"` // 分页数量
	Total    int `json:"total"`    // 数据总数
}
