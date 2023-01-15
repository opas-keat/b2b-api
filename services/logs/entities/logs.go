package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Logs struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	CreatedBy *string `gorm:"column:created_by"`
	Detail    string  `gorm:"column:detail"`
}

func (*Logs) TableName() string {
	return "t_logs"
}

func (d *Logs) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		d.ID = uuid.NewString()
	}
	d.CreatedAt = time.Now()

	return nil
}
