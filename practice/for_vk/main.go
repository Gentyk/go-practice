package main

import (
	"github.com/gin-gonic/gin"
	"practice/for_vk/domain"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": domain.GetData(),
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
