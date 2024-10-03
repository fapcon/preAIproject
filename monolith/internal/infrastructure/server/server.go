package server

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"time"
)

type Server interface {
	Serve(ctx context.Context) error
}

type HTTPServer struct {
	conf     config.Server
	logger   *zap.Logger
	srv      *http.Server
	registry *prometheus.Registry
}

func NewHTTPServer(conf config.Server, server *http.Server, logger *zap.Logger, registry *prometheus.Registry) Server {
	return &HTTPServer{conf: conf, logger: logger, srv: server, registry: registry}
}

func (s *HTTPServer) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)
	go func() {
		if err = s.ServePrometheus(); err != nil {
			chErr <- err
		}
	}()

	go func() {
		s.logger.Info("server started", zap.String("port", s.conf.Port))
		if err = s.srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Error("http listen and serve error", zap.Error(err))
			chErr <- err
		}
	}()

	select {
	case <-chErr:
		return err
	case <-ctx.Done():
	}

	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.conf.ShutdownTimeout*time.Second)
	defer cancel()
	err = s.srv.Shutdown(ctxShutdown)

	return err
}

func (s *HTTPServer) ServePrometheus() error {
	metricsHandler := promhttp.HandlerFor(s.registry, promhttp.HandlerOpts{})
	http.Handle("/metrics", promhttp.Handler())
	s.logger.Info("Prometheus metrics endpoint started", zap.String("port", s.conf.PrometheusPort))
	err := http.ListenAndServe(":"+s.conf.PrometheusPort, metricsHandler)
	if err != nil {
		s.logger.Error("Prometheus metrics endpoint error", zap.Error(err))
	}

	return err
}
