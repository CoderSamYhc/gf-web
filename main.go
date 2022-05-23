package main

import (
	_ "gf-web/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-web/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
