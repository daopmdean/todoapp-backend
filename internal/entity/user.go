package entity

import (
	"errors"
	"todoapp/internal/entity/internal"
)

type User struct {
	Username     string `bson:"username,omitempty" gorm:"primaryKey"`
	Password     string `bson:"password,omitempty"`
	PasswordSalt string `bson:"passwordSalt,omitempty"`
	FirstName    string `bson:"firstName,omitempty"`
	LastName     string `bson:"lastName,omitempty"`
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
