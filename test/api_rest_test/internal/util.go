package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"todoapp/internal/adapter/http/rest"
	"todoapp/internal/adapter/memdb"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"github.com/gin-gonic/gin"
)

type serverRouter struct {
	router   *gin.Engine
	userRepo repo.UserRepo
	todoRepo repo.TodoRepo
}

func (sr *serverRouter) GetUserRepo() repo.UserRepo {
	return sr.userRepo
}

func (sr *serverRouter) GetTodoRepo() repo.TodoRepo {
	return sr.todoRepo
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
	todoRepo := memdb.NewTodoRepo(userRepo)
	initData(userRepo)

	server := rest.NewServer(userRepo, todoRepo)
	return &serverRouter{
		router:   server.GetRouter(),
		userRepo: userRepo,
		todoRepo: todoRepo,
	}
}

func initData(userRepo repo.UserRepo) {
	dao := entity.User{
		Username:  "daopham",
		FirstName: "Dao",
		LastName:  "Pham",
	}
	dao.SetPassword("12345678")
	hung := entity.User{
		Username:  "hungpham",
		FirstName: "Hung",
		LastName:  "Pham",
	}
	hung.SetPassword("87654321")
	userRepo.SaveUser(&dao)
	userRepo.SaveUser(&hung)
}
