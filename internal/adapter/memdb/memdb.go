package memdb

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewUserRepo() repo.UserRepo {
	return &userRepo{
		users: []entity.User{},
	}
}

type userRepo struct {
	users []entity.User
}

func (ur *userRepo) SaveUser(user *entity.User) error {
	ur.users = append(ur.users, *user)
	return nil
}
