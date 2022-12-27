package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"user/utils"
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
	return string("t_users")
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

//type RoleName string
//
//func (e RoleName) String() string {
//	return string(e)
//}
//
//func ConvertToDBRole(role models.Role) RoleName {
//	switch role {
//	case models.Admin:
//		return AdminRole
//	case models.AdminAssistant:
//		return AdminAssistantRole
//	default: // User
//		return UserRole
//	}
//}
//
//const (
//	AdminRole          RoleName = "ADMIN"
//	AdminAssistantRole RoleName = "ADMIN_ASSISTANT"
//	UserRole           RoleName = "USER"
//)
//
//// Scan is ...
//func (e RoleName) Scan(value interface{}) error {
//	val, _ := value.(string)
//	*e = RoleName(val)
//	return nil
//}
//
//// Value is ...
//func (e RoleName) Value() (driver.Value, error) {
//	return string(e), nil
//}

//type SignUpInput struct {
//	Name            string `json:"name" binding:"required"`
//	Email           string `json:"email" binding:"required"`
//	Password        string `json:"password" binding:"required,min=8"`
//	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
//	Photo           string `json:"photo" binding:"required"`
//}
//
//type SignInInput struct {
//	Email    string `json:"email"  binding:"required"`
//	Password string `json:"password"  binding:"required"`
//}
//
//type UserResponse struct {
//	ID         string    `json:"id,omitempty"`
//	Name       string    `json:"name,omitempty"`
//	Email      string    `json:"email,omitempty"`
//	DealerCode string    `json:"dealer_code,omitempty"`
//	Role       string    `json:"role,omitempty"`
//	Photo      string    `json:"photo,omitempty"`
//	Provider   string    `json:"provider"`
//	CreatedAt  time.Time `json:"created_at"`
//	UpdatedAt  time.Time `json:"updated_at"`
//}
