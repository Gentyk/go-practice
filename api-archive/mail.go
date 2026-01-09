package main

import (
	"github.com/gin-gonic/gin"

	"example/web-api-archive/internal"
)

func main() {
	router := gin.Default()
	router.GET("/ping", internal.Ping)

	router.Run("localhost:8080")
}
