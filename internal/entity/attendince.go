package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Attendince struct {
	bun.BaseModel `bun:"table:attendince"`

	UserID    string     `json:"user_id" ban:"user_id,notnull"`
	Type      string     `json:"type" ban:"type,notnull"`
	CreatedAt *time.Time `json:"created_at" bun:"created_at,default:CURRENT_TIMESTAMP"`
}
