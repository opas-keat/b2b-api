package user

import (
	"context"
	"models/gateway"
	"models/user"
)

type Service interface {
	Register(ctx context.Context, req user.RegisterUserRequest) (*user.CreateUserResponse, error)
	Login(ctx context.Context, username, password string) (*user.LoginResponse, error)
	VerifyEmail(ctx context.Context, verificationCode string) (*user.VerifyEmailResponse, error)
	Me(ctx context.Context, userDetail gateway.User) (*user.MeResponse, error)
}
