package product

import (
	"context"
	"models"
	"product/entities"
)

type Repo interface {
	Create(ctx context.Context, member entities.Product) (*entities.Product, error)
	Get(ctx context.Context, filter entities.Product) (*entities.Product, error)
	ListBrandAndModel(ctx context.Context, pagination *models.Pagination, count *models.Count, brandName []string, productGroup []string) (*[]entities.BrandAndModel, int64, error)
	//Update(ctx context.Context, dealerId string, updateModel entities.Dealer) error
	//CreateBatch(ctx context.Context, dealers []entities.Dealer) error
}
