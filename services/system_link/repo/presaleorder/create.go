package presaleorder

import (
	"context"
	"systemlink/entities"
)

func (g Gorm) Create(ctx context.Context, preSaleOrder entities.PreSaleOrder) (*entities.PreSaleOrder, error) {
	if err := g.WithTxFromCtx(ctx).Create(&preSaleOrder).Error; err != nil {
		return nil, err
	}
	return &preSaleOrder, nil
}
