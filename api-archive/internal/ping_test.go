package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	// Создаем Gin router
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ping", Ping) // Регистрируем маршрут

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Создаем тестовый HTTP запрос
	req, _ := http.NewRequestWithContext(ctx, "GET", "/ping", nil)

	// Создаем ResponseRecorder для захвата ответа
	w := httptest.NewRecorder()

	// Выполняем запрос
	r.ServeHTTP(w, req)

	// Проверяем статус код
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	} 
}
