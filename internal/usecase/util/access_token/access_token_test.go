package access_token_test

import (
	"testing"
	"time"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/util/access_token"
)

func TestCreateToken(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}

	_, err := access_token.CreateToken(&user, time.Now())

	if err != nil {
		t.Errorf("expect nil err, got %v", err)
	}
}

func TestParseToken(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}
	tokenStr, _ := access_token.CreateToken(&user, time.Now())

	_, err := access_token.ParseToken(tokenStr)

	if err != nil {
		t.Errorf("expect nil err, got %v", err)
	}
}

func TestExtractToken(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}
	tokenStr, _ := access_token.CreateToken(&user, time.Now())
	token, _ := access_token.ParseToken(tokenStr)

	tokenData, err := access_token.ExtractTokenClaims(token)

	if err != nil {
		t.Errorf("expect nil err, got %v", err)
	}

	if tokenData.Username != "daopham" {
		t.Errorf("expect username daopham, got %s", tokenData.Username)
	}
}
