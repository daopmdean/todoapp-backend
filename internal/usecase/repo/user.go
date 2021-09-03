package repo

import "todoapp/internal/entity"

type UserRepo interface {
	GetUserByUsername(username string) *entity.User
	SaveUser(*entity.User) error
}
