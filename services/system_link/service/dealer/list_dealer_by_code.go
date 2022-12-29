package dealer

import (
	"context"
	"github.com/AlekSi/pointer"
	lop "github.com/samber/lo/parallel"
	"models"
	"models/dealer"
	"strconv"
	"systemlink/entities"
)

func (s ServiceImpl) ListDealerByCode(ctx context.Context, filter dealer.ListDealerRequest, pagination models.Pagination) (*models.ListResponse[dealer.DealerResponse], error) {
	var (
	//dealers *[]entities.Dealer
	//count   int64
	//err     error
	)

	result, count, err := s.dealerRepo.List(ctx, entities.Dealer{
		FTCustCode: filter.FTCustCode,
	}, &pagination, pointer.To(models.Count{}))
	if err != nil {
		return nil, err
	}

	rows := lop.Map[entities.Dealer, dealer.DealerResponse](*result, func(d entities.Dealer, _ int) dealer.DealerResponse {
		return dealer.DealerResponse{
			ID:      strconv.FormatUint(uint64(d.FNMSysCustId), 10),
			Code:    d.FTCustCode,
			Name:    d.FTCustName,
			Address: d.FTCustAddressInv,
			Phone:   d.FTCustPhoneInv,
		}
	})
	return &models.ListResponse[dealer.DealerResponse]{
		Rows:       rows,
		TotalCount: count,
	}, nil
}
