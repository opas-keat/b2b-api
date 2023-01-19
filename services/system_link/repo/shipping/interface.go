package shipping

import (
	"context"
	"models"
	"systemlink/entities"
)

type Repo interface {
	Get(ctx context.Context, filter entities.Shipping) (*entities.Shipping, error)
	List(ctx context.Context, filter entities.Shipping, pagination *models.Pagination, count *models.Count) (*[]entities.Shipping, int64, error)
}
