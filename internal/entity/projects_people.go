package entity

import "github.com/uptrace/bun"

type ProjectsPeople struct {
	bun.BaseModel `bun:"table:projects_people"`

	UserID    string `json:"user_id" ban:"user_id,notnull"`
	ProjectID string `json:"project_id" ban:"project_id,notnull,notnull"`
	Position  string `json:"position" ban:"position,notnull"`
}
