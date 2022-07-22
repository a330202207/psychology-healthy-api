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
	ID        int64  `v:"required#菜单ID不能为空" json:"id"`
	Name      string `v:"required#菜单名称不能为空" json:"name"`
	URL       string `json:"url"`
	Icon      string `json:"icon"`
	Status    uint   `json:"status"`
	Sort      uint   `json:"sort"`
	PID       int64  `json:"pid"`
	Type      uint   `json:"type"`
	Perms     string `json:"perms"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Query     string `json:"query"`
	IsCache   bool   `json:"isCache"`
	IsLeaf    bool   `json:"isLeaf"`
	Level     uint   `json:"level"`
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
