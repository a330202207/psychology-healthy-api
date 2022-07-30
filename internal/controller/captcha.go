// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 13:44
// @Package controller

package controller

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/a330202207/psychology-healthy-api/api/v1"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

var Captcha = cCaptcha{}

type cCaptcha struct{}

// Get 获取验证码
func (c *cCaptcha) Get(ctx context.Context, input *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-api-admin-captcha-get")
	defer span.End()

	idKeyC, base64s, err := service.Captcha().GetVerifyImgString(ctx)
	if err != nil {
		return
	}

	res = &v1.CaptchaRes{
		Key: idKeyC,
		Img: base64s,
	}

	return
}
