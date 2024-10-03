package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type PrometheusMetrics struct {
	counter   *prometheus.GaugeVec
	histogram *prometheus.HistogramVec
}

func NewPrometheusMetrics(registry *prometheus.Registry) *PrometheusMetrics {
	counter := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_request_count",
		Help: "Total number of HTTP requests",
	}, []string{"path"})

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "time_counter_of_methods",
		Help: "request processing time",
	}, []string{"layer", "method"})

	registry.MustRegister(counter, histogram)
	return &PrometheusMetrics{counter: counter, histogram: histogram}
}

func (p *PrometheusMetrics) CountRequest(url string) {
	p.counter.With(prometheus.Labels{"path": url}).Inc()
}

func (p *PrometheusMetrics) TimeCounting(layer, method string, start time.Time) {
	p.histogram.WithLabelValues(layer, method).Observe(time.Since(start).Seconds())
}
