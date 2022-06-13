package cmd

import (
	"context"
	"fmt"
	"gf-web/internal/queue"
	"gf-web/internal/queue/user"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (

	QueueList = make(map[string]interface{})

	QueueService = &gcmd.Command{
		Name:        "queue",
		Usage:       "queue",
		Brief:       "start queue server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {



			q := queue.QM
			defer q.Repo.Client.Close(ctx)

			QueueList["userLogin"] = &user.UserLoginQueue{}

			rcLis := make(chan queue.RecoverData, len(QueueList))
			q.SetRecoverCh(rcLis)

			for _, v := range QueueList {
				err = q.RegisterQueue("demo", "demo1", v)
				if err != nil {
					panic(err)
				}
			}

			//err = Queue.Push(&queue.QueuePayload{"demo", "demo1", "body"})
			//
			//if err != nil {
			//	panic(err)
			//}
			for {
				select {
				case rcData := <-q.RecoverCh:
					err := q.Pop(&rcData)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			return nil
		},
	}
)
