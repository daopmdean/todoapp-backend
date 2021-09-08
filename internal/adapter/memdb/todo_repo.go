package memdb

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

// check if missing implementation
var _ repo.TodoRepo = new(todoRepo)

func NewTodoRepo() repo.TodoRepo {
	return &todoRepo{
		todos: []entity.Todo{},
	}
}

type todoRepo struct {
	todos []entity.Todo
}

func (tr *todoRepo) GetTodoListOf(username string) []entity.Todo {
	result := []entity.Todo{}
	for _, todo := range tr.todos {
		if todo.Username == username {
			result = append(result, todo)
		}
	}
	return result
}

func (tr *todoRepo) CreateTodo(td *entity.Todo) error {
	tr.todos = append(tr.todos, *td)
	return nil
}

func (tr *todoRepo) CheckTodo(td *entity.Todo) error {
	return nil
}

func (tr *todoRepo) DeleteTodo(td *entity.Todo) error {
	return nil
}
