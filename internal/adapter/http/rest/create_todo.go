package rest

import (
	"net/http"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTodo(c *gin.Context) {
	accessToken := s.getWebAccessToken(c)

	authenUsecase := usecase.NewAuthenticationUsecase()
	username, err := authenUsecase.AuthenticateWebAccessToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	var input usecase.CreateTodoInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind: " + err.Error(),
		})
		return
	}
	input.Username = username

	todoUsecase := usecase.NewCreateTodoUsecase(s.todoRepo)

	err = todoUsecase.CreateTodo(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
