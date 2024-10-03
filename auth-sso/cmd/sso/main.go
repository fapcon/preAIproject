package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/clients/auth"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/server"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/server/handlers"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/app"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/lib/logger"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("cannot load .env file: " + err.Error())
	}
	// Init config
	cfg := config.MustLoad()

	// Init logger
	log := logger.Get(cfg.Logger.Level)
	log.Debug().Msgf("Starting sso service: %s ", cfg.Env)

	// Init auth sso application
	application := app.NewApp(log, cfg.GRPC.Port, cfg, cfg.TokenTTL)

	httpAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	grpcAddr := fmt.Sprintf(":%d", cfg.GRPC.Port)

	ssoClient, err := auth.NewClient(context.Background(), log, grpcAddr, cfg.Client)
	if err != nil {
		log.Error().Msgf("Failed to create clients: %v", err)
	}

	authHandler := handlers.NewAuth(ssoClient)

	httpServer := server.CreateHTTPServer(httpAddr, authHandler)

	go func() {
		log.Info().Msgf("http server is running on: %s", httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msgf("Failed to start http server: %v", err)
		}
	}()

	go func() {
		application.GRPCServer.MustRun()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sigName := <-done

	log.Info().Msgf("Stopping sso service: %s", sigName)

	application.GRPCServer.Stop()

	log.Info().Msgf("sso service stopped gracefully")
}
