package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"os"
	"practice/for_vk/domain"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	slog.SetDefault(slog.New(handler))
	prometheus.MustRegister(domain.DataCallsTotal, domain.DataLatencySeconds)

	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/data", func(c *gin.Context) {
		traceID := uuid.New().String()
		ctx := context.WithValue(context.Background(), "traceID", traceID)
		c.JSON(200, gin.H{
			"message": domain.GetData(ctx),
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
