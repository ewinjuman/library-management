package proto

import (
	"context"
	Error "github.com/ewinjuman/go-lib/error"
	Session "github.com/ewinjuman/go-lib/session"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"library-management/UserService/app/domain/queries"
	"library-management/UserService/pkg/utils"
)

type server struct {
	UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *GetUserRequest) (*UserResponse, error) {
	session := ctx.Value(Session.AppSession).(*Session.Session)

	category := queries.NewUserQueries(session)
	result, err := category.GetByID(int(in.Id))
	if err != nil {
		parseError := Error.ParseError(err)
		return nil, status.Error(codes.Code(parseError.ErrorCode), parseError.Message)
	}

	response := &UserResponse{
		Id:       uint32(result.ID),
		Name:     result.Name,
		Username: result.Username,
	}

	return response, nil
}

func (s *server) JwtInterceptor(ctx context.Context, in *JwtInterceptorRequest) (*JwtInterceptorResponse, error) {
	//session := ctx.Value(Session.AppSession).(*Session.Session)

	result, err := utils.JWTInterceptor(in.Token)
	if err != nil {
		parseError := Error.ParseError(err)
		return nil, status.Error(codes.Code(parseError.ErrorCode), parseError.Message)
	}

	response := &JwtInterceptorResponse{
		Id:       uint32(result.UserID),
		Username: result.Username,
	}

	return response, nil
}
