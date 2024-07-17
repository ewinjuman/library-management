package usecase

import (
	"github.com/ewinjuman/go-lib/session"
	"library-management/AuthorService/app/domain/entities"
	"library-management/AuthorService/app/domain/queries"
)

type (
	AuthorUsecaseService interface {
		Create(request *entities.Author) (interface{}, error)
		Get() (interface{}, error)
		GetAuthor(id int) (interface{}, error)
	}

	authorUsecase struct {
		session     *session.Session
		authorQuery queries.AuthorQueriesService
	}
)

func NewAuthorUsecase(session *session.Session) (item AuthorUsecaseService) {
	return &authorUsecase{
		session:     session,
		authorQuery: queries.NewAuthorQueries(session),
	}
}

func (h *authorUsecase) Create(request *entities.Author) (response interface{}, err error) {
	response, err = h.authorQuery.Create(request)
	return
}

func (h *authorUsecase) Get() (response interface{}, err error) {
	response, err = h.authorQuery.Get()
	return
}

func (h *authorUsecase) GetAuthor(id int) (response interface{}, err error) {
	response, err = h.authorQuery.GetByID(id)
	return
}
