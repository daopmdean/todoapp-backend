package repomock

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewTodoRepoMockBuilder() *todoRepoMockBuilder {
	return &todoRepoMockBuilder{
		&todoRepoMock{},
	}
}

type todoRepoMockBuilder struct {
	todoRepoMock *todoRepoMock
}

func (mb *todoRepoMockBuilder) WithGetTodoListOfMock(mockFunc func(string) []entity.Todo) *todoRepoMockBuilder {
	mb.todoRepoMock.getTodoListOfMock = mockFunc
	return mb
}

func (mb *todoRepoMockBuilder) WithCreateTodoMock(mockFunc func(*entity.Todo) error) *todoRepoMockBuilder {
	mb.todoRepoMock.createTodoMock = mockFunc
	return mb
}

func (mb *todoRepoMockBuilder) Build() repo.TodoRepo {
	return mb.todoRepoMock
}
