package user

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/jinzhu/copier"
	shareModels "models/status_code"
	models "models/user"
	"shareerrors"
	"strings"
	"user/constant"
	"user/entities"
)

func (s ServiceImpl) Register(ctx context.Context, req models.RegisterUserRequest) (*models.CreateUserResponse, error) {
	req.Email = strings.ToLower(req.Email)

	_, err := s.userRepo.Get(ctx, entities.User{Email: req.Email})
	if !shareerrors.IsCode(err, shareModels.NotFound) { // Found
		return nil, shareerrors.NewError(shareModels.Duplicate, "email is duplicate")
	}

	mQuery, err := s.userRepo.Create(ctx, entities.User{
		Email:        req.Email,
		PasswordHash: req.Password,
		DealerCode:   req.DealerCode,
		CreatedBy:    pointer.ToString(constant.SystemUID),
		UpdatedBy:    pointer.ToString(constant.SystemUID),
		RoleName:     "USER",
	})
	if err != nil {
		return nil, err
	}

	resp := new(models.CreateUserResponse)
	if err := copier.Copy(resp, mQuery); err != nil {
		return nil, err
	}

	//accessToken, refreshToken, err := s.createLoginToken(ctx, *mQuery, constant.PrivateKey)
	//if err != nil {
	//	return nil, err
	//}
	//resp.LoginResponse.AccessToken = pointer.GetString(accessToken)
	//resp.LoginResponse.RefreshToken = pointer.GetString(refreshToken)
	return resp, nil
}
