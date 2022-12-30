package entities

import (
	"github.com/google/uuid"
	"time"
)

type Dealer struct {
	ID         string `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  *string `gorm:"column:created_by"`
	UpdatedBy  *string `gorm:"column:updated_by"`
	Name       string  `gorm:"column:name"`
	Address    string  `gorm:"column:address"`
	Email      string  `gorm:"column:email"`
	Phone      string  `gorm:"column:phone"`
	DealerCode string  `gorm:"column:dealer_code;uniqueIndex;not null"`
}

func (*Dealer) TableName() string {
	return "t_dealers"
}

func (d *Dealer) BeforeCreate() (err error) {
	if d.ID == "" {
		d.ID = uuid.NewString()
	}
	d.CreatedAt = time.Now()

	return nil
}
