package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/testGRPC/greet_grpc"
	"syscall"
)

func main() {
	app := &App{}
	app.RegisterGrpc()
	go func() {
		log.Println("Server starting")
		if err := app.Serve(); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server exit")
	app.grpcServer.GracefulStop()

}

type App struct {
	grpcServer *grpc.Server
	greet_grpc.UnimplementedGreetServiceServer
}

func (a *App) RegisterGrpc() {
	a.grpcServer = grpc.NewServer()
	greet_grpc.RegisterGreetServiceServer(a.grpcServer, a)
}

func (a *App) Serve() error {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		return err
	}

	return a.grpcServer.Serve(listen)
}

func (a *App) Hello(ctx context.Context, req *greet_grpc.Request) (*greet_grpc.Message, error) {
	log.Println("grpc hello")
	out := a.GreetWithGrpc(req.GetName())
	return &greet_grpc.Message{
		ID:     int64(len(out)),
		Answer: out,
	}, nil
}

func (a *App) Bye(ctx context.Context, req *greet_grpc.Request) (*greet_grpc.Message, error) {
	log.Println("grpc bye")
	out := a.SayByeToGrpc(req.GetName())
	return &greet_grpc.Message{
		ID:     int64(len(out)),
		Answer: out,
	}, nil
}

func (a *App) GreetWithGrpc(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func (a *App) SayByeToGrpc(name string) string {
	return fmt.Sprintf("Bye, %s...", name)
}
