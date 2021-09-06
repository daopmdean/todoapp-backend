package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestPingRoute(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	method := "GET"
	path := "/api/ping"

	response := internal.RequestServer(router, method, path, "", nil)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	expectedBody := `{"message":"pong"}`
	if response.Body.String() != expectedBody {
		t.Errorf(`expected %s, got %s`, expectedBody, response.Body.String())
	}
}
