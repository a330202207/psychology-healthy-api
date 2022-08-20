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
	if err = service.Menu().Edit(ctx, &model.MenuEditInput{
		ID:        input.ID,
		PID:       input.PID,
		Name:      input.Name,
		Code:      input.Code,
		Icon:      input.Icon,
		Status:    input.Status,
		Type:      input.Type,
		Perms:     input.Perms,
		Sort:      input.Sort,
		Path:      input.Path,
		Component: input.Component,
		Query:     input.Query,
		IsCache:   input.IsCache,
		IsLeaf:    input.IsLeaf,
		Level:     input.Level,
		Tree:      input.Tree,
	}); err != nil {
		return
	}
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
	out, err := service.Menu().List(ctx, &model.MenuListInput{
		Name:   input.Name,
		Status: input.Status,
		Type:   input.Type,
		PageBaseInfo: model.PageBaseInfo{
			Page:     input.Page,
			PageSize: input.PageSize,
		},
	})

	if err != nil {
		return
	}

	res = &v1.MenuListRes{
		Menus: out.List,
		PageInfo: v1.PageInfo{
			PageBaseInfo: v1.PageBaseInfo{
				Page:     out.Page,
				PageSize: out.PageSize,
			},
			Total: out.Total,
		},
	}

	return
}

// Get 获取菜单信息
func (c *cMenu) Get(ctx context.Context, input *v1.MenuGetReq) (res *v1.MenuGetRes, err error) {
	return
}

// GetAll 获取所有菜单
func (c *cMenu) GetAll(ctx context.Context, input *v1.MenuGetAllReq) (res *v1.MenuGetAllRes, err error) {
	out, err := service.Menu().GetAll(ctx)
	if err != nil {
		return
	}
	res = &v1.MenuGetAllRes{
		List: out,
	}

	return
}
