package entities

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
