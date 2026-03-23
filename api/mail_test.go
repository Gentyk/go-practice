package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkGetAlbums(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/albums", getAlbums)

	req := httptest.NewRequest(http.MethodGet, "/albums", nil)

	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
	}
}

func BenchmarkPostAlbums(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/albums", postAlbums)

	initialAlbums := append([]album(nil), albums...)
	body := []byte(`{"id":"10","title":"Bench","artist":"Bench","price":10.0}`)

	b.Cleanup(func() {
		albums = initialAlbums
	})

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/albums", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
	}
}
