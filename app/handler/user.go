package handler

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateStandardUser(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	success, err := h.UserService.CreateUser(c, newUser, model.Standard)

	if success {
		c.IndentedJSON(http.StatusCreated, newUser)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
