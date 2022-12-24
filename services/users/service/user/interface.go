package user

import (
	"context"
	"models/gateway"
	"models/users"
)

type Service interface {
	Create(ctx context.Context, userDetail gateway.User, req users.CreateMemberRequest) (*users.CreateMemberResponse, error)
}
