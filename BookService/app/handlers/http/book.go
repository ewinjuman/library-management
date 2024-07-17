package http

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/gofiber/fiber/v2"
	"library-management/BookService/app/domain/entities"
	"library-management/BookService/app/models"
	"library-management/BookService/app/usecase"
	"library-management/BookService/pkg/base"
	"library-management/BookService/pkg/repository"
	"library-management/BookService/pkg/utils"
)

func Create(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	request := &entities.Book{}

	// Checking received data from JSON body.
	if err := ctx.BodyParser(request); err != nil {
		// Return status 400 and error message.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()
	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		// Return, if some fields are not valid.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}
	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.Create(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Get(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.Get()
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Borrow(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	request := &models.BorrowBookRequest{}

	// Checking received data from JSON body.
	if err := ctx.BodyParser(request); err != nil {
		// Return status 400 and error message.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()
	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		// Return, if some fields are not valid.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}
	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.Borrow(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Return(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	request := &models.ReturnBookRequest{}

	// Checking received data from JSON body.
	if err := ctx.BodyParser(request); err != nil {
		// Return status 400 and error message.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()
	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		// Return, if some fields are not valid.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}
	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.Return(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func GetBorrow(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.GetBorrow()
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func GetBook(c *fiber.Ctx) error {
	ctx := base.NewContext(c)
	id, _ := c.ParamsInt("id", 0)
	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.GetBook(id)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func GetRecommend(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.GetRecommend()
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Search(c *fiber.Ctx) error {
	ctx := base.NewContext(c)
	keyword := c.Query("keyword")
	book := usecase.NewBookUsecase(ctx.Session)
	result, err := book.Search(keyword)
	// Return status 200 OK.
	return ctx.Response(result, err)
}
