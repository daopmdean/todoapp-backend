package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestCreateTodoItem(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	method := "POST"
	path := "/api/todos"
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "daopham",
	}
	body := `{"content":"Wake up"}`

	response := internal.RequestServer(router, method, path, body, headers)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}
}
