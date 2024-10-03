package logger

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once

var log zerolog.Logger

func Get(logLevel string) *zerolog.Logger {
	once.Do(
		func() {
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
			zerolog.TimeFieldFormat = time.RFC822

			logLevel, err := strconv.Atoi(logLevel)
			if err != nil {
				logLevel = int(zerolog.InfoLevel) // default to INFO
			}

			log = zerolog.New(
				zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: time.RFC822,
				},
			).
				Level(zerolog.Level(logLevel)).
				With().
				Timestamp().
				Str("service", "auth-sso").
				Caller().
				Logger()
		},
	)

	return &log
}
