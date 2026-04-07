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

func GetData(ctx context.Context) string {
	DataCallsTotal.Inc()
	slog.LogAttrs(ctx, slog.LevelInfo, "GetData running")
	time.Sleep(10 * time.Millisecond)
	return "data"
}
