package entity_test

import (
	"testing"
	"todoapp/internal/entity"
)

func TestCheckTodoWhenFalse(t *testing.T) {
	todo := entity.Todo{}

	todo.Check()

	if !todo.IsDone {
		t.Errorf("expected IsDone true, but got %v", todo.IsDone)
	}
}

func TestCheckTodoWhenTrue(t *testing.T) {
	todo := entity.Todo{IsDone: true}

	todo.Check()

	if todo.IsDone {
		t.Errorf("expected IsDone false, but got %v", todo.IsDone)
	}
}
