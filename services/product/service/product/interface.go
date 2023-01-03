package product

import (
	"context"
	"models"
	"models/gateway"
	"models/product"
)

type Service interface {
	Create(ctx context.Context, userDetail gateway.User, req product.CreateProductRequest) (*product.CreateProductResponse, error)
	ListBrandAndModel(
		ctx context.Context,
		filter product.ListBrandAndModelRequest,
		pagination models.Pagination,
	) (*models.ListResponse[product.BrandAndModelResponse], error)
}
