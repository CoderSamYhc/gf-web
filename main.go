package main

import (
	_ "gf-web/internal/packed"
	"gf-web/internal/queue"
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
	err := Main.AddCommand(cmd.Http, cmd.QueueService)
	if err != nil {
		panic(err)
	}
	q := queue.NewQueue(ctx)
	defer q.Repo.Client.Close(ctx)
	Main.Run(ctx)
}
