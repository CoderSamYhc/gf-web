package entity

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}
