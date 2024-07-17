package queries

import (
	"github.com/ewinjuman/go-lib/session"
	"library-management/CategoryService/app/domain/entities"
	"library-management/CategoryService/pkg/repository"
	"library-management/CategoryService/platform/database"
)

type (
	CategoryQueriesService interface {
		Create(req *entities.Category) (*entities.Category, error)
		Get() (result []entities.Category, err error)
		GetByID(id int) (result entities.Category, err error)
	}

	categoryQueries struct {
		session *session.Session
	}
)

func NewCategoryQueries(session *session.Session) (rep CategoryQueriesService) {
	return &categoryQueries{session: session}
}

func (r *categoryQueries) Create(req *entities.Category) (result *entities.Category, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newUser := new(entities.Category)
	err = db.Omit("updated_at").Create(req).Scan(newUser).Error
	if err != nil {
		err = repository.HandleMysqlError(err)
		return
	}
	return newUser, nil
}

func (r *categoryQueries) Get() (result []entities.Category, err error) {
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

func (r *categoryQueries) GetByID(id int) (result entities.Category, err error) {
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
