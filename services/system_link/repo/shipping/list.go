package shipping

import (
	"context"
	"models"
	"models/status_code"
	"shareerrors"
	"systemlink/entities"
)

func (g Gorm) List(ctx context.Context, filter entities.Shipping, pagination *models.Pagination, count *models.Count) (*[]entities.Shipping, int64, error) {
	var (
		shippingList []entities.Shipping
		c            int64
	)

	query := g.WithTxFromCtx(ctx).Table(`DBWebApp.dbo.V_MasterShipping sh`)
	if filter.FTShipNameTH != "" {
		query = query.Where("sh.FTShipNameTH like ?", "%"+filter.FTShipNameTH+"%")
	}
	// if filter.FTShipCode != "" {
	// 	query = query.Where("sh.FTCustCode like ?", filter.FTShipCode+"%")
	// }

	if filter.FTCentralRegion == "1" {
		query = query.Or("sh.FTCentralRegion = '1'")
	}
	if filter.FTTheNorthEastRegion == "1" {
		query = query.Or("sh.FTTheNorthEastRegion = '1'")
	}
	if filter.FTNorthRegion == "1" {
		query = query.Or("sh.FTNorthRegion = '1'")
	}
	if filter.FTWestRegion == "1" {
		query = query.Or("sh.FTWestRegion = '1'")
	}
	if filter.FTSouthernRegion == "1" {
		query = query.Or("sh.FTSouthernRegion = '1'")
	}
	if filter.FTEastRegion == "1" {
		query = query.Or("sh.FTEastRegion = '1'")
	}
	if filter.FTBKKRegion == "1" {
		query = query.Or("sh.FTBKKRegion = '1'")
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
			Find(&shippingList).Error; err != nil {
			return nil, 0, shareerrors.NewError(status_code.NotFound, "shipping not found")
		}
		return &shippingList, c, nil
	}
	return &shippingList, c, nil
}
