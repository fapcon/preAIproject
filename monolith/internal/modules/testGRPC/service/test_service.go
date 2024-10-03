package service

import (
	"context"
	"google.golang.org/grpc"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/testGRPC/greet_grpc"
)

type TestServiceGRPC struct {
	client greet_grpc.GreetServiceClient
}

func NewTestService(conn *grpc.ClientConn) *TestServiceGRPC {
	return &TestServiceGRPC{greet_grpc.NewGreetServiceClient(conn)}
}

func (t *TestServiceGRPC) Hello() (string, error) {
	out := &greet_grpc.Request{Name: "Test"}
	message, err := t.client.Hello(context.Background(), out)
	if err != nil {
		return "", err
	}
	return message.GetAnswer(), nil
}

func (t *TestServiceGRPC) Bye() (string, error) {
	out := &greet_grpc.Request{Name: "Test"}
	message, err := t.client.Bye(context.Background(), out)
	if err != nil {
		return "", err
	}
	return message.GetAnswer(), nil
}
