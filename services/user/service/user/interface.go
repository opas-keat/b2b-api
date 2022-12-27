package user

import (
	"context"
	"models/user"
)

type Service interface {
	Register(ctx context.Context, req user.RegisterUserRequest) (*user.CreateUserResponse, error)
	//Me(ctx context.Context, userDetail gateway.User) (*user.MeResponse, error)
	Login(ctx context.Context, username, password string) (*user.LoginResponse, error)
}
