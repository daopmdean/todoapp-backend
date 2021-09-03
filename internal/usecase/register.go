package usecase

import (
	"errors"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewRegisterUsecase(ur repo.UserRepo) *RegisterUsecase {
	return &RegisterUsecase{userRepo: ur}
}

type RegisterUsecase struct {
	userRepo repo.UserRepo
}

type RegisterInput struct {
	//TODO: fix warning
	Username  string `json: "username"`
	Password  string `json: "password"`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}

func (i *RegisterInput) validate() error {
	// TODO: check white space

	if len(i.Username) == 0 {
		return errors.New("username must not be empty")
	}
	if len(i.Password) == 0 {
		return errors.New("password must not be empty")
	}
	return nil
}

func (ru *RegisterUsecase) Register(input RegisterInput) error {
	if err := input.validate(); err != nil {
		return errors.New("register input invalid, got error: " + err.Error())
	}

	user := &entity.User{
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err := user.SetPassword(input.Password)
	if err != nil {
		return errors.New("cannot set user password: " + err.Error())
	}

	err = ru.userRepo.SaveUser(user)
	if err != nil {
		return errors.New("cannot save user: " + err.Error())
	}

	return nil
}
