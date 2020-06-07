package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type PlanStatusType string

const (
	Draft      PlanStatusType = "DRAFT"
	Review                    = "REVIEW"
	Schedules                 = "SCHEDULED"
	InProgress                = "IN_PROGRESS"
	Aborted                   = "ABORTED"
	Completed                 = "COMPLETED"
)

// Plan model
type Plan struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	StartTime *time.Time     `json:"start_dt"`
	EndTime   *time.Time     `json:"end_dt"`
	Operator  string         `json:"operator"`
	Status    PlanStatusType `json:"status"`
	Overview  string         `json:"overview" gorm:"default:''"`
}

func (plan *Plan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
