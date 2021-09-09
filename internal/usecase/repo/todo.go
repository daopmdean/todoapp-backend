package repo

import "todoapp/internal/entity"

type TodoRepo interface {
	GetTodoListOf(username string) []entity.Todo
	GetTodoByID(id string) *entity.Todo
	CreateTodo(*entity.Todo) error
	DeleteTodo(todoID string) error
}
