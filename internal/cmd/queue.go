package cmd

import (
	"context"
	"fmt"
	"gf-web/internal/queue"
	"gf-web/internal/queue/user"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (

	Queue = &gcmd.Command{
		Name:  "queue",
		Usage: "queue",
		Brief: "start queue server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			Rds, err := g.Redis().Conn(ctx)
			fmt.Println(11)

			if err != nil {
				panic(err)
			}
			q := queue.NewQueue(Rds, ctx)
			fmt.Println(44)

			defer Rds.Close(ctx)

			rcLis := make(chan queue.RecoverData, 5)
			q.SetRecoverCh(rcLis)
			err = q.RegisterQueue("demo", "demo1", &user.UserLoginQueue{})
			fmt.Println(22)

			if err != nil {
				panic(err)
			}
			err = q.Push(&queue.QueuePayload{"demo", "demo1", "body"})
			if err != nil {
				panic(err)
			}
			fmt.Println(33)
			for  {
				select {
				case rcData := <- q.RecoverCh:
					err := q.Pop(&rcData)
					fmt.Println(err)
				}
			}

			return nil
		},
	}
)
