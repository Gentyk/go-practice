package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Создаем Gin router
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ping", Ping)

	// Создаем тестовый HTTP запрос
	req := httptest.NewRequest("GET", "/ping", nil)

	// Создаем ResponseRecorder для захвата ответа
	w := httptest.NewRecorder()

	// Выполняем запрос
	r.ServeHTTP(w, req)

	// Проверяем статус код
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
