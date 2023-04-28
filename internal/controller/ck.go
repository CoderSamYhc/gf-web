package controller

import (
	"context"
	v1 "gf-web/api/v1"
	"gf-web/internal/common"
	"gf-web/internal/consts"
	"gf-web/utility"
)

type cCk struct {
}

var Ck = cCk{

}

func (ck *cCk ) ShowTables(ctx context.Context, req *v1.CkReq) (res *v1.ShowTablesRes, err error) {

	r := common.Response{
		Ctx: ctx,
	}
	client := utility.NewClickHouse(ctx)
	result, err := client.ShowTables(ctx)
	if err != nil {
		r.Error(consts.ERROR, err.Error())
		return
	}
	res = &v1.ShowTablesRes{
		ShowTables :&result,
	}

	r.Success(res)
	return
}
