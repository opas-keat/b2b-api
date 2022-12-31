package user

import (
	dealerconnector "connector/dealer"
	"member/repo/user"
)

type ServiceImpl struct {
	userRepo        user.Repo
	dealerConnector dealerconnector.DealerConnector
}

func New(
	userRepo user.Repo,
	dealerConnector dealerconnector.DealerConnector,
) (Service, error) {
	return &ServiceImpl{
		userRepo,
		dealerConnector,
	}, nil
}
