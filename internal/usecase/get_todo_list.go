package usecase

import (
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"
)

func NewGetTodoListUsecase(todoRepo repo.TodoRepo) *getTodoListUsecase {
	return &getTodoListUsecase{todoRepo}
}

type getTodoListUsecase struct {
	tr repo.TodoRepo
}

func (u *getTodoListUsecase) GetTodoListOf(username string) []entity.Todo {
	return u.tr.GetTodoListOf(username)
}
