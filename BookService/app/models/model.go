package models

type BorrowBookRequest struct {
	UserID int64 `json:"user_id"  validate:"required,min=1"`
	BookID int64 `json:"book_id"  validate:"required,min=1"`
}

type ReturnBookRequest struct {
	UserID int64 `json:"user_id"  validate:"required,min=1"`
	BookID int64 `json:"book_id"  validate:"required,min=1"`
}
