package models

import (
	"time"
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
	BaseModel
	Title        string         `json:"title"`
	Author       string         `json:"author"`
	StartTime    *time.Time     `json:"start"`
	EndTime      *time.Time     `json:"end"`
	Operator     string         `json:"operator"`
	Status       PlanStatusType `json:"status"`
	Overview     string         `json:"overview" gorm:"default:''"`
	RollbackPlan string         `json:"rollbackPlan" gorm:"default:''"`
	PlanComments []PlanComment  `json:"planComments"`
}
