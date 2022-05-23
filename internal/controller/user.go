package controller

import (
	"context"
	v1 "gf-web/api/v1"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = cUser{}
)

type cUser struct {}

func (u *cUser) Info(ctx context.Context, params *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	g.RequestFromCtx(ctx).Response.WriteJson(struct {
		Code int
		Message string
	}{
		200,
		"ok",
	})
	return
}
