package entity

import "errors"

type User struct {
	Username     string
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

	u.Password = password
	//TODO: Hash password

	return nil
}
