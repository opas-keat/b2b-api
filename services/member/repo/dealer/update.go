package dealer

import (
	"context"
	"member/entities"
)

func (g Gorm) Update(ctx context.Context, dealerId string, updateModel entities.Dealer) error {
	return g.WithTxFromCtx(ctx).
		Model(&entities.Dealer{}).
		Where("id = ?", dealerId).
		Updates(updateModel).Error
}
