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

func (mb *todoRepoMockBuilder) WithGetTodoByID(mockFunc func(string) *entity.Todo) *todoRepoMockBuilder {
	mb.todoRepoMock.getTodoByIDMock = mockFunc
	return mb
}

func (mb *todoRepoMockBuilder) WithCreateTodoMock(mockFunc func(*entity.Todo) error) *todoRepoMockBuilder {
	mb.todoRepoMock.createTodoMock = mockFunc
	return mb
}

func (mb *todoRepoMockBuilder) WithDeleteTodoMock(mockFunc func(string) error) *todoRepoMockBuilder {
	mb.todoRepoMock.deleteTodoMock = mockFunc
	return mb
}

func (mb *todoRepoMockBuilder) WithUpdateTodoMock(mockFunc func(*entity.Todo) error) *todoRepoMockBuilder {
	mb.todoRepoMock.updateTodoMock = mockFunc
	return mb
}

func (mb *todoRepoMockBuilder) Build() repo.TodoRepo {
	return mb.todoRepoMock
}
