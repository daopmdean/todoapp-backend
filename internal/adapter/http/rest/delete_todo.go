package rest

import (
	"net/http"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) deleteTodo(c *gin.Context) {
	accessToken := s.getWebAccessToken(c)

	authenUsecase := usecase.NewAuthenticationUsecase()
	username, err := authenUsecase.AuthenticateWebAccessToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	todoID := c.Param("id")

	todoOwner := s.todoRepo.GetTodoByID(todoID)
	if todoOwner.Username != username {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	usecase := usecase.NewDeleteTodoUsecase(s.todoRepo)
	err = usecase.DeleteTodo(todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
