package queries

import (
	"github.com/ewinjuman/go-lib/session"
	"github.com/lib/pq"
	"library-management/UserService/app/domain/entities"
	"library-management/UserService/pkg/repository"
	"library-management/UserService/platform/database"
)

type (
	UserQueriesService interface {
		Create(req *entities.User) (*entities.User, error)
		Get() (books []entities.User, err error)
		GetByUsername(username string) (result entities.User, err error)
		GetByID(id int) (result entities.User, err error)
	}

	userQueries struct {
		session *session.Session
	}
)

func NewUserQueries(session *session.Session) (rep UserQueriesService) {
	return &userQueries{session: session}
}

func (r *userQueries) Create(req *entities.User) (result *entities.User, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newResult := new(entities.User)
	err = db.Omit("updated_at").Create(req).Scan(newResult).Error
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			println(pqErr.Code)
		}
		err = repository.HandleSqlError(err)
		return
	}
	return newResult, nil
}

func (r *userQueries) Get() (result []entities.User, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Find(&result).Error
	if err != nil {
		err = repository.HandleSqlError(err)
		return
	}
	return
}

func (r *userQueries) GetByUsername(username string) (result entities.User, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	err = db.Where("username = ?", username).First(&result).Error
	if err != nil {
		err = repository.HandleSqlError(err)
		return
	}
	return
}

func (r *userQueries) GetByID(id int) (result entities.User, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	result.ID = uint(id)
	err = db.Find(&result).Error
	if err != nil {
		err = repository.HandleSqlError(err)
		return
	}
	return
}
