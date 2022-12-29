package product

type ProductResponse struct {
	ID              string  `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	MatSize         string  `json:"mat_size"`
	Color           string  `json:"color"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Width           string  `json:"width"`
	Offset          string  `json:"offset"`
	ThreadWare      string  `json:"thread_ware"`
	PitchCircleCode string  `json:"pitch_circle_code"`
	Price           float64 `json:"price"`
	DealerPrice1    float64 `json:"dealer_price_1"`
}

type BrandAndModelResponse struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}
