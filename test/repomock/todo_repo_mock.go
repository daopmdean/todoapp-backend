package repomock

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

var _ repo.TodoRepo = new(todoRepoMock)

type todoRepoMock struct {
	getTodoListOfMock func(string) []entity.Todo
	createTodoMock    func(*entity.Todo) error
}

func (m *todoRepoMock) GetTodoListOf(username string) []entity.Todo {
	return m.getTodoListOfMock(username)
}

func (m *todoRepoMock) CreateTodo(todo *entity.Todo) error {
	return m.createTodoMock(todo)
}
