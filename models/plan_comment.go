package models

import (
	"github.com/google/uuid"
)

// PlanComment model
type PlanComment struct {
	BaseModel
	Author string    `json:"author"`
	PlanID uuid.UUID `json:"planId"`
	Body   string    `json:"body"`
}
