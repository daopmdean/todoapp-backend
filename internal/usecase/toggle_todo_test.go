package usecase_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestToggleTodoUsecase(t *testing.T) {
	isUpdateTodoMockCalled := false
	var inputTodo entity.Todo
	todoRepo := repomock.NewTodoRepoMockBuilder().
		WithGetTodoByID(func(s string) *entity.Todo {
			return &entity.Todo{
				ID: "1234567890",
			}
		}).
		WithUpdateTodoMock(func(input *entity.Todo) error {
			isUpdateTodoMockCalled = true
			inputTodo = *input
			return nil
		}).
		Build()
	ctu := usecase.NewToggleTodoUsecase(todoRepo)
	todoIDToToggle := "1234567890"

	err := ctu.ToggleTodo(todoIDToToggle)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !isUpdateTodoMockCalled {
		t.Errorf("expected repo to call update todo")
	}
	if isUpdateTodoMockCalled && inputTodo.ID != "1234567890" {
		t.Errorf("expected inputCalled 1234567890, got %s", inputTodo.ID)
	}
}
