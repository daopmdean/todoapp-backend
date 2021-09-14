package entity

import (
	"errors"
	"todoapp/internal/entity/internal"
)

type User struct {
	Username     string `gorm:"primaryKey"`
	Password     string
	PasswordSalt string
	FirstName    string
	LastName     string
}

func (u *User) SetPassword(password string) error {
	if len([]rune(password)) < 8 {
		return errors.New("password must be longer than 8")
	}
	//TODO: check password must have uppercase, lowercase, number, special character

	u.PasswordSalt = internal.Rand50Str()
	u.Password = internal.HashHMACSHA256(password, u.PasswordSalt)

	return nil
}

func (u *User) CheckPassword(password string) bool {
	passwordHashed := internal.HashHMACSHA256(password, u.PasswordSalt)
	if u.Password == passwordHashed {
		return true
	}

	return false
}
