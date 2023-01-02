package product

import (
	"context"
	"product/entities"
)

type Repo interface {
	Create(ctx context.Context, member entities.Product) (*entities.Product, error)
	Get(ctx context.Context, filter entities.Product) (*entities.Product, error)
	//Update(ctx context.Context, dealerId string, updateModel entities.Dealer) error
	//CreateBatch(ctx context.Context, dealers []entities.Dealer) error
}
