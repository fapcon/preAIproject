## Usage
С помощью конструктора встроить структуру (имплементирующую интерфейс MetricMeter) в структуру, методы которой вы хотите фиксировать.

```go
type Example struct {
    metrics MetricMeter
}

func NewExample(metrics MetricMeter) *Example {
    return &Example{metrics: metrics}
}
```

Для подсчета количества запросов на сервер указываем URI:
```go
func (e *Example) GetByID(w http.ResponseWriter, r *http.Request) {
    e.metrics.CountRequest("/id")
}
```

Для подсчета времени выполнения указываем модуль проекта, метод  и время когда началось выполнение метода:
```go
func (e *Example) Save() {
    start := time.Now()
    defer e.metrics.TimeCounting("example", "Save", start)
}
```

Определить URI для хэндлера прометеуса. Далее именно по этому адресу будут отображаться все ваши метрики:
```go
func NewApiRouter() http.Handler {
    r := chi.NewRouter()
    r.Get("/metrics", promhttp.Handler().ServeHTTP)
    return r
```