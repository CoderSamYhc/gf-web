package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfoReq struct {
	g.Meta `path:"/info" tags:"User" method:"get" summary:"user info"`
}

type UserInfoRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
