package dealer

import (
	"context"
	"errors"
	"member/entities"
	"models/dealer"
)

func (s ServiceImpl) GetDealerByCode(ctx context.Context, code string) (*dealer.Dealer, error) {
	d, err := s.dealerRepo.Get(ctx, entities.Dealer{DealerCode: code})
	if err != nil {
		return nil, errors.New("get dealer by code not found")
	}
	return &dealer.Dealer{
		ID:      d.ID,
		Code:    d.DealerCode,
		Address: d.Address,
		Name:    d.Name,
		Phone:   d.Phone,
	}, nil
}
