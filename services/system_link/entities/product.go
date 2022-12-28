package entities

type VMasterProduct struct {
	FNMSysProdId        uint    `gorm:"column:FNMSysProdId" json:"fNMSysProdId"`
	FTProdCode          string  `gorm:"column:FTProdCode" json:"fTProdCode"`
	FTProdNameTH        string  `gorm:"column:FTProdNameTH" json:"fTProdNameTH"`
	FTProdMatSizeNameTH string  `gorm:"column:FTProdMatSizeNameTH" json:"fTProdMatSizeNameTH"`
	FTProdColorNameTH   string  `gorm:"column:FTProdColorNameTH" json:"fTProdColorNameTH"`
	FTBrandNameTH       string  `gorm:"column:FTBrandNameTH" json:"fTBrandNameTH"`
	FTModelNameTH       string  `gorm:"column:FTModelNameTH" json:"fTModelNameTH"`
	FTWidthNameTH       string  `gorm:"column:FTWidthNameTH" json:"fTWidthNameTH"`
	FTOffsetNameTH      string  `gorm:"column:FTOffsetNameTH" json:"fTOffsetNameTH"`
	FTTreadwareNameTH   string  `gorm:"column:FTTreadwareNameTH" json:"fTTreadwareNameTH"`
	FTPitchCircleCode   string  `gorm:"column:FTPitchCircleCode" json:"fTPitchCircleCode"`
	FNPrice             float64 `gorm:"column:FNPrice" json:"fNPrice"`
	FPProdImage         []byte  `gorm:"column:FPProdImage" json:"-"`
	FTProdImage         string  `json:"fTProdImage"`
	FNDealerPrice1      float64 `gorm:"column:FNDealerPrice1" json:"fNDealerPrice1"`
}

func (*VMasterProduct) TableName() string {
	return string("DBWebApp.dbo.V_MasterProd")
}
