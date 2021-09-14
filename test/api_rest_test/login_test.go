package api_rest_test

import (
	"testing"
	"todoapp/test/api_rest_test/internal"
)

func TestLoginSuccess(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	userRepo := serverRouter.GetUserRepo()
	internal.SeedUserData(userRepo)
	request := internal.Request{
		Method: "POST",
		Path:   "/api/auth/login",
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: `{
			"username": "daopham",
			"password": "12345678"
		}`,
	}

	response := serverRouter.HandleRequest(request)

	expectedBody := `{"token":"daopham"}`
	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}
	if string(response.Body) != expectedBody {
		t.Errorf("expected %s, got %s", expectedBody, string(response.Body))
	}
}

func TestLoginFail(t *testing.T) {
	serverRouter := internal.CreateServerRouterForApiTest()
	userRepo := serverRouter.GetUserRepo()
	internal.SeedUserData(userRepo)
	request := internal.Request{
		Method: "POST",
		Path:   "/api/auth/login",
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: `{
			"username": "daopham",
			"password": "12345678a"
		}`,
	}

	response := serverRouter.HandleRequest(request)

	expectedBody := `{"error":"cannot login: username or password incorrect"}`
	if response.Code != 400 {
		t.Errorf("expected 400, got %d", response.Code)
	}
	if string(response.Body) != expectedBody {
		t.Errorf("expected %s, got %s", expectedBody, string(response.Body))
	}
}
