package product

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"models/status_code"
	"product/entities"
	"shareerrors"
)

func (g Gorm) Create(ctx context.Context, product entities.Product) (*entities.Product, error) {
	if err := g.WithTxFromCtx(ctx).Create(&product).Error; err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, shareerrors.NewError(status_code.Duplicate)
			}
		}
		return nil, err
	}
	return &product, nil
}
