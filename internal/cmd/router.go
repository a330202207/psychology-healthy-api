// Package cmd @Author NedRen 2022/7/6 08:51:00
package cmd

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/a330202207/psychology-healthy-api/internal/controller"
)

func BindRouter(g *ghttp.RouterGroup) {
	g.Group("/api.v1/admin", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Auth,
		)
	})
}
