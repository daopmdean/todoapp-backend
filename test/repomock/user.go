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
	// if username != "daopham" {
	// 	return nil
	// }

	// user := &entity.User{
	// 	Username: "daopham",
	// }
	// user.SetPassword("12345678")

	// return user
}

func NewUserRepoMockBuilder() *userRepoMockBuilder {
	return new(userRepoMockBuilder)
}

type userRepoMockBuilder struct {
	getUserByUsernameMock func(string) *entity.User
	saveUserMock          func(*entity.User) error
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
