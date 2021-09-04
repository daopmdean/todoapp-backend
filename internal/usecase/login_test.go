package usecase_test

import (
	"testing"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestLoginSuccess(t *testing.T) {
	urm := repomock.NewUserRepoMockBuilder().Build()
	lu := usecase.NewLoginUsecase(urm)
	li := usecase.LoginInput{
		Username: "daopham",
		Password: "12345678",
	}

	_, err := lu.Login(li)

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestLoginFailWithInvalidUsername(t *testing.T) {
	urm := repomock.NewUserRepoMockBuilder().Build()
	lu := usecase.NewLoginUsecase(urm)
	li := usecase.LoginInput{
		Username: "daophama",
		Password: "12345678",
	}

	_, err := lu.Login(li)

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func TestLoginFailWithInvalidPassword(t *testing.T) {
	urm := repomock.NewUserRepoMockBuilder().Build()
	lu := usecase.NewLoginUsecase(urm)
	li := usecase.LoginInput{
		Username: "daopham",
		Password: "0000000",
	}

	_, err := lu.Login(li)

	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}
