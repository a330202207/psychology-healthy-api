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

var Role = cRole{}

type cRole struct {
}

// Edit 添加/编辑角色
func (c *cRole) Edit(ctx context.Context, input *v1.RoleEditReq) (res *v1.RoleEditRes, err error) {
	if err = service.Role().Edit(ctx, &model.RoleEditInput{
		ID:      input.ID,
		Name:    input.Name,
		MenuIds: input.MenuIds,
		Status:  input.Status,
		Sort:    input.Sort,
	}); err != nil {
		return
	}
	return
}

// Del 删除角色
func (c *cRole) Del(ctx context.Context, input *v1.RoleDelReq) (res *v1.RoleDelRes, err error) {
	if err = service.Role().Del(ctx, &model.RoleBaseInput{ID: input.ID}); err != nil {
		return
	}
	return
}

// List 角色列表
func (c *cRole) List(ctx context.Context, input *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	out, err := service.Role().List(ctx, &model.RoleListInput{
		ID:     input.ID,
		Name:   input.Name,
		Status: input.Status,
		PageBaseInfo: model.PageBaseInfo{
			Page:     input.Page,
			PageSize: input.PageSize,
		},
	})
	if err != nil {
		return
	}

	res.Roles = out.List
	res.PageInfo.Page = out.Page
	res.PageInfo.PageSize = out.PageSize
	res.PageInfo.Total = out.Total

	return
}

// GetAll 获取所有角色
func (c *cRole) GetAll(ctx context.Context, input *v1.RoleGetAllReq) (res *v1.RoleGetAllRes, err error) {
	if res.Roles, err = service.Role().GetAll(ctx); err != nil {
		return
	}
	return
}

// Get 获取角色信息
func (c *cRole) Get(ctx context.Context, input *v1.RoleGetReq) (res *v1.RoleGetRes, err error) {
	return
}
