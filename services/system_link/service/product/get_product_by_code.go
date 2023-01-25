package product

import (
	"context"
	"errors"
	"fmt"
	"models/product"
	"strconv"
	"systemlink/entities"
)

func (s ServiceImpl) GetProductByCode(ctx context.Context, code string) (*product.SystemLinkProductResponse, error) {
	fmt.Println("---- GetProductByCode--- code = " + code)
	p, err := s.productRepo.Get(ctx, entities.Product{FTProdCode: code})
	if err != nil {
		return nil, errors.New("get product by code not found")
	}
	fmt.Println("---- GetProductByCode--- group code = " + p.FTProdGrpCode)
	year, week := "", ""
	if p.FTProdGrpCode == "TIRE-NEW PP" {
		year, week, err = s.productRepo.GetTiresYear(ctx, strconv.FormatUint(uint64(p.FNMSysProdId), 10))
		if err != nil {
			return nil, errors.New("get year by product id not found")
		}
	}
	balance, err := s.productRepo.GetBalance(ctx, strconv.FormatUint(uint64(p.FNMSysProdId), 10))
	if err != nil {
		return nil, errors.New("get year by product id not found")
	}
	fmt.Println("---- GetProductByCode--- balance = " + fmt.Sprintf("%f", balance))

	return &product.SystemLinkProductResponse{
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
		TreadWare:       p.FTTreadwareNameTH,
		Width:           p.FTWidthNameTH,
		GroupCode:       p.FTProdGrpCode,
		LoadIndex:       p.FTLoadIndexNameTH,
		SpeedIndex:      p.FTSpeedIndexNameTH,
		Year:            year,
		Week:            week,
		Balance:         balance,
	}, nil
}
