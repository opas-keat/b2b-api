package user

import (
	"context"
	"errors"
	"models/user"
	"strings"
	"time"
	"user/constant"
	"user/entities"
	"user/utils"
)

//func createAccessToken(userID string, role models.Role, expiresAt time.Time, jwtPrivateKey *rsa.PrivateKey) (accessToken, tokenID *string, err error) {
//	id := fmt.Sprintf("at-%v-%v:%v", userID, time.Now().UTC().Unix(), time.Now().UTC().UnixNano())
//	println(id)
//	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, entities.UserClaims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			Issuer:    "user-service",
//			ExpiresAt: jwt.NewNumericDate(expiresAt),
//			IssuedAt:  jwt.NewNumericDate(time.Now()),
//			ID:        id,
//		},
//		UserID: userID,
//		Role:   role,
//	})
//
//	token, err := claims.SignedString(jwtPrivateKey)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return pointer.ToString(token), pointer.ToString(id), nil
//}

//func createAccessToken(userID string, role string, expiresAt time.Time, jwtPrivateKey *rsa.PrivateKey) (accessToken, tokenID *string, err error) {
//	id := fmt.Sprintf("at-%v-%v:%v", userID, time.Now().UTC().Unix(), time.Now().UTC().UnixNano())
//	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, entities.UserClaims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			Issuer:    "user-service",
//			ExpiresAt: jwt.NewNumericDate(expiresAt),
//			IssuedAt:  jwt.NewNumericDate(time.Now()),
//			ID:        id,
//		},
//		UserID: userID,
//		Role:   role,
//	})
//
//	token, err := claims.SignedString(jwtPrivateKey)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return pointer.ToString(token), pointer.ToString(id), nil
//}
//
//func createRefreshToken(userID, accessTokenID string, expiresAt time.Time, jwtPrivateKey *rsa.PrivateKey) (accessToken, tokenID *string, err error) {
//	id := fmt.Sprintf("rft-%v:%v", time.Now().UTC().Unix(), time.Now().UTC().UnixNano())
//	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, entities.RefreshClaims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			Issuer:    "user-service",
//			ExpiresAt: jwt.NewNumericDate(expiresAt),
//			IssuedAt:  jwt.NewNumericDate(time.Now()),
//			ID:        id,
//		},
//		AccessTokenID: accessTokenID,
//		UserID:        userID,
//	})
//
//	token, err := claims.SignedString(jwtPrivateKey)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return pointer.ToString(token), pointer.ToString(id), nil
//}
//
//func (s ServiceImpl) createLoginToken(ctx context.Context, m entities.User) (accessToken, refreshToken *string, err error) {
//	var (
//		accessTokenExpiresAt  = time.Now().Add(10 * time.Minute)
//		refreshTokenExpiresAt = time.Now().Add(8 * time.Hour)
//		//role                  = models.MappingToRole[string(m.RoleName)]
//		role = m.RoleName
//	)
//
//	accessToken, accessTokenID, err := createAccessToken(m.ID, role, accessTokenExpiresAt, s.accessTokenKey)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	refreshToken, refreshTokenID, err := createRefreshToken(m.ID, pointer.GetString(accessTokenID), refreshTokenExpiresAt, s.refreshTokenKey)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	println(refreshTokenID)
//
//	return accessToken, refreshToken, nil
//}

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

	//accessToken, refreshToken, err := s.createLoginToken(ctx, *m)
	//if err != nil {
	//	return nil, err
	//}

	//var (
	//	accessTokenExpiresAt  = time.Now().Add(10 * time.Minute)
	//	refreshTokenExpiresAt = time.Now().Add(8 * time.Hour)
	//	//role                  = models.MappingToRole[string(m.RoleName)]
	//	role = m.RoleName
	//)

	token, err := utils.GenerateToken(60*time.Minute, m.ID, constant.PrivateKey)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &user.LoginResponse{
		AccessToken:  token,
		RefreshToken: token,
	}, nil
}
