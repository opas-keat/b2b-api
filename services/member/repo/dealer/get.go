package dealer

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"member/entities"
	"models/status_code"
	"shareerrors"
)

func (g Gorm) Get(ctx context.Context, filter entities.Dealer) (*entities.Dealer, error) {
	var dealer entities.Dealer
	if err := g.GetDB(ctx).Where(filter).First(&dealer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shareerrors.NewError(status_code.NotFound)
		}
		return nil, err
	}

	return &dealer, nil
}
