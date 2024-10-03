package grpc_clients

import (
	"google.golang.org/grpc"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/grpc/auth"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/grpc/user"
)

type GRPCClients struct {
	User user.UserServiceGRPCClient
	Auth auth.AuthServiceGRPCClient
}

func NewGRPCClients(conn *grpc.ClientConn) *GRPCClients {
	userClient := user.NewUserServiceGRPCClient(conn)
	authClient := auth.NewAuthServiceGRPCClient(conn)

	return &GRPCClients{
		User: userClient,
		Auth: authClient}
}
