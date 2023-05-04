package controller

import (
	"context"
	v1 "gf-web/api/v1"
	"gf-web/internal/common"
	"gf-web/internal/consts"
	"gf-web/internal/model/entity"
	"gf-web/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (u *cUser) Info(ctx context.Context, params *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {

	var (
		user *entity.User
		r    = common.Response{
			Ctx: ctx,
		}
	)
	user, err = service.UserService().Info(ctx, params.Id)
	if err != nil {
		r.Error(consts.ERROR, err.Error())
		return
	}

	res = &v1.UserInfoRes{
		User: user,
	}

	r.Success(res)
	return
}
