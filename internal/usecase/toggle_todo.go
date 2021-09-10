package usecase

import "todoapp/internal/usecase/repo"

func NewToggleTodoUsecase(todoRepo repo.TodoRepo) *ToggleTodoUsecase {
	return &ToggleTodoUsecase{todoRepo}
}

type ToggleTodoUsecase struct {
	rp repo.TodoRepo
}

func (uc *ToggleTodoUsecase) ToggleTodo(todoId string) error {
	err := uc.rp.ToggleTodo(todoId)
	if err != nil {
		return err
	}
	return nil
}
