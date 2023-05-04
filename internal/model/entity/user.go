package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	CreatedAt int        `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
}
