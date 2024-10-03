package metrics

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type ChartData struct {
	Timestamps []time.Time
	Values     []float64
}

func InitECharts() *components.Page {
	page := components.NewPage()
	page.PageTitle = "Prometheus Metrics Visualization"
	page.AddCharts(
		createChart("Line Chart", "http_request_count"),
		createChart("Bar Chart", "time_counter_of_methods"),
	)
	return page
}

func createChart(title, query string) *charts.Line {
	chart := charts.NewLine()
	chart.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
	)

	result := queryPrometheus(query)
	xAxisData := make([]string, len(result.Timestamps))
	yAxisData := make([]opts.LineData, len(result.Timestamps))
	for i, t := range result.Timestamps {
		xAxisData[i] = t.Format("2006-01-02 15:04:05")
		yAxisData[i] = opts.LineData{
			Value: result.Values[i],
		}
	}

	chart.SetXAxis(xAxisData).
		AddSeries(title, yAxisData)

	return chart
}

func queryPrometheus(query string) *ChartData {
	prometheusURL := "prometheus url"

	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		fmt.Println("Error creating Prometheus client:", err)
		return nil
	}

	v1api := v1.NewAPI(client)
	now := time.Now()

	result, _, err := v1api.QueryRange(context.TODO(), query, v1.Range{
		Start: now.Add(-1 * time.Hour),
		End:   now,
		Step:  time.Minute,
	})
	if err != nil {
		fmt.Println("Error querying Prometheus:", err)
		return nil
	}

	matrix, ok := result.(model.Matrix)
	if !ok {
		fmt.Println("Prometheus result is not a matrix")
		return nil
	}

	data := &ChartData{
		Timestamps: make([]time.Time, 0, len(matrix)),
		Values:     make([]float64, 0, len(matrix)),
	}

	for _, sample := range matrix {
		for _, s := range sample.Values {
			timestamp := time.Unix(int64(s.Timestamp), 0)
			data.Timestamps = append(data.Timestamps, timestamp)
			data.Values = append(data.Values, float64(s.Value))
		}
	}

	return data
}

func InitHTTPHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := InitECharts()
		page.Render(w)
	})
}
