package repo

import "todoapp/internal/entity"

type UserRepo interface {
	GetUserByUsername(username string) *entity.User
	CreateUser(*entity.User) error
}
