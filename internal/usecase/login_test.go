package usecase_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/internal/usecase/repo"
	"todoapp/test/repomock"
)

func createUserRepoMockWithData() repo.UserRepo {
	return repomock.NewUserRepoMockBuilder().
		WithGetUserByUsernameMock(func(username string) *entity.User {
			if username != "daopham" {
				return nil
			}

			user := &entity.User{
				Username: "daopham",
			}
			user.SetPassword("12345678")

			return user
		}).
		Build()
}

func TestLoginSuccess(t *testing.T) {
	urm := createUserRepoMockWithData()
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
	urm := createUserRepoMockWithData()
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
	urm := createUserRepoMockWithData()
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
