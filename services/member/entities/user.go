package entities

import (
	"b2b/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CreatedBy        *string `gorm:"column:created_by"`
	UpdatedBy        *string `gorm:"column:updated_by"`
	ID               string  `gorm:"type:varchar(36);primary_key"`
	Email            string  `gorm:"uniqueIndex;not null"`
	DealerCode       string  `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash     string  `gorm:"column:password_hash;not null"`
	RoleName         string  `gorm:"column:role_name;not null"`
	VerificationCode string  `gorm:"column:verification_code"`
	Verified         bool    `gorm:"column:verified;not null"`
}

func (*User) TableName() string {
	return "t_users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}

	hashPassword, err := utils.HashPassword(u.PasswordHash)
	if err != nil {
		return err
	}
	u.PasswordHash = hashPassword
	u.CreatedAt = time.Now()

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.PasswordHash != "" {
		hashPassword, err := utils.HashPassword(u.PasswordHash)
		if err != nil {
			return err
		}
		u.PasswordHash = hashPassword
	}
	return nil
}
