package rest

import (
	"net/http"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) toggleTodo(c *gin.Context) {
	accessToken := s.getWebAccessToken(c)

	authenUsecase := usecase.NewAuthenticationUsecase()
	username, err := authenUsecase.AuthenticateWebAccessToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	authorUsecase := usecase.NewAuthorizationUsecase(s.todoRepo)
	todoID := c.Param("id")
	isAuthorize := authorUsecase.CheckIfUserIsAllowedToModifyTodo(username, todoID)
	if !isAuthorize {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	usecase := usecase.NewToggleTodoUsecase(s.todoRepo)
	err = usecase.ToggleTodo(todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
