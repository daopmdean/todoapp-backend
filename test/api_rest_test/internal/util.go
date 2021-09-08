package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"

	"github.com/gin-gonic/gin"
)

func CreateServerRouterForApiTest() *gin.Engine {
	userRepo := memdb.NewUserRepo()
	_ = memdb.NewTodoRepo()
	server := rest.NewServer(userRepo)
	return server.GetRouter()
}

func RequestServer(router *gin.Engine, method, path, body string, headers map[string]string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	router.ServeHTTP(w, req)

	return w
}
