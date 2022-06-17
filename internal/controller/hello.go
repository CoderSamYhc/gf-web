package controller

import (
	"context"
	"gf-web/internal/queue"
	"github.com/gogf/gf/v2/frame/g"

	"gf-web/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {

	err = queue.QM.Push(ctx, &queue.QueuePayload{"demo", "demo1", "body"})

	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
