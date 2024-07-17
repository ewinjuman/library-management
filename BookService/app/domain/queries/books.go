package queries

import (
	"github.com/ewinjuman/go-lib/session"
	"library-management/BookService/app/domain/entities"
	"library-management/BookService/pkg/repository"
	"library-management/BookService/platform/database"
)

type (
	BookQueriesService interface {
		Create(req *entities.Book) (*entities.Book, error)
		Get() (books []entities.Book, err error)
		GetByID(id uint) (result entities.Book, err error)
		AddStock(bookID uint) (err error)
		ReduceStock(bookID uint) (err error)
	}

	bookQueries struct {
		session *session.Session
	}
)

func NewBookQueries(session *session.Session) (rep BookQueriesService) {
	return &bookQueries{session: session}
}

func (r *bookQueries) Create(req *entities.Book) (result *entities.Book, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newResult := new(entities.Book)
	err = db.Omit("updated_at").Create(req).Scan(newResult).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return newResult, nil
}

func (r *bookQueries) Get() (result []entities.Book, err error) {
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

func (r *bookQueries) GetByID(id uint) (result entities.Book, err error) {
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

func (r *bookQueries) AddStock(bookID uint) (err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Exec("update books set stock = stock+1 where id = ?", bookID).Error

	return
}

func (r *bookQueries) ReduceStock(bookID uint) (err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Exec("update books set stock = stock-1 where id = ?", bookID).Error

	return
}
