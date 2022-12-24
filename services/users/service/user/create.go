package user

import (
	"context"
	"models/gateway"
	"models/users"
)

func (s ServiceImpl) Create(ctx context.Context, userDetail gateway.User, req users.CreateMemberRequest) (*users.CreateMemberResponse, error) {
	resp := new(users.CreateMemberResponse)
	return resp, nil
}
