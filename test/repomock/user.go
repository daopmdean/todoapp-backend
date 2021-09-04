package repomock

import "todoapp/internal/entity"

type userRepoMock struct {
}

func (urm *userRepoMock) SaveUser(user *entity.User) error {
	return nil
}

func (urm *userRepoMock) GetUserByUsername(username string) *entity.User {
	if username != "daopham" {
		return nil
	}

	user := &entity.User{
		Username: "daopham",
	}
	user.SetPassword("12345678")

	return user
}

func NewUserRepoMockBuilder() *userRepoMockBuilder {
	return new(userRepoMockBuilder)
}

type userRepoMockBuilder struct {
}

func (urmb *userRepoMockBuilder) WithUser(user *entity.User) *userRepoMockBuilder {
	return urmb
}

func (urmb *userRepoMockBuilder) Build() *userRepoMock {
	return new(userRepoMock)
}
