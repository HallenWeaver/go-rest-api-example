package handler

import (
	"alexandre/gorest/app/helper"
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	UserService service.IUserService
}

func NewAuthenticationHandler(userService service.IUserService) *AuthenticationHandler {
	return &AuthenticationHandler{
		UserService: userService,
	}
}

func (h *AuthenticationHandler) LoginUser(c *gin.Context) {
	var loginUser model.TokenRequest

	if err := c.BindJSON(&loginUser); err != nil {
		errormsg := fmt.Sprintf("Unable to process request payload; error: %+v", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": errormsg})
		return
	}

	user, err := h.UserService.LoginUser(c, loginUser)
	if err != nil {
		errormsg := fmt.Sprintf("Unable to find user with given credentials; error: %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errormsg})
		return
	}

	tokenString, err := helper.GenerateJWT(user.ID.Hex())
	if err != nil {
		errormsg := fmt.Sprintf("Unable to generate token for user; error: %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errormsg})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
