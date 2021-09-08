package usecase_test

import (
	"strconv"
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func TestCreateTodo(t *testing.T) {
	hasCreateTodoFuncCalled := false
	var inputCalled entity.Todo
	todoRepo := repomock.NewTodoRepoMockBuilder().
		WithCreateTodoMock(func(t *entity.Todo) error {
			hasCreateTodoFuncCalled = true
			inputCalled = *t
			return nil
		}).
		Build()
	ctu := usecase.NewCreateTodoUsecase(todoRepo)
	input := usecase.CreateTodoInput{
		Username: "daopham",
		Content:  "Eat breakfast",
	}

	err := ctu.CreateTodo(input)

	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}

	if !hasCreateTodoFuncCalled {
		t.Errorf("expected repo's func createTodo called")
	}

	if hasCreateTodoFuncCalled {
		tests := []struct {
			expected string
			got      string
		}{
			{"daopham", inputCalled.Username},
			{"Eat breakfast", inputCalled.Content},
			{"false", strconv.FormatBool(inputCalled.IsDone)},
		}

		for _, val := range tests {
			if val.expected != val.got {
				t.Errorf("expected %s, got %s", val.expected, val.got)
			}
		}
	}
}
