package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestLoginSuccess(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	method := "POST"
	path := "/api/auth/login"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := `{
		"username": "daopham",
		"password": "12345678"
	}`

	response := internal.RequestServer(router, method, path, body, headers)

	expectedBody := `{"token":"daopham"}`
	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}
	if response.Body.String() != expectedBody {
		t.Errorf("expected %s, got %s", expectedBody, response.Body.String())
	}
}
