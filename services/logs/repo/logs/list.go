package logs

import (
	"context"
	"fmt"
	"logs/entities"
	"models"
)

func (g Gorm) List(ctx context.Context, pagination *models.Pagination, count *models.Count, filter entities.Logs) (*[]entities.Logs, int64, error) {
	var (
		logs         []entities.Logs
		c            int64
		sortType     string
		orderByColum string
	)
	query := g.WithTxFromCtx(ctx).Table(`t_logs lg`)
	if *filter.CreatedBy != "" {
		query = query.Where("lg.created_by = ?", filter.CreatedBy)
	}

	orderByColum = "lg.created_at"
	sortType = "DESC"
	query = query.Order(fmt.Sprintf("%s %s", orderByColum, sortType))

	if count != nil {
		if err := query.Count(&c).Error; err != nil {
			return nil, 0, err
		}

		if count.OnlyCount {
			return nil, c, nil
		}
	}

	if pagination != nil {
		if err := query.
			Limit(pagination.Limit).
			Offset(pagination.Offset).
			Find(&logs).Error; err != nil {
			return nil, 0, err
		}
		return &logs, c, nil
	}

	if err := query.Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return &logs, c, nil
}
