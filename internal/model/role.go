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

// RoleEditInput .
type RoleEditInput struct {
	ID      int64
	Name    string  // 菜单名称
	MenuIds []int64 // 菜单ID
	Status  uint    // 状态:10-正常，20-关闭，30-待验证
	Sort    uint    // 排序
}

// RoleBaseInput .
type RoleBaseInput struct {
	ID int64
}

// RoleListInput .
type RoleListInput struct {
	ID     int64
	Name   string // 角色名称
	Status uint   // 状态:10-正常，20-关闭，30-待验证
	PageBaseInfo
}

// RoleItem .
type RoleItem struct {
	*entity.SysRole
}

// RoleListOut .
type RoleListOut struct {
	List []*RoleItem
	PageInfo
}
