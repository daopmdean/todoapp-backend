package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
)

func TestPingRoute(t *testing.T) {
	userRepo := memdb.NewUserRepo()
	server := rest.NewServer(userRepo)
	router := server.GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200, got %d", w.Code)
	}

	if w.Body.String() != `{"message":"pong"}` {
		t.Errorf(`expected {"message":"pong"}, got %s`, w.Body.String())
	}
}
