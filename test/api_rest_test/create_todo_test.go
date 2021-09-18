package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestCreateTodo(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	tokenStr := internal.GenAccessToken("daopham")
	request := internal.Request{
		Method: "POST",
		Path:   "/api/todos",
		Headers: map[string]string{
			"Content-Type":  "application/json; charset=utf-8",
			"Authorization": "Bearer " + tokenStr,
		},
		Body: `{"content":"Wake up"}`,
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	todoRepo := serverRouter.GetTodoRepo()
	todos := todoRepo.GetTodoListOf("daopham")
	isTodoItemAdded := false
	for _, val := range todos {
		if val.Content == "Wake up" {
			isTodoItemAdded = true
			break
		}
	}

	if !isTodoItemAdded {
		t.Errorf(`expected todo "Wake up" in todo list`)
	}
}
