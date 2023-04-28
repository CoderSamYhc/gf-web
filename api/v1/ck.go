package v1

import (
	"gf-web/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CkReq struct {
	g.Meta `path:"/" tags:"ck" method:"get" summary:"You first hello api"`
}
type ShowTablesRes struct {
	*entity.ShowTables
}

type ListReq struct {
	g.Meta `path:"/list" tags:"ck" method:"get" summary:"You first hello api"`
}

type ListRes struct {
	Data []*entity.Test
}
