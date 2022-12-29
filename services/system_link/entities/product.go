package entities

type Product struct {
	FTModelNameTH       string  `gorm:"column:FTModelNameTH"`
	FTOffsetNameTH      string  `gorm:"column:FTOffsetNameTH"`
	FTProdNameTH        string  `gorm:"column:FTProdNameTH"`
	FTProdMatSizeNameTH string  `gorm:"column:FTProdMatSizeNameTH"`
	FTProdColorNameTH   string  `gorm:"column:FTProdColorNameTH"`
	FTBrandNameTH       string  `gorm:"column:FTBrandNameTH"`
	FTProdCode          string  `gorm:"column:FTProdCode"`
	FTProdGrpCode       string  `gorm:"column:FNDealerPrice1"`
	FTPitchCircleCode   string  `gorm:"column:FTPitchCircleCode"`
	FTTreadwareNameTH   string  `gorm:"column:FTTreadwareNameTH"`
	FTWidthNameTH       string  `gorm:"column:FTWidthNameTH"`
	FNPrice             float64 `gorm:"column:FNPrice"`
	FNDealerPrice1      float64 `gorm:"column:FNDealerPrice1"`
	FNMSysProdId        uint    `gorm:"column:FNMSysProdId"`
}

func (*Product) TableName() string {
	return "DBWebApp.dbo.V_MasterProd"
}

type BrandAndModel struct {
	FTBrandNameTH string `gorm:"column:FTBrandNameTH"`
	FTModelNameTH string `gorm:"column:FTModelNameTH"`
}
