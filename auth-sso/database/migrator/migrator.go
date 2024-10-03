package migrator

import (
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/models"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	migrateConfig "gitlab.com/golight/migrator/db/config"
	"gitlab.com/golight/migrator/db/migrate"
	"gitlab.com/golight/migrator/db/scanner"
)

// Migrate миграция базы данных: пользователи и приложения
func Migrate(storage *sqlx.DB, cfg *config.Config) error {
	tableScanner := scanner.NewTableScanner()
	// регистрация таблиц
	tableScanner.RegisterTable(&models.User{}, &models.App{})

	var migrateCfg migrateConfig.DB
	migrateCfg = config.ConvertConfigs(cfg, &migrateCfg)
	migrator := migrate.NewMigrator(storage, migrateCfg, tableScanner)
	err := migrator.Migrate()
	if err != nil {
		return err
	}

	log.Info().Msgf("migrations completed")

	return nil
}
