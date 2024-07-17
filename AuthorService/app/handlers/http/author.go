package http

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/gofiber/fiber/v2"
	"library-management/AuthorService/app/domain/entities"
	"library-management/AuthorService/app/usecase"
	"library-management/AuthorService/pkg/base"
	"library-management/AuthorService/pkg/repository"
	"library-management/AuthorService/pkg/utils"
)

func Create(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	request := &entities.Author{}

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
	book := usecase.NewAuthorUsecase(ctx.Session)
	result, err := book.Create(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Get(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	book := usecase.NewAuthorUsecase(ctx.Session)
	result, err := book.Get()
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func GetAuthor(c *fiber.Ctx) error {
	ctx := base.NewContext(c)
	id, _ := ctx.ParamsInt("id", 0)
	book := usecase.NewAuthorUsecase(ctx.Session)
	result, err := book.GetAuthor(id)
	// Return status 200 OK.
	return ctx.Response(result, err)
}
