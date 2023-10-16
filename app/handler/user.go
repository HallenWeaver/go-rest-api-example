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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	storedUser, err := h.UserService.CreateUser(c, newUser, model.Standard)

	if storedUser != nil {
		c.IndentedJSON(http.StatusCreated, storedUser)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
	}
}
