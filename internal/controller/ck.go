package controller

import (
	"context"
	v1 "gf-web/api/v1"
	"gf-web/internal/common"
	"gf-web/internal/consts"
	"gf-web/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type cCk struct {
}

var Ck = cCk{}

func (ck *cCk) ShowTables(ctx context.Context, req *v1.CkReq) (res *v1.ShowTablesRes, err error) {

	r := common.Response{
		Ctx: ctx,
	}

	result, err := g.DB("clickhouse").Query(ctx, "SHOW TABLES")
	if err != nil {
		r.Error(consts.ERROR, err.Error())
		return
	}
	var tables []*entity.ShowTables
	result.Structs(&tables)
	res = &v1.ShowTablesRes{
		Data: tables,
	}

	r.Success(res)
	return
}

func (ck *cCk) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {

	r := common.Response{
		Ctx: ctx,
	}

	list, err := g.DB("clickhouse").Model("test").All()

	if err != nil {
		r.Error(consts.ERROR, err.Error())
		return
	}

	var result []*entity.Test
	list.Structs(&result)
	res = &v1.ListRes{
		Data: result,
	}

	r.Success(res)
	return
}
