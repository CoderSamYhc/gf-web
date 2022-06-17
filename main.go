package main

import (
	_ "gf-web/internal/packed"
	"gf-web/internal/queue"
	"gf-web/internal/queue/user"
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
	QueueList = make(map[string]interface{})

)

func main() {


	q := queue.NewQueue(ctx)
	defer q.Repo.Client.Close(ctx)

	QueueList["userLogin"] = &user.UserLoginQueue{}

	rcLis := make(chan queue.RecoverData, len(QueueList))
	q.SetRecoverCh(rcLis)

	for _, v := range QueueList {
		err := q.RegisterQueue("demo", "demo1", v)
		if err != nil {
			panic(err)
		}
	}

	err := Main.AddCommand(cmd.Http, cmd.QueueService)
	if err != nil {
		panic(err)
	}
	Main.Run(ctx)
}
