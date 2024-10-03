package responder

import (
	"context"
	"errors"
	"github.com/ptflp/godecoder"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/metrics"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/response"
	"time"

	"go.uber.org/zap"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})

	ErrorUnauthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorForbidden(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct {
	log *zap.Logger
	godecoder.Decoder
	pm *metrics.PrometheusMetrics
}

func NewResponder(decoder godecoder.Decoder, logger *zap.Logger, pm *metrics.PrometheusMetrics) Responder {
	return &Respond{log: logger, Decoder: decoder, pm: pm}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if err := r.Encode(w, responseData); err != nil {
		r.log.Error("responder json encode error", zap.Error(err))
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	start := time.Now()
	r.log.Info("http response bad request status code", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := r.Encode(w, response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Info("response writer error on write", zap.Error(err))
	}
	r.pm.TimeCounting(`responder`, `ErrorBadRequest`, start)
	r.pm.CountRequest(`error_bad_request`)
}

func (r *Respond) ErrorForbidden(w http.ResponseWriter, err error) {
	start := time.Now()
	r.log.Warn("http resposne forbidden", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	if err := r.Encode(w, response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("response writer error on write", zap.Error(err))
	}
	r.pm.TimeCounting(`responder`, `ErrorForbidden`, start)
	r.pm.CountRequest(`error_forbidden`)
}

func (r *Respond) ErrorUnauthorized(w http.ResponseWriter, err error) {
	start := time.Now()
	r.log.Warn("http resposne Unauthorized", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	if err := r.Encode(w, response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("response writer error on write", zap.Error(err))
	}
	r.pm.TimeCounting(`responder`, `ErrorUnauthorized`, start)
	r.pm.CountRequest(`error_unauthorized`)
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	start := time.Now()
	if errors.Is(err, context.Canceled) {
		return
	}
	r.log.Error("http response internal error", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := r.Encode(w, response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("response writer error on write", zap.Error(err))
	}
	r.pm.TimeCounting(`responder`, `ErrorInternal`, start)
	r.pm.CountRequest(`error_internal`)
}
