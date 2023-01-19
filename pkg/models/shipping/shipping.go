package shipping

type Shipping struct {
	ID     string `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Tel    string `json:"phone"`
	Mobile string `json:"mobile"`
	Note   string `json:"note"`
}

type ListShippingRequest struct {
	FTShipCode           string `json:"code"`
	FTShipNameTH         string `json:"name"`
	FTCentralRegion      string `json:"centralRegion"`
	FTTheNorthEastRegion string `json:"theNorthEastRegion"`
	FTNorthRegion        string `json:"northRegion"`
	FTWestRegion         string `json:"westRegion"`
	FTSouthernRegion     string `json:"southernRegion"`
	FTEastRegion         string `json:"eastRegion"`
	FTBKKRegion          string `json:"bkkRegion"`
	//GeneralSearch *string `json:"general_search"`
	//SortBy        *string `json:"sort_by" validate:"omitempty,oneof=member_id balance last_login withdraw deposit"`
	//SortType      *string `json:"sort_type" validate:"omitempty,oneof=ASC DESC"`
}
