package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       *string `gorm:"column:created_by"`
	UpdatedBy       *string `gorm:"column:updated_by"`
	ID              string  `gorm:"type:varchar(36);primary_key"`
	Model           string  `gorm:"column:model"`
	OffsetName      string  `gorm:"column:offset_name"`
	Name            string  `gorm:"column:name"`
	MatSize         string  `gorm:"column:mat_size"`
	Color           string  `gorm:"column:color"`
	Brand           string  `gorm:"column:brand"`
	Code            string  `gorm:"column:code"`
	ProdGrpCode     string  `gorm:"column:prod_grp_code"`
	PitchCircleCode string  `gorm:"column:pitch_circle_code"`
	Treadware       string  `gorm:"column:tread_ware"`
	WidthName       string  `gorm:"column:width_name"`
	LinkId          string  `gorm:"column:link_id;unique;not null;"`
	Price           float64 `gorm:"column:price"`
	DealerPrice     float64 `gorm:"column:dealer_price"`
	QuantityBal     float64 `gorm:"column:quantity_balance"`
	QuantityBalB    float64 `gorm:"column:quantity_balance_b"`
}

func (*Product) TableName() string {
	return "t_products"
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	p.CreatedAt = time.Now()

	return nil
}

type BrandAndModel struct {
	Brand string `gorm:"column:brand"`
	Model string `gorm:"column:model"`
}
