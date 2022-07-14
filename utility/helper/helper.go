// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/8 17:18
// @Package helper

package helper

import (
	"context"
	"math/rand"
	"time"
	"unsafe"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	helperUtilSnowflake      = "helper.util.snowflake"
	helperUtilIDGenSnowflake = "helper.util.idgen.snowflake"

	// userAgent .
	httpHeaderUserAgent = `Mozilla/5.0 (lanren; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36`
	telemetrySDKName    = "opentelemetry"
)

var uHelper = utilHelper{}

// Helper .
func Helper() *utilHelper {
	return &uHelper
}

type utilHelper struct {
}

var src = rand.NewSource(time.Now().UnixNano())

// InitRandStr RandStringBytesMaskImprSrcUnsafe
func (u *utilHelper) InitRandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// GetClientIp 获取客户端IP
func (u *utilHelper) GetClientIp(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetClientIp()
}

// IsExists 判断字符或切片字符是否存在指定字符
func (u *utilHelper) IsExists(elems interface{}, search string) bool {
	switch elems.(type) {
	case []string:
		elem := gconv.Strings(elems)
		for i := 0; i < len(elem); i++ {
			if gconv.String(elem[i]) == search {
				return true
			}
		}
	default:
		return gconv.String(elems) == search
	}

	return false
}

// IsExceptAuth 是否不需要验证权限的路由地址
func (u *utilHelper) IsExceptAuth(ctx context.Context, path string) bool {
	var pathList []string

	except, _ := g.Cfg().Get(ctx, "router.admin.exceptAuth")
	pathList = except.Strings()

	for i := 0; i < len(pathList); i++ {
		if u.IsExists(pathList[i], path) {
			return true
		}
	}

	return false
}

// IsExceptLogin 是否不需要登录的路由地址
func (u *utilHelper) IsExceptLogin(ctx context.Context, path string) bool {
	var pathList []string

	except, _ := g.Cfg().Get(ctx, "router.admin.exceptLogin")
	pathList = except.Strings()

	for i := 0; i < len(pathList); i++ {
		if u.IsExists(pathList[i], path) {
			return true
		}
	}

	return false
}
