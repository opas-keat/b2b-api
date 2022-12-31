package user

import (
	"context"
	"models/user"
)

type Service interface {
	Login(ctx context.Context, username, password string) (*user.LoginResponse, error)
}
