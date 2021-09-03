package repo

import "todoapp/internal/entity"

type UserRepo interface {
	SaveUser(*entity.User) error
}
