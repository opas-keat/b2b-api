package dealer

import (
	"context"
	"member/entities"
)

func (g Gorm) Create(ctx context.Context, dealer entities.Dealer) (*entities.Dealer, error) {
	if err := g.WithTxFromCtx(ctx).Create(&dealer).Error; err != nil {
		return nil, err
	}
	return &dealer, nil
}
