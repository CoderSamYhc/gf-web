package entity

import "time"

type ShowTables struct {
	Name string `json:"name" ch:"name"`
}

type Test struct {
	Id uint32 `ch:"id"`
	Name string `ch:"name"`
	CreateDate time.Time `ch:"create_date"`
	CreatedAt time.Time `ch:"created_at"`
}


