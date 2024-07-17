package usecase

import (
	"github.com/ewinjuman/go-lib/session"
	"library-management/CategoryService/app/domain/entities"
	"library-management/CategoryService/app/domain/queries"
)

type (
	CategoryUsecaseService interface {
		Create(request *entities.Category) (interface{}, error)
		Get() (interface{}, error)
		GetCategory(id int) (response interface{}, err error)
	}

	categoryUsecase struct {
		session       *session.Session
		categoryQuery queries.CategoryQueriesService
	}
)

func NewCategoryUsecase(session *session.Session) (item CategoryUsecaseService) {
	return &categoryUsecase{
		session:       session,
		categoryQuery: queries.NewCategoryQueries(session),
	}
}

func (h *categoryUsecase) Create(request *entities.Category) (response interface{}, err error) {
	response, err = h.categoryQuery.Create(request)
	return
}

func (h *categoryUsecase) Get() (response interface{}, err error) {
	response, err = h.categoryQuery.Get()
	return
}

func (h *categoryUsecase) GetCategory(id int) (response interface{}, err error) {
	response, err = h.categoryQuery.GetByID(id)
	return
}
