package util

import (
	"alexandre/gorest/app/model"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func parseUserDataFromToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
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

	claims, ok := token.Claims.(model.JWTTokenClaims)
	if !ok {
		err = errors.New("unable to parse claims")
		return "", err
	}
	if claims.ExpiresAt.Time.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", err
	}

	return claims.Username, nil
}
