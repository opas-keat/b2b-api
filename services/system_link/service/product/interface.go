package product

import (
	"context"
	"models/product"
)

type Service interface {
	GetProductByCode(ctx context.Context, code string) (*product.ProductResponse, error)
}
