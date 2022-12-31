package user

import (
	"context"
	"errors"
	"member/constant"
	"member/entities"
	"member/utils"
	"models/user"
	"strings"
	"time"
)

func (s ServiceImpl) Login(ctx context.Context, username, password string) (*user.LoginResponse, error) {
	username = strings.ToLower(username)
	m, err := s.userRepo.Get(ctx, entities.User{Email: username})
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !m.Verified {
		return nil, errors.New("please verify your email")
	}
	if !utils.CheckPasswordHash(password, m.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}
	println(m.DealerCode)
	dc := "CL1713"
	dealerResp, err := s.dealerConnector.GetDealer(ctx, dc)
	if err != nil {
		return nil, err
	}
	println(dealerResp.ID)
	println(dealerResp.Code)
	println(dealerResp.Name)
	println(dealerResp.Phone)
	println(dealerResp.Address)
	//dealerListResp, err := s.dealerConnector.FindDealers(ctx, dc)
	//if err != nil {
	//	return nil, err
	//}
	//println(dealerListResp)
	token, err := utils.GenerateToken(60*time.Minute, m.ID, constant.PrivateKey)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user.LoginResponse{
		AccessToken:  token,
		RefreshToken: token,
	}, nil
}
