package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestCreateTodoItem(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	request := internal.Request{
		Method: "POST",
		Path:   "/api/todos",
		Headers: map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "daopham",
		},
		Body: `{"content":"Wake up"}`,
	}

	response := serverRouter.HandleRequest(request)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}
}
