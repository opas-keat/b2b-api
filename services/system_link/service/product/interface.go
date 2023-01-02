package product

import (
	"context"
	"models"
	"models/product"
)

type Service interface {
	GetProductByCode(
		ctx context.Context,
		code string,
	) (*product.SystemLinkProductResponse, error)
	ListBrandAndModel(
		ctx context.Context,
		filter product.ListBrandAndModelRequest,
		pagination models.Pagination,
	) (*models.ListResponse[product.BrandAndModelResponse], error)
}
