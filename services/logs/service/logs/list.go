package logs

import (
	"context"
	"logs/entities"
	"models"
	"models/logs"

	lop "github.com/samber/lo/parallel"

	"github.com/AlekSi/pointer"
)

func (s ServiceImpl) List(ctx context.Context, filter logs.ListLogsRequest, pagination models.Pagination) (*models.ListResponse[logs.ListLogsResponse], error) {

	result, count, err := s.logsRepo.List(ctx, &pagination, pointer.To(models.Count{}), entities.Logs{CreatedBy: &filter.CreatedBy, Detail: filter.Detail})
	if err != nil {
		return nil, err
	}
	rows := lop.Map[entities.Logs, logs.ListLogsResponse](*result, func(l entities.Logs, _ int) logs.ListLogsResponse {
		return logs.ListLogsResponse{
			ID:        l.ID,
			Detail:    *&l.Detail,
			CreatedBy: *l.CreatedBy,
			CreatedAt: l.CreatedAt.Format("02/01/2006 15:04:05"),
		}
	})
	return &models.ListResponse[logs.ListLogsResponse]{
		Rows:       rows,
		TotalCount: count,
	}, nil
}
