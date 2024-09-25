package model

import "github.com/golang-jwt/jwt/v5"

type MyClaims struct {
	UserID    int    `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
