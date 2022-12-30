package user

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"user/configs"
	"user/repo/user"
)

type ServiceImpl struct {
	userRepo                        user.Repo
	accessTokenKey, refreshTokenKey *rsa.PrivateKey
	refreshTokenPub                 *rsa.PublicKey
}

func New(
	userRepo user.Repo,
	config *configs.AppConfig,
) (Service, error) {
	priString := config.AccessTokenPrivate
	refreshPriString := config.RefreshTokenPrivate
	refreshPubString := config.RefreshTokenPublic

	//println("priString: " + priString)
	//println("refreshPriString: " + refreshPriString)
	//println("refreshPubString: " + refreshPubString)

	if config.Mode.IsLocal() {
		priString = strings.ReplaceAll(priString, "\\n", "\n")
		refreshPriString = strings.ReplaceAll(refreshPriString, "\\n", "\n")
		refreshPubString = strings.ReplaceAll(refreshPubString, "\\n", "\n")
	}

	accessTokenKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(priString))
	if err != nil {
		return nil, err
	}

	refreshTokenKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(refreshPriString))
	if err != nil {
		return nil, err
	}

	refreshTokenPub, err := jwt.ParseRSAPublicKeyFromPEM([]byte(refreshPubString))
	if err != nil {
		return nil, err
	}

	return &ServiceImpl{
		userRepo,
		accessTokenKey,
		refreshTokenKey,
		refreshTokenPub,
	}, nil
}
