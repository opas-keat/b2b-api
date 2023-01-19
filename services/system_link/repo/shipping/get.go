package shipping

import (
	"context"
	"errors"
	"models/status_code"
	"shareerrors"
	"systemlink/entities"

	"gorm.io/gorm"
)

func (g Gorm) Get(ctx context.Context, filter entities.Shipping) (*entities.Shipping, error) {
	var shipping entities.Shipping
	if err := g.GetDB(ctx).Where(filter).First(&shipping).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shareerrors.NewError(status_code.NotFound)
		}
		return nil, err
	}
	return &shipping, nil
}
