package users

import "time"

type CreateMemberRequest struct {
	Username     string  `json:"username" validate:"required,min=8,max=20,username"`
	Password     string  `json:"password" validate:"required,min=6,max=32,password"`
	Phone        string  `json:"phone" validate:"required,len=10,phone"`
	PhoneConfirm bool    `json:"phone_confirm"`
	Status       *bool   `json:"status"`
	StatusRemark *string `json:"status_remark,omitempty"`
	Birthdate    *string `json:"birthdate"`
}

type CreateMemberResponse struct {
	MemberID  string    `json:"member_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
