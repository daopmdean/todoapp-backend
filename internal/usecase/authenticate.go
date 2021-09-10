package usecase

import "errors"

func NewAuthenticationUsecase() *AuthenticationUsecase {
	return &AuthenticationUsecase{}
}

type AuthenticationUsecase struct {
}

func (au *AuthenticationUsecase) AuthenticateWebAccessToken(accessToken string) (username string, err error) {
	//TODO: implement jwt
	if accessToken == "" {
		return "", errors.New("empty access token")
	}
	return accessToken, nil
}
