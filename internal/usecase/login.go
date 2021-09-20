package usecase

import (
	"errors"
	"time"
	"todoapp/internal/usecase/repo"
	"todoapp/internal/usecase/util/access_token"
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

	token, err := access_token.CreateToken(user, time.Now())
	if err != nil {
		return nil, errors.New("cannot create jwt token: " + err.Error())
	}

	return &LoginOutput{
		Token: token,
	}, nil
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}
