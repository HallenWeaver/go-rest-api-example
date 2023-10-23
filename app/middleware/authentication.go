package middleware

import (
	"alexandre/gorest/app/model"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := validateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}

func validateToken(signedToken string) (err error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// Validating Algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected token signing method: %v", token.Header["alg"])
		}

		// Return the key
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(model.JWTTokenClaims)
	if !ok {
		err = errors.New("unable to parse claims")
		return
	}
	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
