package cache

import (
	"context"

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
)

var insRedis = redisUtil{}

type redisUtil struct {
}

// RedisCache redis cache
func RedisCache() *redisUtil {
	return &insRedis
}

// DefaultConnection 默认
func (c *redisUtil) DefaultConnection(ctx context.Context) string {
	return DefaultConnection
}

// AdminLoginCode 登陆短信key
func (c *redisUtil) AdminLoginCode(ctx context.Context) string {
	return AdminLoginCode
}

// AdminUserTokenArr 后台用户TokenArr
func (c *redisUtil) AdminUserTokenArr(ctx context.Context) string {
	return AdminUserTokenArr
}

// AdminToken 后台用户Token
func (c *redisUtil) AdminToken(ctx context.Context, adminNo uint64) string {
	return AdminToken + gconv.String(adminNo)
}

// AdminUserTokenExpire .
func (c *redisUtil) AdminUserTokenExpire(ctx context.Context) uint {
	return AdminUserTokenExpire
}
