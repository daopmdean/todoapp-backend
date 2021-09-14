package mssqldb

import (
	"errors"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"gorm.io/gorm"
)

// check if missing implementation
var _ repo.UserRepo = new(userRepo)

func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}

func (ur *userRepo) CreateUser(user *entity.User) error {
	result := ur.db.Create(user)
	if result.Error != nil {
		return errors.New("cannot create user with gorm: " + result.Error.Error())
	}
	return nil
}

func (ur *userRepo) GetUserByUsername(username string) *entity.User {
	var user entity.User
	ur.db.First(&user, username)
	return &user
}
