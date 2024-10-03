# Внедрение Prometheus в приложение

> Далее идут примеры кода и конфигурации.

Определить структуру и методы для обработки метрик

```golang
type PrometheusMetrics struct {
    counter   *prometheus.GaugeVec
    histogram *prometheus.HistogramVec
}

func NewPrometheusMetrics() *PrometheusMetrics {
    counter := prometheus.NewGaugeVec(prometheus.GaugeOpts{
    Name: "http_request_count",
    Help: "Total number of HTTP requests",
    }, []string{"path"})

    histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
    Name: "time_counter_of_methods",
    Help: "request processing time",
    }, []string{"method"})
    prometheus.MustRegister(counter, histogram)
    return &PrometheusMetrics{counter: counter, histogram: histogram}
}

func (p *PrometheusMetrics) CountRequest(url string) {
    p.counter.With(prometheus.Labels{"path": url}).Inc()
}

func (p *PrometheusMetrics) TimeCounting(method string, start time.Time) {
    p.histogram.WithLabelValues(method).Observe(time.Since(start).Seconds())
}
```
Определить хэндлер, через который Prometheus будет собирать метрики

```golang
func (r *Router) Metrics(router *gin.RouterGroup) {
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
```
Настроить конфигурационный файл prometheus.yml

```yaml
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

scrape_configs:
  - job_name: 'prometheus'
    scrape_interval: 5s
    metrics_path: '/metrics'
    static_configs:
      - targets: ['api:8080']

```

Настроить docker-compose.yml

```yaml
version: '3.7'
services:
  api:
    build: ./
    container_name: api
    ports:
      - 8080:8080
  prometheus:
    image: prom/prometheus
    container_name: prometheus
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - 9090:9090
```