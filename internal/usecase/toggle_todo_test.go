package usecase_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestToggleTodoUsecase(t *testing.T) {
	isToggleTodoMockCalled := false
	todoRepo := repomock.NewTodoRepoMockBuilder().
		WithUpdateTodoMock(func(input *entity.Todo) error {
			isToggleTodoMockCalled = true
			return nil
		}).
		Build()
	ctu := usecase.NewToggleTodoUsecase(todoRepo)
	todoIDToToggle := "1234567890"

	err := ctu.ToggleTodo(todoIDToToggle)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !isToggleTodoMockCalled {
		t.Errorf("expected repo to call toggle todo")
	}
	// if isToggleTodoMockCalled & {
	// 	t.Errorf("expected inputCalled 1234567890, got %s", inputCalled)
	// }
}
