package rest

import (
	"net/http"
	"time"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) getTodoList(c *gin.Context) {
	accessToken := s.getWebAccessToken(c)

	authenUsecase := usecase.NewAuthenticationUsecase()
	username, err := authenUsecase.AuthenticateWebAccessToken(accessToken, time.Now())
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	todoUsecase := usecase.NewGetTodoListUsecase(s.todoRepo)

	todos := todoUsecase.GetTodoListOf(username)
	c.JSON(http.StatusOK, todos)
}
