package memdb

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewUserRepo() repo.UserRepo {
	dao := entity.User{Username: "daopham", FirstName: "Dao", LastName: "Pham"}
	dao.SetPassword("12345678")
	hung := entity.User{Username: "hungpham", FirstName: "Hung", LastName: "Pham"}
	hung.SetPassword("87654321")

	return &userRepo{
		users: []entity.User{dao, hung},
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
