package usecase

import "todoapp/internal/usecase/util/access_token"

func NewAuthenticationUsecase() *AuthenticationUsecase {
	return &AuthenticationUsecase{}
}

type AuthenticationUsecase struct {
}

func (au *AuthenticationUsecase) AuthenticateWebAccessToken(accessToken string) (string, error) {
	token, err := access_token.ParseToken(accessToken)
	if err != nil {
		return "", err
	}

	data, err := access_token.ExtractTokenClaims(token)
	if err != nil {
		return "", err
	}

	return data.Username, nil
}
