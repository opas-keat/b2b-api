package product

import (
	"context"
	"models"
	"models/product"
)

type Service interface {
	GetProductByCode(ctx context.Context, code string) (*product.ProductResponse, error)
	ListBrandAndModel(ctx context.Context, productType string) (*models.ListResponse[product.BrandAndModelResponse], error)
}
