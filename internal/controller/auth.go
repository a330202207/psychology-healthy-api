// Package controller @Author NedRen 2022/7/6 17:33:00
package controller

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/a330202207/psychology-healthy-api/api/v1"
	"github.com/a330202207/psychology-healthy-api/internal/model"
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sAuth struct {
}

var insAuth = sAuth{}

func Auth() *sAuth {
	return &insAuth
}

// Authorization .
func (s *sAuth) Authorization(ctx context.Context, req *v1.AuthReq) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-api-admin-auth-auth")
	defer span.End()

	service.Auth().Authorization(ctx, &model.AuthInput{
		Account:    req.Account,
		Passwd:     req.Passwd,
		AuthType:   req.AuthType,
		VerifyCode: req.VerifyCode,
		VerifyKey:  req.VerifyKey,
	})

	return
}
