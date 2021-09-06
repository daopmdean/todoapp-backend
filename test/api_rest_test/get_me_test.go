package api_rest_test

import (
	"encoding/json"
	"testing"
	"todoapp/internal/usecase"
	"todoapp/test/api_rest_test/internal"
)

func TestGetMeSuccess(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	method := "GET"
	path := "/api/me"
	headers := map[string]string{
		"Authorization": "Bearer daopham", //TODO: change after implement jwt
	}

	response := internal.RequestServer(router, method, path, "", headers)

	if response.Code != 200 {
		t.Errorf("expected 200, got %d", response.Code)
	}

	result := &usecase.GetMeOutput{}
	err := json.Unmarshal(response.Body.Bytes(), result)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if result.Username != "daopham" {
		t.Errorf("expected username %s, got %s", "daopham", result.Username)
	}

	if result.FirstName != "Dao" {
		t.Errorf("expected firstName %s, got %s", "Dao", result.FirstName)
	}

	if result.LastName != "Pham" {
		t.Errorf("expected lastName %s, got %s", "Pham", result.LastName)
	}
}

func TestGetMeFailWithNoToken(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	method := "GET"
	path := "/api/me"

	response := internal.RequestServer(router, method, path, "", nil)

	if response.Code != 401 {
		t.Errorf("expected 401, got %d", response.Code)
	}
}

func TestGetMeFailWithWrongUsernameFromToken(t *testing.T) {
	router := internal.CreateServerRouterForApiTest()
	method := "GET"
	path := "/api/me"
	headers := map[string]string{
		"Authorization": "Bearer daophamaa", //TODO: change after implement jwt
	}

	response := internal.RequestServer(router, method, path, "", headers)

	if response.Code != 500 {
		t.Errorf("expected 500, got %d", response.Code)
	}
}
