package proto

import (
	"context"
	Error "github.com/ewinjuman/go-lib/error"
	Session "github.com/ewinjuman/go-lib/session"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"library-management/CategoryService/app/domain/queries"
)

type server struct {
	UnimplementedCategoryServiceServer
}

func (s *server) GetCategory(ctx context.Context, in *GetCategoryRequest) (*CategoryResponse, error) {
	session := ctx.Value(Session.AppSession).(*Session.Session)

	category := queries.NewCategoryQueries(session)
	result, err := category.GetByID(int(in.Id))
	if err != nil {
		parseError := Error.ParseError(err)
		return nil, status.Error(codes.Code(parseError.ErrorCode), parseError.Message)
	}

	response := &CategoryResponse{
		Category: &Category{
			Id:   uint32(result.ID),
			Name: result.Name,
		},
	}

	return response, nil
}
