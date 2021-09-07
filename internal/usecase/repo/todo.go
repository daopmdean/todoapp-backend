package repo

import "todoapp/internal/entity"

type TodoRepo interface {
	GetTodoListOf(username string) []entity.Todo
	CreateTodo(*entity.Todo) error
	// CheckTodoItem([]entity.TodoItem, *entity.TodoItem) error
	// DeleteTodoItem([]entity.TodoItem, *entity.TodoItem) error
}
