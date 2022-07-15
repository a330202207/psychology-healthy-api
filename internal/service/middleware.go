// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 10:41
// @Package service

package service

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/a330202207/psychology-healthy-api/internal/consts"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/utility/cache"
	"github.com/a330202207/psychology-healthy-api/utility/helper"
	"github.com/a330202207/psychology-healthy-api/utility/jwt"
)

type sMiddleware struct{}

var insMiddleware = sMiddleware{}

func Middleware() *sMiddleware {
	return &insMiddleware
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	var (
		ctx           = r.Context()
		user          = new(model.ContextUser)
		authorization = jwt.Jwt().GetAuthorization(r)
		jwtToken      = cache.RedisCache().JwtToken() + gmd5.MustEncryptString(authorization)
		jwtSign, _    = g.Cfg().Get(ctx, "jwt.sign", "hotgo")
	)

	// 替换模块前缀
	routerPrefix, _ := g.Cfg().Get(ctx, "router.admin.prefix", "/api.v1/admin")
	path := gstr.Replace(r.URL.Path, routerPrefix.String(), "", 1)

	// 不需要验证登陆路由
	if helper.Helper().IsExceptLogin(ctx, path) {
		r.Middleware.Next()
		return
	}

	if authorization == "" {
		g.Log().Error(ctx, "Auth authorization is empty")
	}

	if _, err := jwt.Jwt().ParseToken(authorization, jwtSign.Bytes()); err != nil {
		g.Log().Error(ctx, "Auth authorization is empty")
	}

	conn := cache.RedisCache().DefaultConnection()

	v, err := g.Redis(conn).Do(ctx, "get", jwtToken)
	if err != nil {

	}

	if v.String() != jwtToken {

	}

	// 是否开启多端登录
	if multiPort, _ := g.Cfg().Get(ctx, "jwt.multiPort", true); !multiPort.Bool() {
		key := cache.RedisCache().JwtUserBind() + consts.AppAdmin + gconv.String(user.ID)
		originJwtToken, err := g.Redis(conn).Do(ctx, "get", key)
		if err != nil {

		}

		if v.IsNil() || v.IsEmpty() {

		}

		if originJwtToken.String() != jwtToken {

		}
	}
	customCtx := &model.Context{}
	Context().Init(r, customCtx)
	if user != nil {
		customCtx.User = &model.ContextUser{
			ID:         user.ID,
			Username:   user.Username,
			Nickname:   user.Nickname,
			Avatar:     user.Avatar,
			Email:      user.Email,
			Mobile:     user.Mobile,
			VisitCount: user.VisitCount,
			LastTime:   user.LastTime,
			LastIp:     user.LastIp,
			Exp:        user.Exp,
			Expires:    user.Expires,
			App:        user.App,
		}
	}

	r.Middleware.Next()
}
