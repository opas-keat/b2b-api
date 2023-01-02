package product

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/jinzhu/copier"
	"models/gateway"
	"models/product"
	"models/status_code"
	"product/constant"
	"product/entities"
	"shareerrors"
	"strconv"
)

func (s ServiceImpl) Create(ctx context.Context, userDetail gateway.User, req product.CreateProductRequest) (*product.CreateProductResponse, error) {

	fPrice, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return nil, err
	}
	fDealerPrice1, err := strconv.ParseFloat(req.DealerPrice1, 64)
	if err != nil {
		return nil, err
	}
	//
	//if _, err := s.productRepo.Get(ctx, entities.Product{LinkId: req.LinkId}); err != nil {
	//	return nil, err
	//}

	p, err := s.productRepo.Create(ctx, entities.Product{
		LinkId:          req.LinkId,
		Name:            req.Name,
		Code:            req.Code,
		Model:           req.Model,
		Brand:           req.Brand,
		Price:           fPrice,
		CreatedBy:       pointer.ToString(userDetail.MemberID),
		PitchCircleCode: req.PitchCircleCode,
		Color:           req.Color,
		ProdGrpCode:     req.GroupCode,
		DealerPrice:     fDealerPrice1,
		MatSize:         req.MatSize,
		OffsetName:      req.Offset,
		Treadware:       req.TreadWare,
		WidthName:       req.Width,
	})
	if err != nil {
		if shareerrors.IsCode(err, status_code.Duplicate) {
			return nil, shareerrors.NewError(status_code.Duplicate, constant.CreateProductDuplicate)
		}
		return nil, err
	}
	resp := new(product.CreateProductResponse)
	if err := copier.Copy(resp, p); err != nil {
		return nil, err
	}
	return resp, nil
}
