package usecase

import (
	"time"
	"todoapp/internal/usecase/util/access_token"
)

func NewAuthenticationUsecase() *AuthenticationUsecase {
	return &AuthenticationUsecase{}
}

type AuthenticationUsecase struct {
}

func (au *AuthenticationUsecase) AuthenticateWebAccessToken(accessToken string, verifyTime time.Time) (username string, err error) {
	data, err := access_token.VerifyToken(accessToken, verifyTime)
	if err != nil {
		return "", err
	}

	return data.Username, nil
}
