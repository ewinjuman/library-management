package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"  validate:"required"`
	Username  string         `json:"username" gorm:"unique"  validate:"required"`
	Password  string         `json:"-" validate:"required,min=6"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
