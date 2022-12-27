package user

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/jinzhu/copier"
	"github.com/thanhpk/randstr"
	"models/email"
	shareModels "models/status_code"
	models "models/user"
	"shareerrors"
	"strings"
	"user/constant"
	"user/entities"
	"user/utils"
)

func (s ServiceImpl) Register(ctx context.Context, req models.RegisterUserRequest) (*models.CreateUserResponse, error) {
	req.Email = strings.ToLower(req.Email)

	_, err := s.userRepo.Get(ctx, entities.User{Email: req.Email})
	if !shareerrors.IsCode(err, shareModels.NotFound) { // Found
		return nil, shareerrors.NewError(shareModels.Duplicate, "email is duplicate")
	}

	newUser, err := s.userRepo.Create(ctx, entities.User{
		Email:        req.Email,
		PasswordHash: req.Password,
		DealerCode:   req.DealerCode,
		CreatedBy:    pointer.ToString(constant.SystemUID),
		UpdatedBy:    pointer.ToString(constant.SystemUID),
		RoleName:     "user",
		Verified:     false,
	})
	if err != nil {
		return nil, err
	}

	resp := new(models.CreateUserResponse)
	if err := copier.Copy(resp, newUser); err != nil {
		return nil, err
	}

	//accessToken, refreshToken, err := s.createLoginToken(ctx, *mQuery, constant.PrivateKey)
	//if err != nil {
	//	return nil, err
	//}
	//resp.LoginResponse.AccessToken = pointer.GetString(accessToken)
	//resp.LoginResponse.RefreshToken = pointer.GetString(refreshToken)

	// Generate Verification Code
	code := randstr.String(20)

	verificationCode := utils.Encode(code)
	em := entities.User{VerificationCode: verificationCode}
	if err := s.userRepo.Update(ctx, newUser.ID, em); err != nil {
		//return resp, err
	}

	// ? Send Email
	emailData := email.Email{
		URL:       constant.ClientOrigin + "/api/v1/verifyemail/" + code,
		FirstName: "Hello",
		Subject:   "Your account verification code",
	}

	utils.SendEmail(&emailData)
	return resp, nil
}
