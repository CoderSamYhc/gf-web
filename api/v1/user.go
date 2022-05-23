package v1

import (
	"gf-web/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoReq struct {
	g.Meta `path:"/info" tags:"User" method:"get" summary:"用户详情"`
	Id int `v:"required"`
}

type UserInfoRes struct {
	*entity.User
}
