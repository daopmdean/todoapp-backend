package usecase_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/internal/usecase/repo"
	"todoapp/test/repomock"
)

func createUserRepoMockForRegister() repo.UserRepo {
	return repomock.NewUserRepoMockBuilder().
		WithCreateUserMock(func(u *entity.User) error {
			return nil
		}).
		Build()
}

func TestRegisterSuccess(t *testing.T) {
	urm := createUserRepoMockForRegister()
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
	urm := createUserRepoMockForRegister()
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "someone",
		Password:  "123456",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got nil error")
	}
}

func TestRegisterFailWhenUsernameAndPasswordEmpty(t *testing.T) {
	urm := createUserRepoMockForRegister()
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "",
		Password:  "",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got nil error")
	}
}

func TestRegisterFailWhenUsernameEmpty(t *testing.T) {
	urm := createUserRepoMockForRegister()
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "",
		Password:  "12345678",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got nil error")
	}
}

func TestRegisterFailWhenPasswordEmpty(t *testing.T) {
	urm := createUserRepoMockForRegister()
	ru := usecase.NewRegisterUsecase(urm)
	u := usecase.RegisterInput{
		Username:  "someone",
		Password:  "",
		FirstName: "Dao",
		LastName:  "Pham",
	}

	err := ru.Register(u)

	if err == nil {
		t.Errorf("expected error, got nil error")
	}
}
