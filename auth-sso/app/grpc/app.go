package grpcapp

import (
	"context"
	"fmt"
	"net"

	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type App struct {
	log        *zerolog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPCApp(log *zerolog.Logger, authService auth.Auther, port int) *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(
			func(p interface{}) error {
				log.Error().Msgf("Recovered from panic: %v", p)

				return status.Errorf(codes.Internal, "internal server error")
			},
		),
	}

	gRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
			logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
		),
	)

	auth.Register(gRPCServer, authService)

	// Рефлексия для тестирования
	reflection.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// InterceptorLogger adapts zerolog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			l := l.With().Fields(fields).Logger()

			switch lvl {
			case logging.LevelDebug:
				l.Debug().Msg(msg)
			case logging.LevelInfo:
				l.Info().Msg(msg)
			case logging.LevelWarn:
				l.Warn().Msg(msg)
			case logging.LevelError:
				l.Error().Msg(msg)
			default:
				panic(fmt.Sprintf("unknown level: %v", lvl))
			}
		},
	)
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With().Str("op", op).Logger()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Error().Err(err).Msgf("failed to listen")

		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info().Msgf("gRPC server started... on port %d", a.port)

	if err := a.gRPCServer.Serve(l); err != nil {
		log.Error().Err(err).Msgf("failed to serve")

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.gRPCServer.Stop()

	log := a.log.With().Str("op", op).Logger()

	log.Info().Msgf("stopping gRPC server on port %d", a.port)

	a.gRPCServer.GracefulStop()
}
