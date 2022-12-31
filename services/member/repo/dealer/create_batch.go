package dealer

import (
	"context"
	"member/entities"
)

func (g Gorm) CreateBatch(ctx context.Context, dealers []entities.Dealer) error {
	if err := g.WithTxFromCtx(ctx).Create(&dealers).Error; err != nil {
		return err
	}
	return nil
}
