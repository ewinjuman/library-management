package author

import (
	"context"
	Error "github.com/ewinjuman/go-lib/error"
	pGrpc "github.com/ewinjuman/go-lib/grpc"
	Session "github.com/ewinjuman/go-lib/session"
	"google.golang.org/grpc/connectivity"
	"os"
)

type (
	AuthorGrpcService interface {
		GetAuthor(id int) (*AuthorResponse, error)
	}

	authorGrpc struct {
		// client to GRPC service
		session *Session.Session
	}
)

type authorGrpcConn struct {
	// client to GRPC service
	authorClient  AuthorServiceClient
	rpcConnection *pGrpc.RpcConnection
}

var rpcConnection *pGrpc.RpcConnection

func getUserClient() (*authorGrpcConn, error) {
	if rpcConnection != nil && rpcConnection.Connection != nil && (rpcConnection.Connection.GetState() == connectivity.Connecting || rpcConnection.Connection.GetState() == connectivity.Ready) {
		ctx := &authorGrpcConn{
			authorClient:  NewAuthorServiceClient(rpcConnection.Connection),
			rpcConnection: rpcConnection,
		}
		return ctx, nil

	}
	conf := pGrpc.Options{
		Address: os.Getenv("AUTHOR_GRPC"),
		Timeout: 20,
	}
	userConn, err := pGrpc.New(conf)
	if err != nil {
		return nil, err
	}
	rpcConnection = userConn
	ctx := &authorGrpcConn{
		authorClient:  NewAuthorServiceClient(rpcConnection.Connection),
		rpcConnection: rpcConnection,
	}
	return ctx, nil
}

// NewServerContext constructor for server context
func NewAuthorGrpc(session *Session.Session) AuthorGrpcService {
	return &authorGrpc{session: session}
}

func (s *authorGrpc) GetAuthor(id int) (*AuthorResponse, error) {
	conn, err := getUserClient()
	if err != nil {
		return &AuthorResponse{}, err
	}
	clientCtx, cancel := conn.rpcConnection.CreateContext(context.Background(), s.session)
	defer cancel()
	request := &GetAuthorRequest{Id: uint32(id)}
	result, err := conn.authorClient.GetAuthor(clientCtx, request)
	if err != nil {
		return &AuthorResponse{}, Error.ParseError(err)
	}

	return result, nil
}
