package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/config"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewPostgres(ctx context.Context, dbConf config.DB, logger *zerolog.Logger) (*sqlx.DB, error) {
	var err error
	dbDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name,
	)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * dbConf.Timeout)

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", dbConf.Timeout, err)
		case <-ticker.C:
			db, err := sqlx.Open(dbConf.Driver, dbDSN)
			if err != nil {
				return nil, errors.Wrap(err, "sqlx.Open() failed")
			}

			err = db.PingContext(ctx)
			if err == nil {
				db.SetMaxOpenConns(dbConf.MaxConn)
				db.SetMaxIdleConns(dbConf.MaxConn)
				return db, nil
			}
			logger.Error().Msgf("Failed to connect to database: %v", err)
		}
	}
}
