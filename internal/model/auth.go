// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/8 09:01
// @Package model

package model

// AuthInput .
type AuthInput struct {
	Account    string // 登陆用户名
	Passwd     string // 登陆密码(短信验证码)
	AuthType   string // 登陆类型
	VerifyCode string // 验证码
	VerifyKey  string // 验证码key
}

// AuthOutput .
type AuthOutput struct {
	Token string
}
