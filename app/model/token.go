package model

import "github.com/golang-jwt/jwt/v5"

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTTokenClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
