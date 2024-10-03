package suite

import (
	"context"
	"net"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/pb/sso"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient sso.AuthClient
}

const grpcHost = "localhost"

func NewSuite(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadPath(configPath())

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(
		func() {
			t.Helper()
			cancelCtx()
		},
	)

	cc, err := grpc.DialContext(
		context.Background(), grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: sso.NewAuthClient(cc),
	}
}

func configPath() string {
	const key = "CONFIG_PATH"

	if v := os.Getenv(key); v != "" {
		return v
	}

	return "../.env"
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
