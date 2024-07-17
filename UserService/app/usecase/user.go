package usecase

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"library-management/UserService/app/domain/entities"
	"library-management/UserService/app/domain/queries"
	"library-management/UserService/app/models"
	"library-management/UserService/pkg/repository"
	"library-management/UserService/pkg/utils"
)

type (
	UserUsecaseService interface {
		Register(request *models.Register) (interface{}, error)
		Login(request *models.Credentials) (interface{}, error)
		Get() (interface{}, error)
		Me() (response interface{}, err error)
	}

	userUsecase struct {
		session   *session.Session
		userQuery queries.UserQueriesService
	}
)

func NewUserUsecase(session *session.Session) (item UserUsecaseService) {
	return &userUsecase{
		session:   session,
		userQuery: queries.NewUserQueries(session),
	}
}

func (h *userUsecase) Register(request *models.Register) (response interface{}, err error) {
	user := &entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: utils.GeneratePassword(request.Password),
	}
	response, err = h.userQuery.Create(user)
	return
}

func (h *userUsecase) Login(request *models.Credentials) (response interface{}, err error) {
	idclaim := fiberUtils.UUID()

	user, err := h.userQuery.GetByUsername(request.Username)
	if err != nil {
		return
	}

	compareUserPassword := utils.ComparePasswords(user.Password, request.Password)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
		return nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, "Username/Password yang diinput salah")
	}

	meta := utils.TokenMetadata{
		Id:       idclaim,
		UserID:   int(user.ID),
		Username: user.Username,
	}
	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(idclaim, meta)
	if err != nil {
		// Return status 500 and token generation error.
		return nil, Error.New(fiber.StatusInternalServerError, repository.FailedStatus, err.Error())
	}

	response = map[string]interface{}{
		"tokens": tokens,
	}
	r, errd := utils.JWTInterceptor(tokens.Access)
	if errd != nil {
		println(errd.Error())
	} else {
		println(r.UserID)
	}
	return
}

func (h *userUsecase) Get() (response interface{}, err error) {
	users, err := h.userQuery.Get()
	if err != nil {
		return
	}
	response = users
	return
}

func (h *userUsecase) Me() (response interface{}, err error) {
	f, e := h.session.Get("meta")
	if e != nil {
		println(e.Error())
	} else {
		if meta, ok := f.(*utils.TokenMetadata); ok {
			users, errg := h.userQuery.GetByID(meta.UserID)
			if errg != nil {
				err = errg
				return
			}
			response = users
			return
		}
	}
	return
}
