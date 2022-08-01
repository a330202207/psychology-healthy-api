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

// MenuEditInput .
type MenuEditInput struct {
	ID        int64
	PID       int64
	Name      string
	Code      string
	Icon      string
	Status    uint
	Type      uint
	Perms     string
	Sort      uint
	Path      string
	Component string
	Query     string
	IsCache   bool
	IsLeaf    bool
	Level     uint
	Tree      string
}

// MenuListInput .
type MenuListInput struct {
	Name   string
	Status uint
	Type   uint
	PageBaseInfo
}

// MenuItem .
type MenuItem struct {
	*entity.SysMenu
}

// MenuListOut .
type MenuListOut struct {
	List []*MenuItem
	PageInfo
}

// MenuBaseInput .
type MenuBaseInput struct {
	ID int64
}

// MenuGetOut .
type MenuGetOut struct {
}

// MenuTree 菜单树
type MenuTree struct {
	*entity.SysMenu
	Children []*MenuTree `json:"children"`
}

// LabelMenuTree 菜单kl树
type LabelMenuTree struct {
	entity.SysMenu
	Key      int64            `json:"key"       description:"键名"`
	Label    string           `json:"label"       description:"键标签"`
	Children []*LabelMenuTree `json:"children"`
}
