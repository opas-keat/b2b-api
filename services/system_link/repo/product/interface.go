package product

import (
	"context"
	"models"
	"systemlink/entities"
)

type Repo interface {
	Get(ctx context.Context, filter entities.Product) (*entities.Product, error)
	ListBrandAndModel(ctx context.Context, pagination *models.Pagination, count *models.Count, brandName []string, productGroup []string) (*[]entities.BrandAndModel, int64, error)
}
