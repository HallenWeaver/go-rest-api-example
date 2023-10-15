package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetAPIHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is up and working fine"})
}
