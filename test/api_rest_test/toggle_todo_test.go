package api_rest_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/test/api_rest_test/internal"
)

func TestToggleTodo(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	todoRepo := serverRouter.GetTodoRepo()
	newTodo := &entity.Todo{
		ID:       "1234567890",
		Username: "daopham",
		Content:  "Do abc",
		IsDone:   false,
	}
	todoRepo.CreateTodo(newTodo)
	tokenStr := internal.GenAccessToken("daopham")
	request := internal.Request{
		Method: "PUT",
		Path:   "/api/todos/1234567890/toggle",
		Headers: map[string]string{
			"Authorization": "Bearer " + tokenStr,
		},
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	todo := todoRepo.GetTodoByID("1234567890")
	if !todo.IsDone {
		t.Errorf("expect todo IsDone true, got %v", todo.IsDone)
	}
}

func TestUnauthorizedWhenToggleOtherUserTodo(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	todoRepo := serverRouter.GetTodoRepo()
	todo := &entity.Todo{
		ID:       "1234567890",
		Username: "daopham",
		Content:  "Do abc",
		IsDone:   false,
	}
	todoRepo.CreateTodo(todo)
	tokenStr := internal.GenAccessToken("hungpham")
	request := internal.Request{
		Method: "PUT",
		Path:   "/api/todos/1234567890/toggle",
		Headers: map[string]string{
			"Authorization": "Bearer " + tokenStr,
		},
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 403 {
		t.Errorf("expected 403, got %d", response.Code)
	}
}
