package models

type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Register struct {
	Name     string `json:"name"  validate:"required"`
	Username string `json:"username" gorm:"unique"  validate:"required"`
	Password string `json:"password" validate:"required"`
}
