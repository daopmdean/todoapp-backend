package usecase

import (
	"errors"
	"todoapp/internal/usecase/repo"
)

func NewGetMeUsecase(ur repo.UserRepo, username string) *GetMeUsecase {
	return &GetMeUsecase{ur, username}
}

type GetMeUsecase struct {
	userRepo repo.UserRepo
	username string
}

func (gmu *GetMeUsecase) GetMe() (*GetMeOutput, error) {
	user := gmu.userRepo.GetUserByUsername(gmu.username)
	if user == nil {
		return nil, errors.New("cannot get user info")
	}

	return &GetMeOutput{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

type GetMeOutput struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
