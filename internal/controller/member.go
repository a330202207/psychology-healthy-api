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

// Add 添加管理员
func (c *cMember) Add(ctx context.Context, input *v1.MemberEditReq) (res *v1.MemberEditRes, err error) {
	service.Member().Edit(ctx, &model.MemberInput{
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
	})

	return
}

// UpdatePassWd 修改密码
func (c *cMember) UpdatePassWd(ctx context.Context, input *v1.MemberUpdatePassWdReq) (res *v1.MemberUpdatePassWdRes, err error) {

	return
}

// GetInfo 获取管理员信息
func (c *cMember) GetInfo(ctx context.Context, input *v1.MemberGetInfoReq) (res *v1.MemberGetInfoRes, err error) {
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

	return
}
