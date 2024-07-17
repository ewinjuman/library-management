package queries

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/fiber/v2"
	"library-management/BookService/app/domain/entities"
	"library-management/BookService/pkg/repository"
	"library-management/BookService/platform/database"
)

type (
	BorrowQueriesService interface {
		Create(req *entities.Borrow) (*entities.Borrow, error)
		Get() (borrow []entities.Borrow, err error)
		GetByUserIDAndBookID(userID uint, bookID uint, status string) (result entities.Borrow, err error)
		UpdateBorrow(borrow *entities.Borrow) (err error)
		GetWithStatus(status ...string) (result []entities.Borrow, err error)
		CountTopBorrowBook() (bookIDs []int, err error)
	}

	borrowQueries struct {
		session *session.Session
	}
)

func NewBorrowQueries(session *session.Session) (rep BorrowQueriesService) {
	return &borrowQueries{session: session}
}

func (r *borrowQueries) Create(req *entities.Borrow) (result *entities.Borrow, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newResult := new(entities.Borrow)
	err = db.Omit("updated_at").Create(req).Scan(newResult).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return newResult, nil
}

func (r *borrowQueries) Get() (result []entities.Borrow, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Find(&result).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (r *borrowQueries) GetWithStatus(status ...string) (result []entities.Borrow, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Where("status in ?", status).Find(&result).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (r *borrowQueries) GetById(id uint) (result entities.Borrow, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	result.ID = id
	err = db.First(&result).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (r *borrowQueries) GetByUserIDAndBookID(userID uint, bookID uint, status string) (result entities.Borrow, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Where("user_id = ? and book_id = ? and status = ?", userID, bookID, status).First(&result).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}

func (r *borrowQueries) UpdateBorrow(borrow *entities.Borrow) (err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	if borrow.ID == 0 {
		err = Error.New(fiber.StatusBadRequest, repository.FailedStatus, "No Data to Update")
		return
	}
	err = db.Model(&borrow).Updates(borrow).Error
	return
}

type BorrowedBook struct {
	BookID     int `gorm:"column:book_id"`
	CountValue int `gorm:"column:count_value"`
}

func (r *borrowQueries) CountTopBorrowBook() (bookIDs []int, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}

	var results []BorrowedBook
	db.Raw(`
        SELECT book_id, COUNT(book_id) AS count_value FROM borrows GROUP BY book_id ORDER BY count_value DESC LIMIT 5;`).Scan(&results)

	for _, result := range results {
		bookIDs = append(bookIDs, result.BookID)
	}
	return
}
