package usecase_test

import (
	"strconv"
	"testing"
	"todoapp/internal/entity"
	"todoapp/internal/usecase"
	"todoapp/test/repomock"
)

func entitiesReturn() []entity.Todo {
	return []entity.Todo{
		{
			Username: "daopham",
			Content:  "Wash hand",
			IsDone:   false,
		},
		{
			Username: "daopham",
			Content:  "Brush teeth",
			IsDone:   false,
		},
	}
}

func TestGetTodoList(t *testing.T) {
	hasGetTodoListFuncCalled := false
	todoRepo := repomock.NewTodoRepoMockBuilder().
		WithGetTodoListOfMock(func(s string) []entity.Todo {
			hasGetTodoListFuncCalled = true
			return entitiesReturn()
		}).
		Build()
	ctu := usecase.NewGetTodoListUsecase(todoRepo)

	todos := ctu.GetTodoListOf("daopham")

	if !hasGetTodoListFuncCalled {
		t.Errorf("expect to call func GetToListOf")
	}

	if hasGetTodoListFuncCalled {
		tests := []struct {
			expected string
			got      string
		}{
			{"daopham", todos[0].Username},
			{"Wash hand", todos[0].Content},
			{"false", strconv.FormatBool(todos[0].IsDone)},
			{"daopham", todos[1].Username},
			{"Brush teeth", todos[1].Content},
			{"false", strconv.FormatBool(todos[1].IsDone)},
		}

		for _, val := range tests {
			if val.expected != val.got {
				t.Errorf("expected %s, got %s", val.expected, val.got)
			}
		}
	}
}
