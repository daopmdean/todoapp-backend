package repo

import "todoapp/internal/entity"

type TodoRepo interface {
	GetTodoList(username string) []entity.Todo
	CreateTodo(*entity.Todo) error
	CheckTodo(*entity.Todo) error
	DeleteTodo(*entity.Todo) error
}
