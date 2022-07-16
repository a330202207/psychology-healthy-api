// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/8 08:44
// @Package auth

package auth

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"

	"github.com/a330202207/psychology-healthy-api/internal/consts"
	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/model/entity"
	"github.com/a330202207/psychology-healthy-api/internal/service"
	"github.com/a330202207/psychology-healthy-api/utility/cache"
	"github.com/a330202207/psychology-healthy-api/utility/helper"
	"github.com/a330202207/psychology-healthy-api/utility/jwt"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

// Authorization .
func (s *sAuth) Authorization(ctx context.Context, in *model.AuthInput) (*model.AuthOutput, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-auth-Authorization")
	defer span.End()

	var (
		logger = gconv.String(ctx.Value("logger"))
		out    *model.AuthOutput
		err    error
	)

	if in.AuthType == "code" {
		if err = s.getAuthCode(ctx, in); err != nil {
			g.Log(logger).Error(ctx, "service Authorization getAuthCode error:", err.Error())
			return nil, err
		}
	} else {
		if !service.Captcha().VerifyCode(ctx, in.VerifyKey, in.VerifyCode) {
			g.Log(logger).Error(ctx, "service Authorization VerifyCode error:", err.Error())
			err = gerror.New("验证码输入错误")
			return nil, err
		}
	}

	if in.Device != consts.AppAdmin && in.Device != consts.AppApi {
		in.Device = consts.AppAdmin
	}

	user, err := s.getMemberInfo(ctx, in)
	if err != nil {
		g.Log(logger).Error(ctx, "service Authorization getMemberInfo error:", err.Error())
		return nil, err
	}

	// 生成token
	token, err := jwt.Jwt().GenerateToken(ctx, user, false)
	if err != nil {
		g.Log(logger).Error(ctx, "service Authorization buildToken error:", err.Error())
		return nil, err
	}
	authKey := gmd5.MustEncryptString(token)
	if err = s.updateLoginInfo(ctx, authKey); err != nil {
		g.Log(logger).Error(ctx, "service Authorization updateLoginInfo error:", err.Error())
		return nil, err
	}

	out.Token = token

	return out, err
}

// getAuthCode 获取登陆短信码
func (s *sAuth) getAuthCode(ctx context.Context, in *model.AuthInput) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-auth-getAuthCode")
	defer span.End()

	conn, err := g.Redis(cache.RedisCache().DefaultConnection()).Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	v, err := conn.Do(ctx, "GET", cache.RedisCache().AdminLoginCode()+in.Account)
	if err != nil {
		return errors.New("登陆失败(001)")
	}

	if v.IsNil() || v.IsEmpty() || v.String() != in.Passwd {
		return errors.New("登陆失败(002)")
	}
	return nil
}

// getMemberInfo 获取用户登陆信息
func (s *sAuth) getMemberInfo(ctx context.Context, in *model.AuthInput) (user *model.ContextUser, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-auth-getMemberInfo")
	defer span.End()

	var member *entity.SysMember
	if err = dao.SysMember.Ctx(ctx).Scan(&member, "account = ?", in.Account); err != nil {
		err = errors.New("登陆失败(003)")
		return
	}

	if member == nil {
		err = errors.New("登陆失败(004)")
		return
	}

	if member.Status != 20 {
		err = errors.New("帐号状态异常，请联系管理员(001)")
		return
	}

	if in.AuthType == "account" && !s.compareHashAndPassword(in.Passwd, member.Password) {
		err = errors.New("密码错误")
		return
	}

	jwtExpires, err := g.Cfg().Get(ctx, "jwt.expires", 1)
	if err != nil {
		err = gerror.New(err.Error())
		return
	}
	// 有效期
	expires := jwtExpires.Int64()

	// 过期时间戳
	exp := gconv.Int64(gtime.Timestamp()) + expires

	user = &model.ContextUser{
		ID:       member.Id,
		Username: member.Username,
		Nickname: member.NickName,
		Avatar:   member.Avatar,
		Email:    member.Email,
		Mobile:   member.Mobile,
		Type:     gconv.Uint(member.Type),
		Exp:      exp,
		Expires:  expires,
		App:      in.Device,
	}

	return
}

// compareHashAndPassword 校验密码
func (s *sAuth) compareHashAndPassword(inputPass, authPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(authPass), []byte(inputPass)); err != nil {
		return false
	}
	return true
}

// EncryptPass .加密处理
func (s *sAuth) EncryptPass(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// updateLoginInfo 更新登陆信息
func (s *sAuth) updateLoginInfo(ctx context.Context, authKey string) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-auth-updateLoginInfo")
	defer span.End()

	if _, err := dao.SysMember.Ctx(ctx).Where("id = ").Update(g.Map{
		"visit_count": &gdb.Counter{
			Field: "visit_count",
			Value: 1,
		},
		"auth_key": authKey,
		"last_at":  gtime.Now(),
		"last_ip":  helper.Helper().GetClientIp(ctx),
	}); err != nil {
		err := gerror.New(err.Error())
		return err
	}

	return nil
}
