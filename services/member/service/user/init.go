package user

import (
	"member/repo/user"
)

type ServiceImpl struct {
	userRepo user.Repo
}

func New(
	userRepo user.Repo,
) (Service, error) {
	return &ServiceImpl{
		userRepo,
	}, nil
}
