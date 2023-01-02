package product

import (
	"context"
	"models/gateway"
	"models/product"
)

type Service interface {
	Create(ctx context.Context, userDetail gateway.User, req product.CreateProductRequest) (*product.CreateProductResponse, error)
}
