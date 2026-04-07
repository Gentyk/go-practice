package domain

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var DataCallsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "data_calls_total",
		Help: "Total number of /data handler business calls.",
	},
)

var DataLatencySeconds = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "data_latency_seconds",
		Help:    "Latency of data business function in seconds.",
		Buckets: prometheus.DefBuckets,
	},
)

func GetData(ctx context.Context) string {
	start := time.Now()
	defer DataLatencySeconds.Observe(time.Since(start).Seconds())

	DataCallsTotal.Inc()
	slog.LogAttrs(ctx, slog.LevelInfo, "GetData running")
	time.Sleep(10 * time.Millisecond)
	return "data"
}
