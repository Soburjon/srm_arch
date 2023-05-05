package entity

import "github.com/uptrace/bun"

type Project struct {
	bun.BaseModel `bun:"table:project"`

	ID          string  `json:"id" bun:"id,type:uuid,pk"`
	Name        string  `json:"name" ban:"name,notnull"`
	Status      string  `json:"status" ban:"status,notnull,default:'new'"`
	StartDate   *string `json:"start_date" ban:"start_date"`
	EndDate     *string `json:"end_date" ban:"end_date"`
	TeamleadID  string  `json:"teamlead_id" ban:"teamlead_id,notnull"`
	Attachments *string `json:"attachments" ban:"attachments"`

	BasicEntity
}
