package usecase

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/session"
	"library-management/BookService/app/domain/entities"
	"library-management/BookService/app/domain/queries"
	"library-management/BookService/app/models"
	"library-management/BookService/pkg/repository"
	"library-management/BookService/platform/grpc/author"
	"library-management/BookService/platform/grpc/category"
	"library-management/BookService/platform/grpc/user"
)

type (
	BookUsecaseService interface {
		Create(request *entities.Book) (interface{}, error)
		Get() (interface{}, error)
		Borrow(request *models.BorrowBookRequest) (response interface{}, err error)
		Return(request *models.ReturnBookRequest) (response interface{}, err error)
		GetBorrow() (response interface{}, err error)
		GetBook(id int) (response interface{}, err error)
		GetRecommend() (response interface{}, err error)
		Search(keyword string) (response interface{}, err error)
	}

	bookUsecase struct {
		session      *session.Session
		bookQuery    queries.BookQueriesService
		borrowQuery  queries.BorrowQueriesService
		authorGrpc   author.AuthorGrpcService
		categoryGrpc category.CategoryGrpcService
		userGrpc     user.UserGrpcService
	}
)

func NewBookUsecase(session *session.Session) (item BookUsecaseService) {
	return &bookUsecase{
		session:      session,
		bookQuery:    queries.NewBookQueries(session),
		borrowQuery:  queries.NewBorrowQueries(session),
		authorGrpc:   author.NewAuthorGrpc(session),
		categoryGrpc: category.NewCategoryGrpc(session),
		userGrpc:     user.NewUserGrpc(session),
	}
}

func (h *bookUsecase) Create(request *entities.Book) (response interface{}, err error) {
	_, err = h.authorGrpc.GetAuthor(int(request.AuthorID))
	if err != nil {
		return
	}
	_, err = h.categoryGrpc.GetCategory(int(request.CategoryID))
	if err != nil {
		return
	}
	response, err = h.bookQuery.Create(request)
	return
}

func (h *bookUsecase) Get() (response interface{}, err error) {
	response, err = h.bookQuery.Get()
	return
}

func (h *bookUsecase) Borrow(request *models.BorrowBookRequest) (response interface{}, err error) {
	_, err = h.userGrpc.GetUser(int(request.UserID))
	if err != nil {
		return
	}
	book, err := h.bookQuery.GetByID(uint(request.BookID))
	if err != nil {
		return
	}
	if book.Stock <= 0 {
		err = Error.NewError(repository.BadRequestCode, repository.FailedStatus, "Out of stock")
	}
	borrow, _ := h.borrowQuery.GetByUserIDAndBookID(uint(request.UserID), uint(request.BookID), "borrowed")
	if borrow.ID > 0 {
		err = Error.NewError(repository.BadRequestCode, repository.FailedStatus, "already borrowed")
		return
	}
	response, err = h.borrowQuery.Create(&entities.Borrow{
		UserID: uint(request.UserID),
		BookID: uint(request.BookID),
		Status: "borrowed",
	})
	if err != nil {
		return
	}
	h.bookQuery.ReduceStock(book.ID)
	return
}

func (h *bookUsecase) Return(request *models.ReturnBookRequest) (response interface{}, err error) {
	borrow, err := h.borrowQuery.GetByUserIDAndBookID(uint(request.UserID), uint(request.BookID), "borrowed")
	if err != nil {
		return
	}
	err = h.borrowQuery.UpdateBorrow(&entities.Borrow{
		ID:     borrow.ID,
		Status: "returned",
	})
	if err != nil {
		return
	}
	err = h.bookQuery.AddStock(borrow.BookID)
	return
}

func (h *bookUsecase) GetBorrow() (response interface{}, err error) {
	response, err = h.borrowQuery.Get()
	return
}

func (h *bookUsecase) GetBook(id int) (response interface{}, err error) {
	response, err = h.bookQuery.GetByID(uint(id))
	return
}

func (h *bookUsecase) GetRecommend() (response interface{}, err error) {
	results := make([]entities.Book, 0)
	f, _ := h.borrowQuery.CountTopBorrowBook()
	if len(f) > 0 {
		c, _ := h.bookQuery.GetRecommend(f)
		for _, v := range c {
			results = append(results, v)
		}
	}
	if len(f) < 5 {
		c, _ := h.bookQuery.GetRecommendRand(f, 5-len(f))
		for _, v := range c {
			results = append(results, v)
		}

	}
	response = results
	return
}

func (h *bookUsecase) Search(keyword string) (response interface{}, err error) {
	response, err = h.bookQuery.Search(keyword)
	return
}
