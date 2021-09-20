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
	verifyTime := time.Unix(1632130727056, 0)
	_, err := au.AuthenticateWebAccessToken("", verifyTime)

	if err == nil {
		t.Errorf("expected error, got no error")
	}
}

func TestAccessToken(t *testing.T) {
	user := &entity.User{
		Username: "daopham",
	}
	au := usecase.NewAuthenticationUsecase()
	createTime := time.Unix(1632130727056, 0)
	verifyTime := createTime.Add(time.Hour*24*7 - time.Second)
	token, _ := access_token.CreateToken(user, createTime)

	username, err := au.AuthenticateWebAccessToken(token, verifyTime)

	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}
	if err == nil && username != "daopham" {
		t.Errorf("expected daopham, got %s", username)
	}
}
