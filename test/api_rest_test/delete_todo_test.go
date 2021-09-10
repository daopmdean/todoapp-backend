package api_rest_test

import (
	"testing"
	"todoapp/internal/entity"
	"todoapp/test/api_rest_test/internal"
)

func TestDeleteTodo(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	todoRepo := serverRouter.GetTodoRepo()
	todo := &entity.Todo{
		ID:       "1234567890",
		Username: "daopham",
		Content:  "Do abc",
		IsDone:   false,
	}
	todoRepo.CreateTodo(todo)
	request := internal.Request{
		Method: "DELETE",
		Path:   "/api/todos/1234567890",
		Headers: map[string]string{
			"Authorization": "Bearer daopham",
		},
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	todos := todoRepo.GetTodoListOf("daopham")
	if internal.CheckIfTodoListContainsTodo(todos, *todo) {
		t.Errorf("expect todo deleted")
	}
}

func TestUnauthorizedWhenDeleteOtherUserTodo(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	todoRepo := serverRouter.GetTodoRepo()
	todo := &entity.Todo{
		ID:       "1234567890",
		Username: "daopham",
		Content:  "Do abc",
		IsDone:   false,
	}
	todoRepo.CreateTodo(todo)
	request := internal.Request{
		Method: "DELETE",
		Path:   "/api/todos/1234567890",
		Headers: map[string]string{
			"Authorization": "Bearer hungpham",
		},
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 403 {
		t.Errorf("expected 403, got %d", response.Code)
	}
}
