package repomock

import "todoapp/internal/entity"

type userRepoMockBuilder struct {
	getUserByUsernameMock func(string) *entity.User
	saveUserMock          func(*entity.User) error
}

func NewUserRepoMockBuilder() *userRepoMockBuilder {
	return new(userRepoMockBuilder)
}

func (urmb *userRepoMockBuilder) WithSaveUserMock(mockFunc func(*entity.User) error) *userRepoMockBuilder {
	urmb.saveUserMock = mockFunc
	return urmb
}

func (urmb *userRepoMockBuilder) WithGetUserByUsernameMock(mockFunc func(string) *entity.User) *userRepoMockBuilder {
	urmb.getUserByUsernameMock = mockFunc
	return urmb
}

func (urmb *userRepoMockBuilder) Build() *userRepoMock {
	return &userRepoMock{
		saveUserMock:          urmb.saveUserMock,
		getUserByUsernameMock: urmb.getUserByUsernameMock,
	}
}