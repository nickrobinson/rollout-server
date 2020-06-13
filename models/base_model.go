package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key`
	CreatedAt time.Time `json:"createdAt"`
}

func (model *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
