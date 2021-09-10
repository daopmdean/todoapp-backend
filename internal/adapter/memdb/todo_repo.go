package memdb

import (
	"errors"
	"todoapp/internal/adapter/memdb/internal"
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

func (tr *todoRepo) GetTodoByID(id string) *entity.Todo {
	for _, todo := range tr.todos {
		if todo.ID == id {
			return &todo
		}
	}
	return nil
}

func (tr *todoRepo) CreateTodo(td *entity.Todo) error {
	if td.ID == "" {
		td.ID = internal.GenId(10)
	}
	tr.todos = append(tr.todos, *td)
	return nil
}

func (tr *todoRepo) DeleteTodo(todoID string) error {
	for i, todo := range tr.todos {
		if todo.ID == todoID {
			tr.todos = remove(tr.todos, i)
			return nil
		}
	}
	return errors.New("todo not found")
}

func remove(slice []entity.Todo, i int) []entity.Todo {
	return append(slice[:i], slice[i+1:]...)
}

func (tr *todoRepo) ToggleTodo(id string) error {
	for i, todo := range tr.todos {
		if todo.ID == id {
			tr.todos[i].IsDone = !tr.todos[i].IsDone
			return nil
		}
	}
	return errors.New("todo not found")
}
