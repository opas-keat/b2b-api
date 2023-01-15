package logs

import (
	"context"
	"logs/entities"
)

func (g Gorm) Create(ctx context.Context, logs entities.Logs) (*entities.Logs, error) {
	if err := g.WithTxFromCtx(ctx).Create(&logs).Error; err != nil {
		return nil, err
	}
	return &logs, nil
}
