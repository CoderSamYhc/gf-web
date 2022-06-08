package cmd

import (
	"context"
	"fmt"
	"gf-web/internal/queue"
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
			if err != nil {
				panic(err)
			}
			q := queue.NewQueue(Rds, ctx)
			defer Rds.Close(ctx)
			//p := parser.GetOpt("a")
			//g.Redis().Conn(ctx)
			//err = q.Push("test", "msg:test")
			//msg, err := q.Pop("test")
			//fmt.Println(msg, err)
			return nil
		},
	}
)
