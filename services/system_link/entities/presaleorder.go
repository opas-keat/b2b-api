package entities

type PreSaleOrder struct {
	FTSaleOrderNo     string  `gorm:"column:FTSaleOrderNo;primaryKey;"`
	FTInsUser         string  `gorm:"column:FTInsUser"`
	FDInsDate         string  `gorm:"column:FDInsDate"`
	FTInsTime         string  `gorm:"column:FTInsTime"`
	FTUpdUser         string  `gorm:"column:FTUpdUser"`
	FDUpdDate         string  `gorm:"column:FDUpdDate"`
	FTUpdTime         string  `gorm:"column:FTUpdTime"`
	FDSaleOrderDate   string  `gorm:"column:FDSaleOrderDate"`
	FTSaleOrderBy     string  `gorm:"column:FTSaleOrderBy"`
	FTQuotationNo     string  `gorm:"column:FTQuotationNo"`
	FNRevised         int     `gorm:"column:FNRevised"`
	FNVatType         int     `gorm:"column:FNVatType"`
	FNPriority        int     `gorm:"column:FNPriority"`
	FNMSysCustId      int     `gorm:"column:FNMSysCustId"`
	FTDeliveryAddress string  `gorm:"column:FTDeliveryAddress"`
	FDDeliveryDate    string  `gorm:"column:FDDeliveryDate"`
	FNMSysCrTermId    int     `gorm:"column:FNMSysCrTermId"`
	FNCreditDay       int     `gorm:"column:FNCreditDay"`
	FNMSysTermOfPMId  int     `gorm:"column:FNMSysTermOfPMId"`
	FNMSysCurId       int     `gorm:"column:FNMSysCurId"`
	FNExchangeRate    float64 `gorm:"column:FNExchangeRate"`
	FNMSysSaleId      int     `gorm:"column:FNMSysSaleId"`
	FTRemark          string  `gorm:"column:FTRemark"`
	FNSOAmt           float64 `gorm:"column:FNSOAmt"`
	// FNDisCountPer     float64 `gorm:"column:FNDisCountPer"`
	// FNDisCountAmt     float64 `gorm:"column:FNDisCountAmt"`
	// FNDisCountPer2      float64 `gorm:"column:FNDisCountPer2"`
	// FNDisCountAmt2      float64 `gorm:"column:FNDisCountAmt2"`
	// FNDisCountPer3      float64 `gorm:"column:FNDisCountPer3"`
	// FNDisCountAmt3      float64 `gorm:"column:FNDisCountAmt3"`
	FNVatAmt            float64 `gorm:"column:FNVatAmt"`
	FNSONetVatAmt       float64 `gorm:"column:FNSONetVatAmt"`
	FNSOReturnAmt       float64 `gorm:"column:FNSOReturnAmt"`
	FNSOGrandAmt        float64 `gorm:"column:FNSOGrandAmt"`
	FNDocumentState     int     `gorm:"column:FNDocumentState"`
	FNMSysShipId        int     `gorm:"column:FNMSysShipId"`
	FTStateApproved     string  `gorm:"column:FTStateApproved"`
	FTStateSendApproved string  `gorm:"column:FTStateSendApproved"`
	FTStateSendAppDate  string  `gorm:"column:FTStateSendAppDate"`
	FTStateSendAppTime  string  `gorm:"column:FTStateSendAppTime"`
}

func (*PreSaleOrder) TableName() string {
	return "DBWebApp.dbo.TMARTPreSaleOrder"
}
