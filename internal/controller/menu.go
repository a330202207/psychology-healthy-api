// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 15:56
// @Package controller

package controller

import (
	"context"

	v1 "github.com/a330202207/psychology-healthy-api/api/v1"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

var Menu = cMenu{}

type cMenu struct {
}

// Edit 添加/编辑菜单
func (c *cMenu) Edit(ctx context.Context, input *v1.MenuEditReq) (res *v1.MenuEditRes, err error) {
	return
}

// Del 删除菜单
func (c *cMenu) Del(ctx context.Context, input *v1.MenuDelReq) (res *v1.MenuDelRes, err error) {
	if err = service.Menu().Del(ctx, &model.MenuBaseInput{ID: input.ID}); err != nil {
		return
	}
	return
}

// List 菜单列表
func (c *cMenu) List(ctx context.Context, input *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	return
}
