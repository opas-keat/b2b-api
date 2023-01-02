package product

import (
	"context"
	"models/product"
)

type Service interface {
	Create(ctx context.Context, req product.CreateProductRequest) (*product.CreateProductResponse, error)
}
