package usecase_test

import (
	"testing"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestDeleteTodoUsecase(t *testing.T) {
	isDeleteTodoMockCalled := false
	var inputCalled string
	todoRepo := repomock.NewTodoRepoMockBuilder().
		WithDeleteTodoMock(func(s string) error {
			isDeleteTodoMockCalled = true
			inputCalled = s
			return nil
		}).
		Build()
	ctu := usecase.NewDeleteTodoUsecase(todoRepo)
	todoIDToDelete := "1234567890"

	err := ctu.DeleteTodo(todoIDToDelete)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !isDeleteTodoMockCalled {
		t.Errorf("expected repo to call delete todo")
	}
	if isDeleteTodoMockCalled && inputCalled != "1234567890" {
		t.Errorf("expected inputCalled 1234567890, got %s", inputCalled)
	}
}
