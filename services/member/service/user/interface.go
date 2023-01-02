package user

import (
	"context"
	"models/gateway"
	"models/user"
)

type Service interface {
	Login(ctx context.Context, username, password string) (*user.LoginResponse, error)
	Register(ctx context.Context, req user.RegisterUserRequest) (*user.CreateUserResponse, error)
	VerifyEmail(ctx context.Context, verificationCode string) (*user.VerifyEmailResponse, error)
	Me(ctx context.Context, userDetail gateway.User) (*user.MeResponse, error)
}
