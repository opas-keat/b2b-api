package entities

type Shipping struct {
	FNMSysShipId         uint   `gorm:"column:FNMSysShipId"`
	FTShipCode           string `gorm:"column:FTShipCode"`
	FTShipNameTH         string `gorm:"column:FTShipNameTH"`
	FTShipNameEN         string `gorm:"column:FTShipNameEN"`
	FTTel                string `gorm:"column:FTTel"`
	FTMobile             string `gorm:"column:FTMobile"`
	FTNote               string `gorm:"column:FTNote"`
	FTStateActive        string `gorm:"column:FTStateActive"`
	FTCentralRegion      string `gorm:"column:FTCentralRegion"`
	FTTheNorthEastRegion string `gorm:"column:FTTheNorthEastRegion"`
	FTNorthRegion        string `gorm:"column:FTNorthRegion"`
	FTWestRegion         string `gorm:"column:FTWestRegion"`
	FTSouthernRegion     string `gorm:"column:FTSouthernRegion"`
	FTEastRegion         string `gorm:"column:FTEastRegion"`
	FTBKKRegion          string `gorm:"column:FTBKKRegion"`
}

func (*Shipping) TableName() string {
	return "DBWebApp.dbo.V_MasterShipping"
}
