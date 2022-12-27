package user

import "models/gateway"

type LoginRequest struct {
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
	DealerCode string `json:"dealer_code" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" validate:"required,jwt"`
}

type LoginSession struct {
	AccessTokenID  string       `json:"access_token_id"`
	RefreshTokenID string       `json:"refresh_token_id"`
	UserDetail     gateway.User `json:"user_detail"`
}
