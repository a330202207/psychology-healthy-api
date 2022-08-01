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

var Rule = cRule{}

type cRule struct {
}

// Edit 添加/编辑角色
func (c *cRule) Edit(ctx context.Context, input *v1.RuleEditReq) (res *v1.RuleEditRes, err error) {
	if err = service.Rule().Edit(ctx, &model.RuleEditInput{
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
func (c *cRule) Del(ctx context.Context, input *v1.RuleDelReq) (res *v1.RuleDelRes, err error) {
	if err = service.Rule().Del(ctx, &model.RuleBaseInput{ID: input.ID}); err != nil {
		return
	}
	return
}

// List 角色列表
func (c *cRule) List(ctx context.Context, input *v1.RuleListReq) (res *v1.RuleListRes, err error) {
	out, err := service.Rule().List(ctx, &model.RuleListInput{
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

	res.Rules = out.List
	res.PageInfo.Page = out.Page
	res.PageInfo.PageSize = out.PageSize
	res.PageInfo.Total = out.Total

	return
}

// GetAll 获取所有角色
func (c *cRule) GetAll(ctx context.Context, input *v1.RuleGetAllReq) (res *v1.RuleGetAllRes, err error) {
	if res.Rules, err = service.Rule().GetAll(ctx); err != nil {
		return
	}
	return
}

// Get 获取角色信息
func (c *cRule) Get(ctx context.Context, input *v1.RuleGetReq) (res *v1.RuleGetRes, err error) {
	return
}
