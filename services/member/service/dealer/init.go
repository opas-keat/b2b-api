package dealer

import (
	"member/repo/dealer"
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
