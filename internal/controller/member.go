// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:26
// @Package controller

package controller

import (
	"context"

	v1 "github.com/a330202207/psychology-healthy-api/api/v1"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

var Member = cMember{}

type cMember struct {
}

// Edit 添加管理员
func (c *cMember) Edit(ctx context.Context, input *v1.MemberEditReq) (res *v1.MemberEditRes, err error) {
	if err = service.Member().Edit(ctx, &model.MemberEditInput{
		ID:         input.ID,
		RuleIds:    input.RuleIds,
		UserName:   input.UserName,
		Passwd:     input.Passwd,
		Type:       input.Type,
		Status:     input.Status,
		NickName:   input.NickName,
		Avatar:     input.Avatar,
		Sex:        input.Sex,
		Email:      input.Email,
		ProvinceID: input.ProvinceID,
		CityID:     input.CityID,
		AreaID:     input.AreaID,
		Address:    input.Address,
		Mobile:     input.Mobile,
		Birthday:   input.Birthday,
	}); err != nil {
		return
	}

	return
}

// UpdatePassWd 修改密码
func (c *cMember) UpdatePassWd(ctx context.Context, input *v1.MemberUpdatePassWdReq) (res *v1.MemberUpdatePassWdRes, err error) {
	if err = service.Member().UpdatePassWd(ctx, &model.MemberUpdatePassWdInput{
		ID:     input.ID,
		Passwd: input.Passwd,
	}); err != nil {
		return
	}
	return
}

// Get 获取管理员信息
func (c *cMember) Get(ctx context.Context, input *v1.MemberGetReq) (res *v1.MemberGetRes, err error) {
	return
}

// List 获取管理员列表
func (c *cMember) List(ctx context.Context, input *v1.MemberListReq) (res *v1.MemberListRes, err error) {
	return
}

// ResetPassWd 重置密码
func (c *cMember) ResetPassWd(ctx context.Context, input *v1.MemberResetPassWdReq) (res *v1.MemberResetPassWdRes, err error) {

	return
}

// Del 删除管理员
func (c *cMember) Del(ctx context.Context, input *v1.MemberDelReq) (res *v1.MemberDelRes, err error) {
	if err = service.Member().Del(ctx, &model.MemberBaseInput{ID: input.ID}); err != nil {
		return
	}
	return
}
