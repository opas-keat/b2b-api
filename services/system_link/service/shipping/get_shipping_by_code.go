package shipping

import (
	"context"
	"errors"
	"models/shipping"
	"strconv"
	"systemlink/entities"
)

func (s ServiceImpl) GetShippingByCode(ctx context.Context, code string) (*shipping.Shipping, error) {
	sh, err := s.shippingRepo.Get(ctx, entities.Shipping{FTShipCode: code})
	if err != nil {
		return nil, errors.New("get shipping by code not found")
	}
	return &shipping.Shipping{
		ID:     strconv.FormatUint(uint64(sh.FNMSysShipId), 10),
		Code:   sh.FTShipCode,
		Name:   sh.FTShipNameTH,
		Tel:    sh.FTTel,
		Mobile: sh.FTMobile,
		Note:   sh.FTNote,
	}, nil
}
