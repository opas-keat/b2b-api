package entities

type Product struct {
	FNMSysProdId        uint    `gorm:"column:FNMSysProdId"`
	FTProdCode          string  `gorm:"column:FTProdCode"`
	FTProdNameTH        string  `gorm:"column:FTProdNameTH"`
	FTProdMatSizeNameTH string  `gorm:"column:FTProdMatSizeNameTH"`
	FTProdColorNameTH   string  `gorm:"column:FTProdColorNameTH"`
	FTBrandNameTH       string  `gorm:"column:FTBrandNameTH"`
	FTModelNameTH       string  `gorm:"column:FTModelNameTH"`
	FTWidthNameTH       string  `gorm:"column:FTWidthNameTH"`
	FTOffsetNameTH      string  `gorm:"column:FTOffsetNameTH"`
	FTTreadwareNameTH   string  `gorm:"column:FTTreadwareNameTH"`
	FTPitchCircleCode   string  `gorm:"column:FTPitchCircleCode"`
	FNPrice             float64 `gorm:"column:FNPrice"`
	FNDealerPrice1      float64 `gorm:"column:FNDealerPrice1"`
}

func (*Product) TableName() string {
	return string("DBWebApp.dbo.V_MasterProd")
}

type BrandAndModel struct {
	FTBrandNameTH string `gorm:"column:FTBrandNameTH"`
	FTModelNameTH string `gorm:"column:FTModelNameTH"`
}
