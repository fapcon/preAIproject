package main

import (
	"os"

	"github.com/joho/godotenv"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/logs"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/run"
)

// @title Auto Trade API
// @version 1.0
// @description Documentation of auto trade API

// @schemes http https
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	conf := config.NewAppConf()
	logger := logs.NewLogger(conf, os.Stdout)
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	app := run.NewApp(conf, logger)
	exitCode := app.Bootstrap().Run()
	os.Exit(exitCode)
}
