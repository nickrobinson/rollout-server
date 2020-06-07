package models

import "time"

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
	ID        uint           `json:"id" gorm:"primary_key`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	StartTime *time.Time     `json:"start_dt"`
	EndTime   *time.Time     `json:"end_dt"`
	Operator  string         `json:"operator"`
	Status    PlanStatusType `json:"status"`
	Overview  string         `json:"overview" gorm:"default:''"`
}
