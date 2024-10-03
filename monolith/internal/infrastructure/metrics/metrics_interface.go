package metrics

import (
	"time"
)

type MetricMeter interface {
	CountRequest(url string)
	TimeCounting(layer, method string, start time.Time)
}
