// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/18 15:05
// @Package model

package model

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
}

// MenuListOut .
type MenuListOut struct {
}

// MenuBaseInput .
type MenuBaseInput struct {
	ID int64
}

// MenuGetOut .
type MenuGetOut struct {
}
