// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/11 17:26
// @Package controller

package controller

import (
	"context"

	v1 "github.com/a330202207/psychology-healthy-api/api/v1"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

var Admin = cAdmin{}

type cAdmin struct {
}

func (c *cAdmin) Add(ctx context.Context, input *v1.AdminAddReq) (res *v1.AdminAddRes, err error) {
	service.Admin().Add(ctx, &model.AdminAddInput{
		Account: input.Account,
		Passwd:  input.Passwd,
		Type:    input.Type,
		Status:  input.Status,
	})
	return
}
