package cache

import (
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	// DefaultConnection 默认连接
	DefaultConnection = "default"

	// MemberLoginCode 登陆短信key
	MemberLoginCode = "member:login:code:"

	// MemberUserTokenArr 后台用户TokenArr
	MemberUserTokenArr = "member:user:token:arr"

	// MemberToken 后台用户Token
	MemberToken = "member:user:token:"

	// MemberUserTokenExpire token有效期 刷新有效期
	MemberUserTokenExpire = 259200

	// JwtToken .
	JwtToken = "jwtToken:"

	// JwtUserBind 用户身份绑定
	JwtUserBind = ""

	// PhoneResetPassWdCode 手机忘记密码短信Key
	PhoneResetPassWdCode = "phone:reset:passwd:code:"

	// EmailResetPassWdCode 邮箱忘记密码短信Key
	EmailResetPassWdCode = "email:reset:passwd:code:"
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

// MemberLoginCode 登陆短信key
func (u *redisUtil) MemberLoginCode() string {
	return MemberLoginCode
}

// MemberUserTokenArr 后台用户TokenArr
func (u *redisUtil) MemberUserTokenArr() string {
	return MemberUserTokenArr
}

// MemberToken 后台用户Token
func (u *redisUtil) MemberToken(memberNo int64) string {
	return MemberToken + gconv.String(memberNo)
}

// MemberUserTokenExpire .
func (u *redisUtil) MemberUserTokenExpire() uint {
	return MemberUserTokenExpire
}

// JwtToken .
func (u *redisUtil) JwtToken() string {
	return JwtToken
}

// JwtUserBind .
func (u *redisUtil) JwtUserBind() string {
	return JwtUserBind
}

// PhoneResetPassWdCode 手机忘记密码短信key
func (u *redisUtil) PhoneResetPassWdCode() string {
	return PhoneResetPassWdCode
}

// EmailResetPassWdCode 邮箱忘记密码短信key
func (u *redisUtil) EmailResetPassWdCode() string {
	return EmailResetPassWdCode
}
