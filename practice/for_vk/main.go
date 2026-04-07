package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/data", func(c *gin.Context) {
		traceID := uuid.NewString()
		logger := slog.Default().With("trace_id", traceID)
		ctx := context.WithValue(c.Request.Context(), "logger", logger)
		c.JSON(200, gin.H{
			"message": domain.GetData(ctx),
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
