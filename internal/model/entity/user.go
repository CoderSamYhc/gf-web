package entity

import "time"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	CreatedAt int `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
