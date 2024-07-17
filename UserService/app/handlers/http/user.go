package http

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/gofiber/fiber/v2"
	"library-management/UserService/app/models"
	"library-management/UserService/app/usecase"
	"library-management/UserService/pkg/base"
	"library-management/UserService/pkg/repository"
	"library-management/UserService/pkg/utils"
)

func Register(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	request := &models.Register{}

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
	user := usecase.NewUserUsecase(ctx.Session)
	result, err := user.Register(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Login(c *fiber.Ctx) error {
	ctx := base.NewContext(c)
	request := &models.Credentials{}

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
	user := usecase.NewUserUsecase(ctx.Session)
	result, err := user.Login(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Get(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	user := usecase.NewUserUsecase(ctx.Session)
	result, err := user.Get()
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func Me(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	user := usecase.NewUserUsecase(ctx.Session)
	result, err := user.Me()
	// Return status 200 OK.
	return ctx.Response(result, err)
}
