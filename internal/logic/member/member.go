// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:28
// @Package admin

package member

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
	"github.com/a330202207/psychology-healthy-api/utility/cache"
)

type sMember struct {
}

func init() {
	service.RegisterMember(New())
}

func New() *sMember {
	return &sMember{}
}

// Edit 新增/修改管理员
func (s *sMember) Edit(ctx context.Context, in *model.MemberEditInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-member-Edit")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysMember.IsUniqueName(ctx, in.ID, in.UserName)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member Edit sql error:", err.Error())
		err = errors.New("操作失败[001]")
		return err
	}

	if !ok {
		err = gerror.New("帐号已存在")
		return err
	}

	passWd, err := service.Auth().EncryptPass(in.Passwd)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member EncryptPass error:", err.Error())
		return err
	}

	in.Passwd = string(passWd)

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member Edit Transaction error:", err.Error())
		err = errors.New("操作失败[002]")
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	if in.ID > 0 {

		// 更新用户
		if _, err = dao.SysMember.Ctx(ctx).TX(tx).Where("id", in.ID).OmitEmpty().Update(in); err != nil {
			g.Log(logger).Error(ctx, "service Member Edit update mysql error:", err.Error())
			err = errors.New("操作失败[003]")
			return
		}
	} else {
		in.Passwd = string(passWd)
		in.ID, err = dao.SysMember.Ctx(ctx).TX(tx).OmitEmpty().InsertAndGetId(in)
		if err != nil {
			g.Log(logger).Error(ctx, "service Member Edit insert mysql error:", err.Error())
			err = errors.New("操作失败[004]")
			return
		}
	}

	// 更新角色
	if err = dao.SysMemberRole.UpdateMemberRoleByIds(ctx, in.ID, in.RuleIds, tx); err != nil {
		g.Log(logger).Error(ctx, "service Member Edit update UpdateMemberRoleByIds error:", err.Error())
		err = errors.New("操作失败[006]")
		return
	}

	return
}

// UpdatePassWd 修改密码
func (s *sMember) UpdatePassWd(ctx context.Context, in *model.MemberUpdatePassWdInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-member-UpdatePassWd")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysMember.IsUniqueMember(ctx, in.ID)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member UpdatePassWd update rule error:", err.Error())
		err = errors.New("操作失败[001]")
		return err
	}

	if !ok {
		g.Log(logger).Error(ctx, "service Member UpdatePassWd member does not exist")
		err = errors.New("用户不存在")
		return err
	}
	passwd, err := service.Auth().EncryptPass(in.Passwd)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member UpdatePassWd EncryptPass error:", err.Error())
		err = errors.New("操作失败[002]")
		return err
	}
	if _, err = dao.SysMember.Ctx(ctx).Where("id", in.ID).Update(g.Map{
		"password": string(passwd),
	}); err != nil {
		g.Log(logger).Error(ctx, "service Member UpdatePassWd error:", err.Error())
		err = errors.New("操作失败[003]")
		return
	}

	return
}

// Del 删除管理员
func (s *sMember) Del(ctx context.Context, in *model.MemberBaseInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-member-Del")
	defer span.End()
	var logger = gconv.String(ctx.Value("logger"))

	ok, err := dao.SysMember.IsUniqueMember(ctx, in.ID)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member IsUniqueMember error:", err.Error())
		err = errors.New("操作失败[001]")
		return err
	}

	if !ok {
		g.Log(logger).Error(ctx, "service Member UpdatePassWd member does not exist")
		err = errors.New("用户不存在")
		return err
	}

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member Edit Transaction error:", err.Error())
		err = errors.New("操作失败[002]")
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	if _, err = dao.SysMember.Ctx(ctx).Delete("id", in.ID); err != nil {
		g.Log(logger).Error(ctx, "service Member Del error:", err.Error())
		err = errors.New("操作失败[003]")
		return
	}

	if _, err = dao.SysMemberRole.Ctx(ctx).Delete("member_id", in.ID); err != nil {
		g.Log(logger).Error(ctx, "service MemberRole Del error:", err.Error())
		err = errors.New("操作失败[004]")
		return
	}

	return
}

// ResetPassWd 重置密码
func (s *sMember) ResetPassWd(ctx context.Context, in *model.MemberResetPassWdInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-member-ResetPassWd")
	defer span.End()

	var (
		key    = cache.RedisCache().PhoneResetPassWdCode()
		logger = gconv.String(ctx.Value("logger"))
	)

	if in.Type == 20 {
		key = cache.RedisCache().EmailResetPassWdCode()
	}

	conn := cache.RedisCache().DefaultConnection()
	v, err := g.Redis(conn).Do(ctx, "GET", key+in.Account)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member ResetPassWd redis error:", err.Error())
		err = errors.New("重置密码错误(001)")
		return
	}

	if v.IsNil() || v.IsEmpty() || v.String() != in.VerifyCode {
		g.Log(logger).Error(ctx, "service Member ResetPassWd VerifyCode error")
		err = errors.New("验证码错误")
		return
	}

	m := dao.SysMember.Ctx(ctx)
	if in.Type == 10 {
		m = m.Where("mobile", in.Account)
	}

	if in.Type == 20 {
		m = m.Where("email", in.Account)
	}

	passwd, err := service.Auth().EncryptPass(in.Passwd)
	if err != nil {
		g.Log(logger).Error(ctx, "service Member ResetPassWd EncryptPass error:", err.Error())
		err = errors.New("重置密码错误(002)")
		return
	}

	if _, err = m.Update(g.Map{
		"password": string(passwd),
	}); err != nil {
		g.Log(logger).Error(ctx, "service Member ResetPassWd error:", err.Error())
		err = errors.New("重置密码错误(003)")
		return
	}

	return
}

// List 管理员列表
func (s *sMember) List(ctx context.Context, in *model.MemberBaseInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-member-List")
	defer span.End()
	return
}

// GetInfo 管理员信息
func (s *sMember) GetInfo(ctx context.Context, in *model.MemberBaseInput) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-member-GetInfo")
	defer span.End()
	return
}
