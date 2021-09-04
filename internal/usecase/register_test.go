package usecase_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
)

type userRepoMock struct {
}

func (urm *userRepoMock) SaveUser(user *entity.User) error {
	return nil
}

func (urm *userRepoMock) GetUserByUsername(username string) *entity.User {
	if username != "daopham" {
		return nil
	}

	user := &entity.User{
		Username: "daopham",
	}
	user.SetPassword("12345678")

	return user
}

func TestRegisterSuccess(t *testing.T) {
	urm := new(userRepoMock)
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "someone",
		Password:  "12345678",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}

func TestRegisterFailWhenPasswordLessThan8Chars(t *testing.T) {
	urm := new(userRepoMock)
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "someone",
		Password:  "123456",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got: %v", err)
	}
}

func TestRegisterFailWhenUsernameAndPasswordEmpty(t *testing.T) {
	urm := new(userRepoMock)
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "",
		Password:  "",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got: %v", err)
	}
}

func TestRegisterFailWhenUsernameEmpty(t *testing.T) {
	urm := new(userRepoMock)
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "",
		Password:  "12345678",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got: %v", err)
	}
}

func TestRegisterFailWhenPasswordEmpty(t *testing.T) {
	urm := new(userRepoMock)
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "someone",
		Password:  "",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got: %v", err)
	}
}
