package main

import (
	_ "psychology-healthy-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"psychology-healthy-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
