package user

import (
	"context"
	"member/entities"
)

func (g Gorm) Create(ctx context.Context, user entities.User) (*entities.User, error) {
	if err := g.WithTxFromCtx(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
