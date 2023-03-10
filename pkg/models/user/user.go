package user

import "time"

type RegisterUserRequest struct {
	//Username string `json:"username" validate:"required,min=8,max=20,username"`
	Email      string `json:"email" validate:"required,min=8,max=100,email"`
	Password   string `json:"password" validate:"required,min=6,max=32,password"`
	DealerCode string `json:"dealer_code" validate:"required,min=6,max=32"`
	//Phone        string  `json:"phone" validate:"required,len=10,phone"`
	//PhoneConfirm bool    `json:"phone_confirm"`
	//Status       *bool   `json:"status"`
	//StatusRemark *string `json:"status_remark,omitempty"`
	//Birthdate    *string `json:"birthdate"`
}

type CreateUserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	//LoginResponse
}
