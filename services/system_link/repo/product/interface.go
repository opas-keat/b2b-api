package product

import (
	"context"
	"systemlink/entities"
)

type Repo interface {
	Get(ctx context.Context, filter entities.Product) (*entities.Product, error)
}
