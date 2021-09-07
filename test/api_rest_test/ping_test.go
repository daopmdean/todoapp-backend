package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestPingRoute(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	request := internal.Request{
		Method: "GET",
		Path:   "/api/ping",
	}

	response := router.HandleRequest(request)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	expectedBody := `{"message":"pong"}`
	if string(response.Body) != expectedBody {
		t.Errorf(`expected %s, got %s`, expectedBody, string(response.Body))
	}
}
