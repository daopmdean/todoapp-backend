package usecase

import "todoapp/internal/usecase/repo"

func NewDeleteTodoUsecase(todoRepo repo.TodoRepo) *DeleteTodoUsecase {
	return &DeleteTodoUsecase{todoRepo}
}

type DeleteTodoUsecase struct {
	todoRepo repo.TodoRepo
}

func (u *DeleteTodoUsecase) DeleteTodo(id string) error {
	err := u.todoRepo.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}
