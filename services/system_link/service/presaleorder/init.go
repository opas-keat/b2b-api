package presaleorder

import (
	"systemlink/repo/presaleorder"
)

type ServiceImpl struct {
	presaleorderRepo presaleorder.Repo
}

func New(
	presaleorderRepo presaleorder.Repo,
) (Service, error) {
	return &ServiceImpl{
		presaleorderRepo,
	}, nil
}
