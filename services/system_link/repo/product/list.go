package product

import (
	"context"
	"models"
	"models/status_code"
	"shareerrors"
	"systemlink/entities"
)

func (g Gorm) ListBrandAndModel(ctx context.Context, pagination *models.Pagination, count *models.Count, brandName []string, productGroup []string) (*[]entities.BrandAndModel, int64, error) {
	var (
		bams []entities.BrandAndModel
		c    int64
	)
	query := g.WithTxFromCtx(ctx).Table(`DBWebApp.dbo.V_MasterProd mp`).Select("mp.FTBrandNameTH,mp.FTModelNameTH")
	query = query.Where("mp.FTBrandNameTH IN ?", brandName)
	query = query.Where("mp.FTProdGrpCode IN ?", productGroup)
	query = query.Where("mp.FTBrandNameTH > ?", "-")
	query = query.Where("mp.FTModelNameTH > ?", "-")
	query = query.Where("mp.FNDealerPrice1 > ?", 0)
	query = query.Group("mp.FTBrandNameTH,FTModelNameTH")

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
