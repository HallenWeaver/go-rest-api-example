package handler

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	var loginUser model.TokenRequest

	if err := c.BindJSON(&loginUser); err != nil {
		errormsg := fmt.Sprintf("Unable to process request payload; error: %+v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errormsg})
		return
	}

	user, err := h.UserService.LoginUser(c, loginUser)
	if err != nil {
		errormsg := fmt.Sprintf("Unable to find user with given credentials; error: %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errormsg})
		return
	}

	tokenString, err := generateJWT(user.ID.String())
	if err != nil {
		errormsg := fmt.Sprintf("Unable to generate token for user; error: %+v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errormsg})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func generateJWT(userID string) (tokenString string, err error) {
	jwtKey := []byte(os.Getenv("JWT_TOKEN"))
	claims := model.JWTTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-rest-api-example",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
