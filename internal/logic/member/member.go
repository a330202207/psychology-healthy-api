// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:28
// @Package admin

package member

import (
	"context"

	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sMember struct {
}

func init() {
	service.RegisterMember(New())
}

func New() *sMember {
	return &sMember{}
}

// Add .
func (s *sMember) Add(ctx context.Context, in *model.MemberAddInput) (out *model.MemberAddOutput, err error) {
	return nil, err
}
