package entity

import "time"

type BasicEntity struct {
	CreatedAt *time.Time `json:"created_at" bun:"created_at,notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" bun:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bun:"deleted_at"`
}
