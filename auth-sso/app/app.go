package app

import (
	"context"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/social"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/social/facebook"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/social/google"
	"time"

	grpcapp "studentgit.kata.academy/eazzyearn/students/mono/auth-sso/app/grpc"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/database/migrator"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/database/postgres"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/repository"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/repository/storage/psql"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth/service"

	"github.com/rs/zerolog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *zerolog.Logger, grpcPort int, cfg *config.Config, tokenTTL time.Duration) *App {
	pgContext := context.Background()

	// init postgres db
	storage, err := postgres.NewPostgres(pgContext, cfg.DB, log)
	if err != nil {
		log.Fatal().Msgf("Failed to connect to database: %v", err)
	}
	// init postgres storage
	pgStorage := psql.NewStorage(storage)

	// init user repository
	userRepo := repository.NewUserRepository(pgStorage)

	// init app repository
	appRepo := repository.NewAppRepository(pgStorage)

	// Migrate database
	if err := migrator.Migrate(storage, cfg); err != nil {
		log.Fatal().Msgf("Failed to migrate database: %v", err)
	}

	//init socials
	googleAuth := google.NewGoogleAuth(cfg.Google)

	facebookAuth := facebook.NewFacebookAuth(cfg.Facebook)

	authSocial := social.NewAuthSocial(googleAuth, facebookAuth)

	// init auth service
	authService := service.NewAuth(log, userRepo, appRepo, authSocial, cfg)

	// init grpc app
	grpcApp := grpcapp.NewGRPCApp(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
