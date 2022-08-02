// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 13:44
// @Package v1

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CaptchaReq .
type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" summary:"获取验证码" dc:"注意直接返回的是图片二进制内容" tags:"公共" `
}

// CaptchaRes .
type CaptchaRes struct {
	g.Meta `mime:"application/json" dc:"验证码二进制内容" `
	Key    string `json:"key"`
	Img    string `json:"img"`
}
