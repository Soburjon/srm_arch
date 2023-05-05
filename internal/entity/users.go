package entity

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users"`

	ID          string  `json:"id" bun:"id,type:uuid,pk"`
	FullName    string  `json:"full_name" bun:"full_name"`
	Avatar      *string `json:"avatar" bun:"avatar,notnull"`
	Role        string  `json:"role" bun:"role"`
	PhoneNumber string  `json:"phone_number" bun:"phone_number"`
	Birthday    string  `json:"birthday" bun:"birthday"`
	Password    string  `json:"password" bun:"password"`
	Position    string  `json:"position" bun:"position"`

	BasicEntity
}
