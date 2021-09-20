package access_token_test

import (
	"strings"
	"testing"
	"time"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/util/access_token"
)

func TestCreateToken(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}
	createTime := time.Unix(1632130727056, 0)
	_, err := access_token.CreateToken(&user, createTime)

	if err != nil {
		t.Errorf("expect nil err, got %v", err)
	}
}

func TestValidateToken(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}
	createTime := time.Unix(1632130727056, 0)
	verifyTime := createTime.Add(time.Hour*24*7 - time.Second)
	tokenStr, _ := access_token.CreateToken(&user, createTime)

	_, err := access_token.VerifyToken(tokenStr, verifyTime)

	if err != nil {
		t.Errorf("expect nil err, got %v", err)
	}
}

func TestValidateExpiredToken(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}
	createTime := time.Unix(1632130727056, 0)
	expiredTime := createTime.Add(time.Hour*24*7 + time.Second)
	tokenStr, _ := access_token.CreateToken(&user, createTime)

	_, err := access_token.VerifyToken(tokenStr, expiredTime)

	if err == nil {
		t.Errorf("expect token expired error, got nil error")
	}
	if err != nil && err.Error() != "token expired" {
		t.Errorf("expected token expired error, got %v", err)
	}
}

func TestValidateTokenWithInvalidSignature(t *testing.T) {
	user := entity.User{
		Username: "daopham",
	}
	createTime := time.Unix(1632130727056, 0)
	verifyTime := createTime.Add(time.Hour*24*7 - time.Second)
	tokenStr, _ := access_token.CreateToken(&user, createTime)
	tokenParts := strings.Split(tokenStr, ".")
	tokenParts[2] = "atv3hobh7t5v2zyfIlZJESVUPVHh8m5isHYxWDNTmwk"
	invalidSigToken := strings.Join(tokenParts, ".")

	_, err := access_token.VerifyToken(invalidSigToken, verifyTime)

	if err == nil {
		t.Errorf("expect err, got nil error")
	}
	if err != nil && err.Error() != "signature is invalid" {
		t.Errorf("expect signature is invalid, got %v", err)
	}
}
