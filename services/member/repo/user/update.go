package user

import (
	"context"
	"member/entities"
)

func (g Gorm) Update(ctx context.Context, userId string, updateModel entities.User) error {
	return g.WithTxFromCtx(ctx).
		Model(&entities.User{}).
		Where("id = ?", userId).
		Updates(updateModel).Error
}
