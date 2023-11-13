package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	healthHandler := NewHealthHandler()

	router := gin.Default()
	router.GET("/health", healthHandler.GetAPIHealth)

	req, _ := http.NewRequest("GET", "/health", nil)
	httpRecorder := httptest.NewRecorder()
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, http.StatusOK, httpRecorder.Code)
}
