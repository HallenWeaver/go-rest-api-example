package handler

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	UserService service.UserService
}

func NewAuthenticationHandler(userService service.UserService) *AuthenticationHandler {
	return &AuthenticationHandler{
		UserService: userService,
	}
}

func (h *AuthenticationHandler) LoginUser(c *gin.Context) {
	var loginUser model.User

	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	loginResult, err := h.UserService.LoginUser(c, loginUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	if !loginResult {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	// Should generate token if logging in is fine
}
