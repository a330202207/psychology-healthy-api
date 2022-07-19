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
	g.Meta `path:"/menu/edit" method:"post" tag:"menuService" summary:"添加/编辑菜单" tags:"添加菜单"`
}

// MenuEditRes .
type MenuEditRes struct {
}

// MenuDelReq .
type MenuDelReq struct {
	g.Meta `path:"/menu/del" method:"post" tag:"menuService" summary:"删除菜单" tags:"删除菜单"`
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
