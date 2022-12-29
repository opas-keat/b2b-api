package entities

type Dealer struct {
	FNMSysCustId     uint   `gorm:"column:FNMSysCustId"`
	FTCustCode       string `gorm:"column:FTCustCode"`
	FTCustName       string `gorm:"column:FTCustName"`
	FTCustAddressInv string `gorm:"column:FTCustAddressInv"`
	FTCustPhoneInv   string `gorm:"column:FTCustPhoneInv"`
}

func (*Dealer) TableName() string {
	return "DBWebApp.dbo.V_Customer"
}
