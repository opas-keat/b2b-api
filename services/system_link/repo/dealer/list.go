package dealer

import (
	"context"
	"models"
	"models/status_code"
	"shareerrors"
	"systemlink/entities"
)

func (g Gorm) List(ctx context.Context, filter entities.Dealer, pagination *models.Pagination, count *models.Count) (*[]entities.Dealer, int64, error) {
	var (
		dealerList []entities.Dealer
		c          int64
	)

	query := g.WithTxFromCtx(ctx).Table(`DBWebApp.dbo.V_Customer d`)
	if filter.FTCustCode != "" {
		query = query.Where("d.FTCustCode like ?", filter.FTCustCode+"%")
	}

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
			Find(&dealerList).Error; err != nil {
			return nil, 0, shareerrors.NewError(status_code.NotFound, "dealer not found")
		}
		return &dealerList, c, nil
	}
	return &dealerList, c, nil
}
