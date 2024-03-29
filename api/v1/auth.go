// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/7 08:55
// @Package v1

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Base .
type Base struct {
	ClientIP string `json:"clientIP"`
	UserNo   uint64 `json:"userNo"`
}

// AuthReq .
type AuthReq struct {
	g.Meta     `path:"/member/auth" method:"post" summary:"授权" tags:"授权服务"`
	Account    string `v:"required#登陆用户名不能为空" json:"account" dc:"登陆用户名"`
	Passwd     string `v:"required|length:4,20#登陆密码(手机验证码)不能为空|登陆密码(手机验证码)长度应当在4到20之间" json:"passwd" dc:"登陆密码(手机验证码)"`
	AuthType   string `v:"required|in:code,account#登陆类型不能为空" json:"authType" dc:"登陆类型"`
	VerifyCode string `v:"required-if|authType,account#验证码不能为空" json:"verifyCode" dc:"验证码"`
	VerifyKey  string `v:"required-if|authType,account" json:"verifyKey" dc:"验证码Key"`
	Device     string `json:"device" dc:"登陆设备"`
}

// AuthRes .
type AuthRes struct {
	Token string `json:"token"`
}
