// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 12:21
// @Package service

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/a330202207/psychology-healthy-api/internal/consts"
	"github.com/a330202207/psychology-healthy-api/internal/model"
)

type sContext struct{}

var insContext = sContext{}

func Context() *sContext {
	return &insContext
}

// Init initializes and injects custom business context object into request context.
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get .
func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser .
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
