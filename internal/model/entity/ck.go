package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ShowTables struct {
	Name string `json:"name" ch:"name"`
}

type Test struct {
	Id         uint32     `ch:"id"`
	Name       string     `ch:"name"`
	CreateDate gtime.Time `ch:"create_date"`
	CreatedAt  gtime.Time `ch:"created_at"`
}
