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

// 修改管理员信息

// 修改登陆密码

// 获取管理员信息

// 获取管理员列表

// 重置密码

// 删除管理员
