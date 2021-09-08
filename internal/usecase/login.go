package usecase

import (
	"errors"
	"todoapp/internal/usecase/repo"
)

func NewLoginUsecase(ur repo.UserRepo) *LoginUsecase {
	return &LoginUsecase{ur}
}

type LoginUsecase struct {
	userRepo repo.UserRepo
}

func (lu *LoginUsecase) Login(input LoginInput) (*LoginOutput, error) {
	user := lu.userRepo.GetUserByUsername(input.Username)

	if user == nil {
		return nil, errors.New("username or password incorrect")
	}

	isPasswordCorrect := user.CheckPassword(input.Password)
	if !isPasswordCorrect {
		return nil, errors.New("username or password incorrect")
	}

	return &LoginOutput{
		Token: user.Username,
	}, nil
}

type LoginInput struct {
	//TODO: fix warning
	Username string `json: "username"`
	Password string `json: "password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}
