// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 10:53
// @Package jwt

package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"

	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/utility/cache"
)

var uJWT = uJwt{}

func Jwt() *uJwt {
	return &uJWT
}

type uJwt struct{}

// GenerateToken 生成token
func (u *uJwt) GenerateToken(ctx context.Context, auth *model.ContextUser, isRefresh bool) (string, error) {
	var (
		jwtVersion, _ = g.Cfg().Get(ctx, "jwt.version", "1.0")
		jwtSign, _    = g.Cfg().Get(ctx, "jwt.sign", "psychology")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          auth.ID,
		"username":    auth.Username,
		"nick_name":   auth.Nickname,
		"avatar":      auth.Avatar,
		"email":       auth.Email,
		"mobile":      auth.Mobile,
		"last_time":   auth.LastTime,
		"last_ip":     auth.LastIp,
		"exp":         auth.Exp,
		"type":        auth.Type,
		"expires":     auth.Expires,
		"app":         auth.App,
		"visit_count": auth.VisitCount,
		"is_refresh":  isRefresh,
		"jwt_version": jwtVersion.String(),
	})

	tokenString, err := token.SignedString(jwtSign.Bytes())
	if err != nil {
		err = gerror.New(err.Error())
		return "", err
	}

	tokenStringMd5 := gmd5.MustEncryptString(tokenString)

	key := cache.RedisCache().JwtToken() + tokenStringMd5

	// 将有效期转为持续时间，单位：秒
	expires, _ := time.ParseDuration(fmt.Sprintf("+%vs", auth.Expires))
	conn := cache.RedisCache().DefaultConnection()
	if _, err = g.Redis(conn).Do(ctx, "set", key, tokenString, expires); err != nil {
		return "", err
	}

	bindStr := cache.RedisCache().JwtUserBind() + auth.App + gconv.String(auth.ID)
	g.Redis(conn).Do(ctx, "set", bindStr, key, expires)

	return tokenString, err
}

// ParseToken 解析token
func (u *uJwt) ParseToken(token string, secret []byte) (jwt.MapClaims, error) {
	if token == "" {
		err := gerror.New("token 为空")
		return nil, err
	}

	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetAuthorization 获取authorization
func (u *uJwt) GetAuthorization(r *ghttp.Request) string {
	var authorization = r.Header.Get("Authorization")
	if authorization == "" {
		return r.Get("authorization").String()
	}

	return gstr.Replace(authorization, "Bearer ", "")
}
