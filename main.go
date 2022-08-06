package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/a330202207/psychology-healthy-api/internal/cmd"
	_ "github.com/a330202207/psychology-healthy-api/internal/logic"
	"github.com/a330202207/psychology-healthy-api/utility/env"
	"github.com/a330202207/psychology-healthy-api/utility/tracing"
)

func main() {
	var (
		ctx         = gctx.New()
		appEnv, err = env.New(ctx)
	)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	_, err = tracing.InitJaeger("tracing-exam-treasured-api-admin", appEnv.JaegerEndpoint(ctx), appEnv.Version(ctx), appEnv.Environment(ctx), appEnv.HostIP(ctx))
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	cmd.Main.Run(gctx.New())
}
