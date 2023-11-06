package util

import (
	"alexandre/gorest/app/model"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseUserDataFromToken(c *gin.Context) (string, error) {
	signedToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(signedToken, &model.JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validating Algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected token signing method: %v", token.Header["alg"])
		}

		// Return the key
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*model.JWTTokenClaims)
	if !ok {
		err = errors.New("unable to parse claims")
		return "", err
	}
	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", err
	}

	fmt.Printf("UserID: %+v\n", claims.UserID)

	return claims.UserID, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
