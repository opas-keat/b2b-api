package logs

import (
	"context"
	"models"
	"models/logs"
)

type Service interface {
	Create(ctx context.Context, req logs.CreateLogsRequest) (*logs.CreateLogsResponse, error)
	List(ctx context.Context, filter logs.ListLogsRequest, pagination models.Pagination) (*models.ListResponse[logs.ListLogsResponse], error)
}
