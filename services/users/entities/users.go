package entities

import "time"

type User struct {
	Id        uint    `gorm:"column:id; primaryKey;"`
	Username  string  `gorm:"column:username; not null; unique" json:"username"`
	Password  string  `gorm:"column:pwd; not null" json:"password"`
	CreatedBy *string `gorm:"column:created_by"`
	CreatedAt time.Time
	UpdatedBy *string `gorm:"column:updated_by"`
	UpdatedAt time.Time
	DeletedBy *string `gorm:"column:deleted_by"`
	DeletedAt *time.Time
}

func (User) TableName() string {
	return "t_users"
}
