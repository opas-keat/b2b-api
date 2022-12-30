package dealer

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"member/entities"
	"models/status_code"
	"shareerrors"
)

func (g Gorm) Get(ctx context.Context, filter entities.User) (*entities.User, error) {
	var user entities.User
	if err := g.GetDB(ctx).Where(filter).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shareerrors.NewError(status_code.NotFound)
		}
		return nil, err
	}

	return &user, nil
}
