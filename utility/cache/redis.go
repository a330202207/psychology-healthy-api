package cache

import (
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	// DefaultConnection 默认连接
	DefaultConnection = "default"

	// AdminLoginCode 登陆短信key
	AdminLoginCode = "admin:login:code:"

	// AdminUserTokenArr 后台用户TokenArr
	AdminUserTokenArr = "admin:user:token:arr"

	// AdminToken 后台用户Token
	AdminToken = "admin:user:token:"

	// AdminUserTokenExpire token有效期 刷新有效期
	AdminUserTokenExpire = 259200

	// JwtToken .
	JwtToken = "jwtToken:"

	// JwtUserBind 用户身份绑定
	JwtUserBind = ""
)

var uRedis = redisUtil{}

type redisUtil struct {
}

// RedisCache redis cache
func RedisCache() *redisUtil {
	return &uRedis
}

// DefaultConnection 默认
func (u *redisUtil) DefaultConnection() string {
	return DefaultConnection
}

// AdminLoginCode 登陆短信key
func (u *redisUtil) AdminLoginCode() string {
	return AdminLoginCode
}

// AdminUserTokenArr 后台用户TokenArr
func (u *redisUtil) AdminUserTokenArr() string {
	return AdminUserTokenArr
}

// AdminToken 后台用户Token
func (u *redisUtil) AdminToken(adminNo int64) string {
	return AdminToken + gconv.String(adminNo)
}

// AdminUserTokenExpire .
func (u *redisUtil) AdminUserTokenExpire() uint {
	return AdminUserTokenExpire
}

func (u *redisUtil) JwtToken() string {
	return JwtToken
}

func (u *redisUtil) JwtUserBind() string {
	return JwtUserBind
}
