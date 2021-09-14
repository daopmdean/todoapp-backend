package repomock

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewUserRepoMockBuilder() *userRepoMockBuilder {
	return &userRepoMockBuilder{
		&userRepoMock{},
	}
}

type userRepoMockBuilder struct {
	userRepoMock *userRepoMock
}

func (urmb *userRepoMockBuilder) WithCreateUserMock(mockFunc func(*entity.User) error) *userRepoMockBuilder {
	urmb.userRepoMock.createUserMock = mockFunc
	return urmb
}

func (urmb *userRepoMockBuilder) WithGetUserByUsernameMock(mockFunc func(string) *entity.User) *userRepoMockBuilder {
	urmb.userRepoMock.getUserByUsernameMock = mockFunc
	return urmb
}

func (urmb *userRepoMockBuilder) Build() repo.UserRepo {
	return urmb.userRepoMock
}
