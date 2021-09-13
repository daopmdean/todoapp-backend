package repomock

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

var _ repo.TodoRepo = new(todoRepoMock)

type todoRepoMock struct {
	getTodoListOfMock func(string) []entity.Todo
	getTodoByID       func(string) *entity.Todo
	createTodoMock    func(*entity.Todo) error
	deleteTodoMock    func(string) error
	updateTodoMock    func(*entity.Todo) error
}

func (m *todoRepoMock) GetTodoListOf(username string) []entity.Todo {
	return m.getTodoListOfMock(username)
}

func (m *todoRepoMock) GetTodoByID(todoID string) *entity.Todo {
	return m.getTodoByID(todoID)
}

func (m *todoRepoMock) CreateTodo(todo *entity.Todo) error {
	return m.createTodoMock(todo)
}

func (m *todoRepoMock) DeleteTodo(todoID string) error {
	return m.deleteTodoMock(todoID)
}

func (m *todoRepoMock) UpdateTodo(todo *entity.Todo) error {
	return m.updateTodoMock(todo)
}
