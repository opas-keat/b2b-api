package dealer

import (
	"context"
	"member/entities"
)

func (s ServiceImpl) CreateBatch(ctx context.Context, d *[]entities.Dealer) (*[]entities.Dealer, error) {
	var dealers []entities.Dealer
	_, err := s.CreateBatch(ctx, d)
	if err != nil {
		return nil, err
	}
	return &dealers, nil
}
