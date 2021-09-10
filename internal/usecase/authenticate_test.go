package usecase_test

import (
	"testing"
	"todoapp/internal/usecase"
)

func TestEmptyAccessToken(t *testing.T) {
	au := usecase.NewAuthenticationUsecase()

	_, err := au.AuthenticateWebAccessToken("")

	if err == nil {
		t.Errorf("expected empty access token error, got no error")
	}
	if err != nil && err.Error() != "empty access token" {
		t.Errorf("expected empty access token error, got %s", err.Error())
	}
}

func TestAccessToken(t *testing.T) {
	au := usecase.NewAuthenticationUsecase()

	username, err := au.AuthenticateWebAccessToken("daopham")

	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	if err == nil && username != "daopham" {
		t.Errorf("expected daopham, got %s", username)
	}
}
