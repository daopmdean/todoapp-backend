package usecase_test

import (
	"testing"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestRegisterSuccess(t *testing.T) {
	urm := repomock.NewUserRepoMockBuilder().Build()
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
	urm := repomock.NewUserRepoMockBuilder().Build()
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
	urm := repomock.NewUserRepoMockBuilder().Build()
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
	urm := repomock.NewUserRepoMockBuilder().Build()
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
	urm := repomock.NewUserRepoMockBuilder().Build()
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
