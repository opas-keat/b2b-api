package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	AccessTokenID string `json:"access_token_id"`
	UserID        string `json:"user_id"`
}
