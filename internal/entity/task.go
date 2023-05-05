package entity

import (
	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:task"`

	ID           string  `json:"id" ban:"id,type:uuid,pk"`
	ProjectID    string  `json:"project_id" ban:"project_id,notnull"`
	Title        string  `json:"title" ban:"title,notnull"`
	Description  string  `json:"description" ban:"description,notnull"`
	Status       string  `json:"status" ban:"status,notnull,default:'new'"`
	StartedAt    *string `json:"started_at" ban:"started_at"`
	FinishedAt   *string `json:"finished_at" ban:"finished_at"`
	StartAt      *string `json:"start_at" ban:"start_at"`
	FinishAt     *string `json:"finish_at" ban:"finish_at"`
	ProgrammerID string  `json:"programmer_id" ban:"programmer_id,notnull"`
	Attachments  *string `json:"attachments" ban:"attachments"`

	BasicEntity
}
