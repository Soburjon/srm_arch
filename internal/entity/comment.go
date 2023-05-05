package entity

import "github.com/uptrace/bun"

type Comment struct {
	bun.BaseModel `bun:"table:comment"`

	TaskID       string `json:"task_id" ban:"task_id,notnull"`
	ProgrammerID string `json:"programmer_id" ban:"programmer_id,notnull"`
	Text         string `json:"text" ban:"text,notnull"`

	BasicEntity
}
