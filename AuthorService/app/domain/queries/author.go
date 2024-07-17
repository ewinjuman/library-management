package queries

import (
	"github.com/ewinjuman/go-lib/session"
	"library-management/AuthorService/app/domain/entities"
	"library-management/AuthorService/pkg/repository"
	"library-management/AuthorService/platform/database"
)

type (
	AuthorQueriesService interface {
		Create(req *entities.Author) (*entities.Author, error)
		Get() (authors []entities.Author, err error)
		GetByID(id int) (result entities.Author, err error)
	}

	authorQueries struct {
		session *session.Session
	}
)

func NewAuthorQueries(session *session.Session) (rep AuthorQueriesService) {
	return &authorQueries{session: session}
}

func (r *authorQueries) Create(req *entities.Author) (result *entities.Author, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newUser := new(entities.Author)
	err = db.Omit("updated_at").Create(req).Scan(newUser).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return newUser, nil
}

func (r *authorQueries) Get() (result []entities.Author, err error) {
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

func (r *authorQueries) GetByID(id int) (result entities.Author, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	result.ID = uint(id)
	err = db.First(&result).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return
}
