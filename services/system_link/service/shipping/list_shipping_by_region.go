package shipping

import (
	"context"
	"models"
	"models/shipping"
	"strconv"
	"systemlink/entities"

	"github.com/AlekSi/pointer"
	lop "github.com/samber/lo/parallel"
)

func (s ServiceImpl) ListShippingByRegion(ctx context.Context, filter shipping.ListShippingRequest, pagination models.Pagination) (*models.ListResponse[shipping.Shipping], error) {
	var (
	//dealers *[]entities.Dealer
	//count   int64
	//err     error
	)

	result, count, err := s.shippingRepo.List(ctx, entities.Shipping{
		FTShipCode:           filter.FTShipCode,
		FTShipNameTH:         filter.FTShipNameTH,
		FTCentralRegion:      filter.FTCentralRegion,
		FTTheNorthEastRegion: filter.FTTheNorthEastRegion,
		FTNorthRegion:        filter.FTNorthRegion,
		FTWestRegion:         filter.FTWestRegion,
		FTSouthernRegion:     filter.FTSouthernRegion,
		FTEastRegion:         filter.FTEastRegion,
		FTBKKRegion:          filter.FTBKKRegion,
	}, &pagination, pointer.To(models.Count{}))
	if err != nil {
		return nil, err
	}

	rows := lop.Map[entities.Shipping, shipping.Shipping](*result, func(sh entities.Shipping, _ int) shipping.Shipping {
		return shipping.Shipping{
			ID:     strconv.FormatUint(uint64(sh.FNMSysShipId), 10),
			Code:   sh.FTShipCode,
			Name:   sh.FTShipNameTH,
			Tel:    sh.FTTel,
			Mobile: sh.FTMobile,
			Note:   sh.FTNote,
		}
	})
	return &models.ListResponse[shipping.Shipping]{
		Rows:       rows,
		TotalCount: count,
	}, nil
}
