package dealer

type DealerResponse struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type ListDealerRequest struct {
	FTCustCode string `json:"dealer_code"`
	//GeneralSearch *string `json:"general_search"`
	//SortBy        *string `json:"sort_by" validate:"omitempty,oneof=member_id balance last_login withdraw deposit"`
	//SortType      *string `json:"sort_type" validate:"omitempty,oneof=ASC DESC"`
}
