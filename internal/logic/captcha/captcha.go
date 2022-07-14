// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 13:44
// @Package captcha

package captcha

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"

	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sCaptcha struct {
}

func init() {
	service.RegisterCaptcha(New())
}

func New() *sCaptcha {
	return &sCaptcha{}
}

var (
	captchaStore  = base64Captcha.DefaultMemStore
	captchaDriver = newDriver()
)

// newDriver 验证码配置
func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

// GetVerifyImgString 获取验证码
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKey string, b64s string, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-captcha-GetVerifyImgString")
	defer span.End()
	captcha := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	idKey, b64s, err = captcha.Generate()
	return
}

// VerifyCode 校验验证码
func (s *sCaptcha) VerifyCode(ctx context.Context, key, code string) bool {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-captcha-VerifyCode")
	defer span.End()

	c := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	code = gstr.ToLower(code)
	return c.Verify(key, code, true)
}
