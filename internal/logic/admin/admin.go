// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:28
// @Package admin

package admin

import (
	"context"

	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sAdmin struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) Add(ctx context.Context, in *model.AdminAddInput) (out *model.AdminAddOutput, err error) {
	return nil, err
}
