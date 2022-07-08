// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/8 08:44
// @Package auth

package auth

import (
	"context"

	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

// Authorization .
func (s *sAuth) Authorization(ctx context.Context, in *model.AuthInput) (out *model.AuthOutput, err error) {
	return nil, err
}
