package repomock

import "todoapp/internal/entity"

type userRepoMock struct {
	getUserByUsernameMock func(string) *entity.User
	saveUserMock          func(*entity.User) error
}

func (urm *userRepoMock) SaveUser(user *entity.User) error {
	return urm.saveUserMock(user)
}

func (urm *userRepoMock) GetUserByUsername(username string) *entity.User {
	return urm.getUserByUsernameMock(username)
}
