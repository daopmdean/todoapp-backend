package repomock

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

var _ repo.UserRepo = new(userRepoMock)

type userRepoMock struct {
	getUserByUsernameMock func(string) *entity.User
	createUserMock        func(*entity.User) error
}

func (urm *userRepoMock) CreateUser(user *entity.User) error {
	return urm.createUserMock(user)
}

func (urm *userRepoMock) GetUserByUsername(username string) *entity.User {
	return urm.getUserByUsernameMock(username)
}
