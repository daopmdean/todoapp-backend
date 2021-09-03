package rest

import (
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) register(c *gin.Context) {
	var input usecase.RegisterInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(200, gin.H{
			"error": "cannot bind",
		})
		return
	}

	registerUsecase := usecase.NewRegisterUsecase(s.userRepo)

	if err := registerUsecase.Register(input); err != nil {
		c.JSON(200, gin.H{
			"error": "cannot register: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
