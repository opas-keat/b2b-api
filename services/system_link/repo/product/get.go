package product

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"models/status_code"
	"shareerrors"
	"systemlink/entities"
)

func (g Gorm) Get(ctx context.Context, filter entities.Product) (*entities.Product, error) {
	var product entities.Product
	if err := g.GetDB(ctx).Where(filter).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shareerrors.NewError(status_code.NotFound)
		}
		return nil, err
	}
	return &product, nil
}
