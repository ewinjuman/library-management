package user

import (
	"context"
	Error "github.com/ewinjuman/go-lib/error"
	pGrpc "github.com/ewinjuman/go-lib/grpc"
	Session "github.com/ewinjuman/go-lib/session"
	"google.golang.org/grpc/connectivity"
	"os"
)

type (
	UserGrpcService interface {
		JwtInterceptor(token string) (*JwtInterceptorResponse, error)
		GetUser(id int) (*UserResponse, error)
	}

	userGrpc struct {
		// client to GRPC service
		session *Session.Session
	}
)

type userGrpcConn struct {
	// client to GRPC service
	userClient    UserServiceClient
	rpcConnection *pGrpc.RpcConnection
}

var rpcConnection *pGrpc.RpcConnection

func getUserClient() (*userGrpcConn, error) {
	if rpcConnection != nil && rpcConnection.Connection != nil && (rpcConnection.Connection.GetState() == connectivity.Connecting || rpcConnection.Connection.GetState() == connectivity.Ready) {
		ctx := &userGrpcConn{
			userClient:    NewUserServiceClient(rpcConnection.Connection),
			rpcConnection: rpcConnection,
		}
		return ctx, nil

	}
	conf := pGrpc.Options{
		Address: os.Getenv("USER_GRPC"),
		Timeout: 20,
	}
	userConn, err := pGrpc.New(conf)
	if err != nil {
		return nil, err
	}
	rpcConnection = userConn
	ctx := &userGrpcConn{
		userClient:    NewUserServiceClient(rpcConnection.Connection),
		rpcConnection: rpcConnection,
	}
	return ctx, nil
}

// NewServerContext constructor for server context
func NewUserGrpc(session *Session.Session) UserGrpcService {
	return &userGrpc{session: session}
}

func (s *userGrpc) JwtInterceptor(token string) (*JwtInterceptorResponse, error) {
	conn, err := getUserClient()
	if err != nil {
		return &JwtInterceptorResponse{}, err
	}
	clientCtx, cancel := conn.rpcConnection.CreateContext(context.Background(), s.session)
	defer cancel()
	request := &JwtInterceptorRequest{Token: token}
	result, err := conn.userClient.JwtInterceptor(clientCtx, request)
	if err != nil {
		return &JwtInterceptorResponse{}, Error.ParseError(err)
	}

	return result, nil
}

func (s *userGrpc) GetUser(id int) (*UserResponse, error) {
	conn, err := getUserClient()
	if err != nil {
		return &UserResponse{}, err
	}
	clientCtx, cancel := conn.rpcConnection.CreateContext(context.Background(), s.session)
	defer cancel()
	request := &GetUserRequest{Id: uint32(id)}
	result, err := conn.userClient.GetUser(clientCtx, request)
	if err != nil {
		return &UserResponse{}, Error.ParseError(err)
	}

	return result, nil
}
