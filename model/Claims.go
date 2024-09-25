package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
