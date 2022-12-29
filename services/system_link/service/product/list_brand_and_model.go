package product

import (
	"context"
	"github.com/AlekSi/pointer"
	lop "github.com/samber/lo/parallel"
	"models"
	"models/product"
	"systemlink/entities"
)

func (s ServiceImpl) ListBrandAndModel(ctx context.Context, filter product.ListBrandAndModelRequest, pagination models.Pagination) (*models.ListResponse[product.BrandAndModelResponse], error) {
	var brandName, productGroup []string
	if filter.ProductType == "1" {
		brandName = []string{"COSMIS", "COSMIS2", "FATTAH", "FORCE", "NAYA", "UNIVERSE", "VALENZA"}
		productGroup = []string{"COMPLICATED", "FACTORY"}
	} else {
		brandName = []string{"COSMIS", "DOUBLESTAR", "ZESTINO"}
		productGroup = []string{"TIRE-NEW PP"}
	}
	result, count, err := s.productRepo.ListBrandAndModel(ctx, &pagination, pointer.To(models.Count{}), brandName, productGroup)
	if err != nil {
		return nil, err
	}

	rows := lop.Map[entities.BrandAndModel, product.BrandAndModelResponse](*result, func(bam entities.BrandAndModel, _ int) product.BrandAndModelResponse {
		return product.BrandAndModelResponse{
			Brand: bam.FTBrandNameTH,
			Model: bam.FTModelNameTH,
		}
	})
	return &models.ListResponse[product.BrandAndModelResponse]{
		Rows:       rows,
		TotalCount: count,
	}, nil
}
