package dealer

import (
	"systemlink/repo/dealer"
)

type ServiceImpl struct {
	dealerRepo dealer.Repo
}

func New(
	dealerRepo dealer.Repo,
) (Service, error) {
	return &ServiceImpl{
		dealerRepo,
	}, nil
}
