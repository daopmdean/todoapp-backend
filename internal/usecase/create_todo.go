package usecase

import (
	"errors"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"github.com/google/uuid"
)

func NewCreateTodoUsecase(tr repo.TodoRepo) *CreateTodoUsecase {
	return &CreateTodoUsecase{tr}
}

type CreateTodoUsecase struct {
	todoRepo repo.TodoRepo
}

func (ctiu *CreateTodoUsecase) CreateTodo(input CreateTodoInput) error {
	todo := entity.Todo{
		ID:       uuid.New().String(),
		Username: input.Username,
		Content:  input.Content,
		IsDone:   false,
	}

	err := ctiu.todoRepo.CreateTodo(&todo)
	if err != nil {
		return errors.New("cannot create todo: " + err.Error())
	}

	return nil
}

type CreateTodoInput struct {
	Username string
	Content  string `json:"content"`
}
