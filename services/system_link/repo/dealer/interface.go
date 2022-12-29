package dealer

import (
	"context"
	"models"
	"systemlink/entities"
)

type Repo interface {
	Get(ctx context.Context, filter entities.Dealer) (*entities.Dealer, error)
	List(ctx context.Context, filter entities.Dealer, pagination *models.Pagination, count *models.Count) (*[]entities.Dealer, int64, error)
}
