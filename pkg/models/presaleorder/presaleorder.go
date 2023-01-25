package presaleorder

type PreSaleOrderRequest struct {
	FTSaleOrderNo   string `json:"fTSaleOrderNo"`
	FTInsUser       string `json:"fTInsUser"`
	FDInsDate       string `json:"fDInsDate"`
	FTInsTime       string `json:"fTInsTime"`
	FTUpdUser       string `json:"fTUpdUser"`
	FDUpdDate       string `json:"fDUpdDate"`
	FTUpdTime       string `json:"fTUpdTime"`
	FDSaleOrderDate string `json:"fDSaleOrderDate"`
	FTSaleOrderBy   string `json:"fTSaleOrderBy"`
	FTQuotationNo   string `json:"fTQuotationNo"`
	// FNRevised           int     `gorm:"column:FNRevised"`
	// FNVatType           int     `gorm:"column:FNVatType"`
	// FNPriority          int     `gorm:"column:FNPriority"`
	FNMSysCustId      string `json:"fNMSysCustId"`
	FTDeliveryAddress string `json:"fTDeliveryAddress"`
	FDDeliveryDate    string `json:"fDDeliveryDate"`
	// FNMSysCrTermId      int     `gorm:"column:FNMSysCrTermId"`
	FNCreditDay      string `json:"fNCreditDay"`
	FNMSysTermOfPMId string `json:"fNMSysTermOfPMId"`
	// FNMSysCurId         int     `gorm:"column:FNMSysCurId"`
	// FNExchangeRate      float64 `gorm:"column:FNExchangeRate"`
	FNMSysSaleId string  `json:"fNMSysSaleId"`
	FTRemark     string  `json:"fTRemark"`
	FNSOAmt      float64 `json:"fNSOAmt"`
	// FNVatAmt            float64 `gorm:"column:FNVatAmt"`
	// FNSONetVatAmt       float64 `gorm:"column:FNSONetVatAmt"`
	// FNSOReturnAmt       float64 `gorm:"column:FNSOReturnAmt"`
	FNSOGrandAmt float64 `json:"fNSOGrandAmt"`
	// FNDocumentState     int     `gorm:"column:FNDocumentState"`
	FNMSysShipId        string `json:"fNMSysShipId"`
	FTStateApproved     string `json:"fTStateApproved"`
	FTStateSendApproved string `json:"fTStateSendApproved"`
	FTStateSendAppDate  string `json:"fTStateSendAppDate"`
	FTStateSendAppTime  string `json:"fTStateSendAppTime"`
}
