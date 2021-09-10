package usecase_test

import (
	"testing"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestToggleTodoUsecase(t *testing.T) {
	isToggleTodoMockCalled := false
	var inputCalled string
	todoRepo := repomock.NewTodoRepoMockBuilder().
		WithToggleTodoMock(func(s string) error {
			isToggleTodoMockCalled = true
			inputCalled = s
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
	if isToggleTodoMockCalled && inputCalled != "1234567890" {
		t.Errorf("expected inputCalled 1234567890, got %s", inputCalled)
	}
}
