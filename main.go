package main

import (
	_ "gf-web/internal/packed"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-web/internal/cmd"
)

var (
	Main = &gcmd.Command{
		Name: "main",
		Brief: "start http server",
		Description: "this is the command entry for starting your process",
	}

	ctx = gctx.New()
)

func main() {
	//err := Main.AddCommand(cmd.Http, cmd.Queue)
	//cmd.Main.Run(gctx.New())
	cmd.Queue.Run(gctx.New())

	//if err != nil {
	//	panic(err)
	//}
	//Main.Run(ctx)
}
