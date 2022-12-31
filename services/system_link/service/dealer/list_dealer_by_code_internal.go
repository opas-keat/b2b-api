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

func (s ServiceImpl) ListDealerByCodeInternal(ctx context.Context, filter dealer.ListDealerRequest, pagination models.Pagination) (*[]dealer.Dealer, error) {
	var (
	//dealers *[]entities.Dealer
	//count   int64
	//err     error
	)

	result, _, err := s.dealerRepo.List(ctx, entities.Dealer{
		FTCustCode: filter.FTCustCode,
	}, &pagination, pointer.To(models.Count{}))
	if err != nil {
		return nil, err
	}

	rows := lop.Map[entities.Dealer, dealer.Dealer](*result, func(d entities.Dealer, _ int) dealer.Dealer {
		return dealer.Dealer{
			ID:      strconv.FormatUint(uint64(d.FNMSysCustId), 10),
			Code:    d.FTCustCode,
			Name:    d.FTCustName,
			Address: d.FTCustAddressInv,
			Phone:   d.FTCustPhoneInv,
		}
	})
	return &rows, nil
}
