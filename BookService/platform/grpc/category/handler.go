package category

import (
	"context"
	Error "github.com/ewinjuman/go-lib/error"
	pGrpc "github.com/ewinjuman/go-lib/grpc"
	Session "github.com/ewinjuman/go-lib/session"
	"google.golang.org/grpc/connectivity"
	"os"
)

type (
	CategoryGrpcService interface {
		GetCategory(id int) (*CategoryResponse, error)
	}

	categoryGrpc struct {
		// client to GRPC service
		session *Session.Session
	}
)

type categoryGrpcConn struct {
	// client to GRPC service
	categoryClient CategoryServiceClient
	rpcConnection  *pGrpc.RpcConnection
}

var rpcConnection *pGrpc.RpcConnection

func getUserClient() (*categoryGrpcConn, error) {
	if rpcConnection != nil && rpcConnection.Connection != nil && (rpcConnection.Connection.GetState() == connectivity.Connecting || rpcConnection.Connection.GetState() == connectivity.Ready) {
		ctx := &categoryGrpcConn{
			categoryClient: NewCategoryServiceClient(rpcConnection.Connection),
			rpcConnection:  rpcConnection,
		}
		return ctx, nil

	}
	conf := pGrpc.Options{
		Address: os.Getenv("CATEGORY_GRPC"),
		Timeout: 20,
	}
	userConn, err := pGrpc.New(conf)
	if err != nil {
		return nil, err
	}
	rpcConnection = userConn
	ctx := &categoryGrpcConn{
		categoryClient: NewCategoryServiceClient(rpcConnection.Connection),
		rpcConnection:  rpcConnection,
	}
	return ctx, nil
}

// NewServerContext constructor for server context
func NewCategoryGrpc(session *Session.Session) CategoryGrpcService {
	return &categoryGrpc{session: session}
}

func (s *categoryGrpc) GetCategory(id int) (*CategoryResponse, error) {
	conn, err := getUserClient()
	if err != nil {
		return &CategoryResponse{}, err
	}
	clientCtx, cancel := conn.rpcConnection.CreateContext(context.Background(), s.session)
	defer cancel()
	request := &GetCategoryRequest{Id: uint32(id)}
	result, err := conn.categoryClient.GetCategory(clientCtx, request)
	if err != nil {
		return &CategoryResponse{}, Error.ParseError(err)
	}

	return result, nil
}
