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

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"

	"github.com/a330202207/psychology-healthy-api/internal/dao"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/model/entity"
	"github.com/a330202207/psychology-healthy-api/internal/service"
	"github.com/a330202207/psychology-healthy-api/utility/cache"
	"github.com/a330202207/psychology-healthy-api/utility/helper"
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
func (s *sAuth) Authorization(ctx context.Context, in *model.AuthInput) (out *model.AuthOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-api-admin-logic-auth-Authorization")
	defer span.End()
	if in.AuthType == "code" {

	}

	token := helper.Helper().InitRandStr(128)
	out.Token = token

	return
}

// checkAdminPasswd 检查用户登陆密码
func (s *sAuth) checkAdminPasswd(ctx context.Context, in model.AuthInput, token, logger string) error {
	var (
		admin *entity.Admin
		err   error
	)
	if err = dao.Admin.Ctx(ctx).Scan(&admin, "account = ?", in.Account); err != nil {
		g.Log(logger).Error(ctx, "checkAdminPasswd select-db error:", err.Error())
		return errors.New("登陆失败(003)")
	}

	if admin == nil {
		return errors.New("登陆失败(004)")
	}

	if admin.State != 200 {
		return errors.New("帐号状态异常，请联系管理员(001)")
	}

	if in.AuthType == "account" && !s.compareHashAndPassword(in.Passwd, admin.Passwd) {
		return errors.New("密码错误)")
	}

	if err = s.saveAdminToken(ctx, admin, token, logger); err != nil {
		return err
	}

	return nil
}

// compareHashAndPassword 校验密码
func (s *sAuth) compareHashAndPassword(inputPass, authPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(authPass), []byte(inputPass)); err != nil {
		return false
	}
	return true
}

// saveAdminToken 更新token
func (s *sAuth) saveAdminToken(ctx context.Context, data *entity.Admin, token, logger string) error {
	var (
		conn     = cache.RedisCache().DefaultConnection(ctx)
		tokenKey = cache.RedisCache().AdminToken(ctx, data.AdminNo)
		val, err = g.Redis(conn).Do(ctx, "GET", tokenKey)
	)

	if err != nil {
		g.Log(logger).Error(ctx, "saveAdminToken from get redis error:", err)
		return errors.New("登陆失败(005)")
	}
	if val.IsNil() || val.IsEmpty() {
		key := cache.RedisCache().AdminUserTokenArr(ctx)
		if val, err = g.Redis(conn).Do(ctx, "HSETNX", key, token, tokenKey); err != nil {
			g.Log(logger).Error(ctx, "saveAdminToken from hsetnx redis error:", err)
			return errors.New("登陆失败(006)")
		}
		TokenInfo := g.Map{
			"id":       data.Id,
			"token":    token,
			"clientIP": helper.Helper().GetClientIp(ctx),
			"time":     gtime.Now().Timestamp(),
		}
		if _, err = g.Redis(conn).Do(ctx, "SETEX", tokenKey, cache.RedisCache().AdminUserTokenExpire(ctx), TokenInfo); err != nil {
			g.Log(logger).Error(ctx, "saveAdminToken from setex redis error:", err)
			return errors.New("登陆失败(007)")
		}
	}

	return nil
}
