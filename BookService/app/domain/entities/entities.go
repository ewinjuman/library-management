package entities

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Title      string         `json:"title" gorm:"unique" validate:"required"`
	AuthorID   uint           `json:"author_id" validate:"required,min=1"`
	CategoryID uint           `json:"category_id" validate:"required,min=1"`
	Stock      int64          `json:"stock" validate:"required,min=1"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type Borrow struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	BookID    uint           `json:"book_id"`
	Status    string         `json:"status"` // e.g., "borrowed", "returned"
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
