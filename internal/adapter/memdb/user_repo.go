package memdb

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

// check if missing implementation
var _ repo.UserRepo = new(userRepo)

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

func (ur *userRepo) GetUserByUsername(username string) *entity.User {
	for _, user := range ur.users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}
