package cmd

import (
	"context"
	"fmt"
	"gf-web/internal/queue"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (


	QueueService = &gcmd.Command{
		Name:        "queue",
		Usage:       "queue",
		Brief:       "start queue server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {



			q := queue.QM
			defer q.Repo.Client.Close(ctx)


			//err = Queue.Push(&queue.QueuePayload{"demo", "demo1", "body"})
			//
			//if err != nil {
			//	panic(err)
			//}

			// 起个协程监听所有的队列数量
			go q.HandleRecover(ctx)
			for {
				select {
				case rcData := <-q.RecoverCh:
					err := q.Pop(ctx, &rcData)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			return nil
		},
	}
)
