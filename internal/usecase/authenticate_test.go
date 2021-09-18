package usecase_test

import (
	"testing"
	"time"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/internal/usecase/util/access_token"
)

func TestEmptyAccessToken(t *testing.T) {
	au := usecase.NewAuthenticationUsecase()

	_, err := au.AuthenticateWebAccessToken("")

	if err == nil {
		t.Errorf("expected empty access token error, got no error")
	}
}

func TestAccessToken(t *testing.T) {
	user := &entity.User{
		Username: "daopham",
	}
	token, _ := access_token.CreateToken(user, time.Now())
	au := usecase.NewAuthenticationUsecase()

	username, err := au.AuthenticateWebAccessToken(token)

	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	if err == nil && username != "daopham" {
		t.Errorf("expected daopham, got %s", username)
	}
}
