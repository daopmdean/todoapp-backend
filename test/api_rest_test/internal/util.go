package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"

	"github.com/gin-gonic/gin"
)

type serverRouter struct {
	router *gin.Engine
}

type Request struct {
	Method  string
	Path    string
	Body    string
	Headers map[string]string
}

type Response struct {
	Code int
	Body []byte
}

func (sr *serverRouter) HandleRequest(request Request) Response {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(request.Method, request.Path, strings.NewReader(request.Body))
	for key, val := range request.Headers {
		req.Header.Set(key, val)
	}
	sr.router.ServeHTTP(w, req)

	return Response{
		Code: w.Code,
		Body: w.Body.Bytes(),
	}
}

func CreateServerRouterForApiTest() *serverRouter {
	userRepo := memdb.NewUserRepo()
	_ = memdb.NewTodoRepo()
	server := rest.NewServer(userRepo)
	return &serverRouter{router: server.GetRouter()}
}
