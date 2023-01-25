package product

import (
	"context"
	"errors"
	"models/status_code"
	"shareerrors"
	"systemlink/entities"

	"gorm.io/gorm"
)

func (g Gorm) Get(ctx context.Context, filter entities.Product) (*entities.Product, error) {
	var product entities.Product
	if err := g.GetDB(ctx).Where(filter).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shareerrors.NewError(status_code.NotFound)
		}
		return nil, err
	}
	return &product, nil
}

func (g Gorm) GetTiresYear(ctx context.Context, productId string) (string, string, error) {
	query := g.WithTxFromCtx(ctx).Table(`DBWebApp.dbo.V_ProdSMB_Balance_Grade bg`).Select("bg.FTYear as Year,bg.FTWeek as Week")
	query = query.Where("bg.FNMSysProdId = ?", productId)
	query = query.Where("FTStateShowSaleKitPP = ?", "1")
	type Result struct {
		Year string `json:"Year"`
		Week string `json:"Week"`
	}
	var result Result
	if response := query.Scan(&result); response.Error != nil {
		return "", "", shareerrors.NewError(status_code.NotFound)
	}
	return result.Year, result.Week, nil
}

func (g Gorm) GetBalance(ctx context.Context, productId string) (float64, error) {
	query := g.WithTxFromCtx(ctx).Table(`DBWebApp.dbo.V_ProdSMB_Balance_Grade b`).Select("b.FNQuantityBal as balance")
	query = query.Where("b.FNMSysProdId = ?", productId)
	query = query.Where("b.FTStateShowSaleKitPP = ?", "1")
	query = query.Where("b.FNQuantityBal > ?", 0)
	query = query.Order("b.FNQuantityBal desc")
	type Result struct {
		Balance float64 `json:"balance"`
	}
	var result Result
	if response := query.Scan(&result); response.Error != nil {
		return 0, shareerrors.NewError(status_code.NotFound)
	}
	return result.Balance, nil
}
