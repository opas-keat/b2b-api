package user

import (
	dealerconnector "connector/dealer"
	"member/repo/dealer"
	"member/repo/user"
)

type ServiceImpl struct {
	userRepo        user.Repo
	dealerRepo      dealer.Repo
	dealerConnector dealerconnector.DealerConnector
}

func New(
	userRepo user.Repo,
	dealerRepo dealer.Repo,
	dealerConnector dealerconnector.DealerConnector,
) (Service, error) {
	return &ServiceImpl{
		userRepo,
		dealerRepo,
		dealerConnector,
	}, nil
}
