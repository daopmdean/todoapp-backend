package usecase

import (
	"errors"
	"todoapp/internal/usecase/repo"
)

func NewToggleTodoUsecase(todoRepo repo.TodoRepo) *ToggleTodoUsecase {
	return &ToggleTodoUsecase{todoRepo}
}

type ToggleTodoUsecase struct {
	todoRepo repo.TodoRepo
}

func (uc *ToggleTodoUsecase) ToggleTodo(todoId string) error {
	todo := uc.todoRepo.GetTodoByID(todoId)
	if todo == nil {
		return errors.New("todo not found")
	}

	todo.IsDone = !todo.IsDone

	err := uc.todoRepo.UpdateTodo(todo)
	if err != nil {
		return err
	}

	return nil
}
