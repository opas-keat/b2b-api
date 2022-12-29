package product

import (
	"context"
	"errors"
	"models/product"
	"strconv"
	"systemlink/entities"
)

func (s ServiceImpl) GetProductByCode(ctx context.Context, code string) (*product.ProductResponse, error) {
	p, err := s.productRepo.Get(ctx, entities.Product{FTProdCode: code})
	if err != nil {
		return nil, errors.New("get product by code not found")
	}
	return &product.ProductResponse{
		ID:              strconv.FormatUint(uint64(p.FNMSysProdId), 10),
		Code:            p.FTProdCode,
		Name:            p.FTProdNameTH,
		Brand:           p.FTBrandNameTH,
		Color:           p.FTProdColorNameTH,
		DealerPrice1:    p.FNDealerPrice1,
		MatSize:         p.FTProdMatSizeNameTH,
		Model:           p.FTModelNameTH,
		Offset:          p.FTOffsetNameTH,
		PitchCircleCode: p.FTPitchCircleCode,
		Price:           p.FNPrice,
		ThreadWare:      p.FTTreadwareNameTH,
		Width:           p.FTWidthNameTH,
		GroupCode:       p.FTProdGrpCode,
	}, nil
}
