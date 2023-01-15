package logs

import (
	"context"
	"logs/entities"
	"models/logs"

	"github.com/jinzhu/copier"
)

func (s ServiceImpl) Create(ctx context.Context, req logs.CreateLogsRequest) (*logs.CreateLogsResponse, error) {
	p, err := s.logsRepo.Create(ctx, entities.Logs{
		Detail:    req.Detail,
		CreatedBy: &req.CreatedBy,
		// CreatedAt: now(),
	})
	if err != nil {
		return nil, err
	}
	resp := new(logs.CreateLogsResponse)
	if err := copier.Copy(resp, p); err != nil {
		return nil, err
	}
	return resp, nil
}
