package proto

import (
	"context"
	Error "github.com/ewinjuman/go-lib/error"
	Session "github.com/ewinjuman/go-lib/session"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"library-management/AuthorService/app/domain/queries"
)

type server struct {
	UnimplementedAuthorServiceServer
}

func (s *server) GetAuthor(ctx context.Context, in *GetAuthorRequest) (*AuthorResponse, error) {
	session := ctx.Value(Session.AppSession).(*Session.Session)

	author := queries.NewAuthorQueries(session)
	result, err := author.GetByID(int(in.Id))
	if err != nil {
		parseError := Error.ParseError(err)
		return nil, status.Error(codes.Code(parseError.ErrorCode), parseError.Message)
	}

	response := &AuthorResponse{
		Author: &Author{
			Id:   uint32(result.ID),
			Name: result.Name,
		},
	}

	return response, nil
}
