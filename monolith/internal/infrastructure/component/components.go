package component

import (
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/metrics"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	eservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
)

type Components struct {
	Conf         config.AppConf
	Notify       service.Notifier
	TokenManager cryptography.TokenManager
	Responder    responder.Responder
	Decoder      godecoder.Decoder
	Logger       *zap.Logger
	Hash         cryptography.Hasher
	ErrChan      chan models.ErrMessage
	RateLimiter  *eservice.RateLimiter
	Metrics      metrics.MetricMeter
}

func NewComponents(conf config.AppConf, notify service.Notifier, tokenManager cryptography.TokenManager, responder responder.Responder, decoder godecoder.Decoder, hash cryptography.Hasher, logger *zap.Logger, errChan chan models.ErrMessage, metrics metrics.MetricMeter) *Components {
	return &Components{Conf: conf, Notify: notify, TokenManager: tokenManager, Responder: responder, Decoder: decoder, Hash: hash, Logger: logger, ErrChan: errChan, RateLimiter: eservice.NewRateLimiter(), Metrics: metrics}
}
