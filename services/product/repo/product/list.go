package product

import (
	"context"
	"models"
	"models/status_code"
	"product/entities"
	"shareerrors"
)

func (g Gorm) ListBrandAndModel(ctx context.Context, pagination *models.Pagination, count *models.Count, brandName []string, productGroup []string) (*[]entities.BrandAndModel, int64, error) {
	var (
		bams []entities.BrandAndModel
		c    int64
	)
	query := g.WithTxFromCtx(ctx).Table(new(entities.Product).TableName()).Select("brand,model")
	query = query.Where("brand IN ?", brandName)
	query = query.Where("prod_grp_code IN ?", productGroup)
	query = query.Group("brand,model")

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
			Find(&bams).Error; err != nil {
			return nil, 0, shareerrors.NewError(status_code.NotFound, "brand and model not found")
		}
		return &bams, c, nil
	}

	if err := query.Find(&bams).Error; err != nil {
		return nil, 0, err
	}

	return &bams, c, nil
}
