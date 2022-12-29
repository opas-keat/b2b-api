package dealer

import (
	"context"
	"errors"
	"models/dealer"
	"strconv"
	"systemlink/entities"
)

func (s ServiceImpl) GetDealerByCode(ctx context.Context, code string) (*dealer.DealerResponse, error) {
	d, err := s.dealerRepo.Get(ctx, entities.Dealer{FTCustCode: code})
	if err != nil {
		return nil, errors.New("get dealer by code not found")
	}
	return &dealer.DealerResponse{
		ID:      strconv.FormatUint(uint64(d.FNMSysCustId), 10),
		Code:    d.FTCustCode,
		Address: d.FTCustAddressInv,
		Name:    d.FTCustName,
		Phone:   d.FTCustPhoneInv,
	}, nil
}
