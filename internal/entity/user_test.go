package entity_test

import (
	"testing"
	"todoapp/internal/entity"
)

// Test set new password with less than 8 characters
func TestUserSetPasswordLessThan8Chars(t *testing.T) {
	user := new(entity.User)
	password := "123456"

	err := user.SetPassword(password)

	if err == nil {
		t.Errorf("expected error")
	}
}

// Test set new password with 8 characters
func TestUserSetPasswordWith8Chars(t *testing.T) {
	user := entity.User{Username: "someone"}
	password := "12345678"

	err := user.SetPassword(password)

	if err != nil {
		t.Errorf("expected nil error")
	}
}

// Test set new password with more than 8 characters
func TestUserSetPasswordMoreThan8Chars(t *testing.T) {
	user := entity.User{Username: "someone"}
	password := "123456789"

	err := user.SetPassword(password)

	if err != nil {
		t.Errorf("expected nil error")
	}
}

func TestUserSetPasswordSuccessWithHash(t *testing.T) {
	user := entity.User{Username: "someone"}
	password := "123456789"

	_ = user.SetPassword(password)

	if user.Password == password {
		t.Errorf("expected password hashed")
	}
}

func TestUserCheckWrongPassword(t *testing.T) {
	user := entity.User{Username: "someone"}
	password := "123456789"
	user.SetPassword(password)
	wrongPassword := "0000000"

	result := user.CheckPassword(wrongPassword)

	if result {
		t.Errorf("expected false, got true")
	}
}

func TestUserCheckCorrectPassword(t *testing.T) {
	user := entity.User{Username: "someone"}
	password := "123456789"
	user.SetPassword(password)

	result := user.CheckPassword(password)

	if !result {
		t.Errorf("expected true, got false")
	}
}
