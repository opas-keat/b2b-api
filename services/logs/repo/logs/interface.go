package logs

import (
	"context"
	"logs/entities"
	"models"
)

type Repo interface {
	Create(ctx context.Context, logs entities.Logs) (*entities.Logs, error)
	List(ctx context.Context, pagination *models.Pagination, count *models.Count, filter entities.Logs) (*[]entities.Logs, int64, error)
}
