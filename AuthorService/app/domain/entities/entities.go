package entities

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" validate:"required"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
