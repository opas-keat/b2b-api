package presaleorder

import "context"

func (g Gorm) GetPreSaleOrderNo(ctx context.Context, orderNo string) (string, error) {
	var preSaleNo string
	if response := g.GetDB(ctx).Raw("SELECT ISNULL(MAX(CAST(SUBSTRING(DBWebApp.dbo.TMARTPreSaleOrder.FTSaleOrderNo,10,LEN(DBWebApp.dbo.TMARTPreSaleOrder.FTSaleOrderNo)-4) AS INT)) +1,1 ) FROM DBWebApp.dbo.TMARTPreSaleOrder WHERE FTSaleOrderNo LIKE ?", orderNo+"%").Scan(&preSaleNo); response.Error != nil {
		print("Error occurred while retrieving preSaleNo from the database: " + response.Error.Error())
		return "", response.Error
	}
	return preSaleNo, nil
}
