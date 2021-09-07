package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTodoItem(c *gin.Context) {
	// todoUsecase := usecase.NewCreateTodoItemUsecase()
	c.JSON(http.StatusOK, gin.H{
		"error": "cannot bind",
	})
}
