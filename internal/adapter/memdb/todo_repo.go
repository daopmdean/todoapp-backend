package memdb

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewTodoRepo(ur repo.UserRepo) repo.TodoRepo {
	return &todoRepo{
		userRepo: ur,
	}
}

type todoRepo struct {
	userRepo repo.UserRepo
}

func (tr *todoRepo) GetTodoList(username string) []entity.TodoItem {

	return nil
}

func (tr *todoRepo) CreateTodoItem(tds []entity.TodoItem, td *entity.TodoItem) error {
	tds = append(tds, *td)
	return nil
}

func (tr *todoRepo) CheckTodoItem(tds []entity.TodoItem, td *entity.TodoItem) error {
	return nil
}

func (tr *todoRepo) DeleteTodoItem(tds []entity.TodoItem, td *entity.TodoItem) error {
	return nil
}
