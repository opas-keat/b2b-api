package product

import "time"

type SystemLinkProductResponse struct {
	ID              string  `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	MatSize         string  `json:"mat_size"`
	Color           string  `json:"color"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Width           string  `json:"width"`
	Offset          string  `json:"offset"`
	TreadWare       string  `json:"tread_ware"`
	PitchCircleCode string  `json:"pitch_circle_code"`
	Price           float64 `json:"price"`
	DealerPrice1    float64 `json:"dealer_price_1"`
	GroupCode       string  `json:"group_code"`
}

type BrandAndModelResponse struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type ListBrandAndModelRequest struct {
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	ProductType string `query:"product_type" json:"product_type"`
}

type ProductRequest struct {
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	MatSize         string  `json:"mat_size"`
	Color           string  `json:"color"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Width           string  `json:"width"`
	Offset          string  `json:"offset"`
	TreadWare       string  `json:"tread_ware"`
	PitchCircleCode string  `json:"pitch_circle_code"`
	Price           float64 `json:"price"`
	DealerPrice1    float64 `json:"dealer_price_1"`
	GroupCode       string  `json:"group_code"`
}

type CreateProductRequest struct {
	LinkId          string `json:"link_id"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	MatSize         string `json:"mat_size"`
	Color           string `json:"color"`
	Brand           string `json:"brand"`
	Model           string `json:"model"`
	Width           string `json:"width"`
	Offset          string `json:"offset"`
	TreadWare       string `json:"tread_ware"`
	PitchCircleCode string `json:"pitch_circle_code"`
	Price           string `json:"price"`
	DealerPrice1    string `json:"dealer_price_1"`
	GroupCode       string `json:"group_code"`
}

type CreateProductResponse struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Model           string    `json:"model"`
	OffsetName      string    `json:"offset_name"`
	Name            string    `json:"name"`
	MatSizeName     string    `json:"mat_size_name"`
	Color           string    `json:"color"`
	Brand           string    `json:"brand"`
	Code            string    `json:"code"`
	ProdGrpCode     string    `json:"prod_grp_code"`
	PitchCircleCode string    `json:"pitch_circle_code"`
	Treadware       string    `json:"tread_ware"`
	WidthName       string    `json:"width_name"`
	LinkId          string    `json:"link_id;"`
	Price           float64   `json:"price"`
	DealerPrice     float64   `json:"dealer_price"`
	QuantityBal     float64   `json:"quantity_balance"`
	QuantityBalB    float64   `json:"quantity_balance_b"`
}
