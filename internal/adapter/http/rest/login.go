package rest

import (
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) login(c *gin.Context) {
	var input usecase.LoginInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(200, gin.H{
			"error": "cannot bind",
		})
		return
	}

	loginUsecase := usecase.NewLoginUsecase(s.userRepo)

	output, err := loginUsecase.Login(input)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "cannot login: " + err.Error(),
		})
		return
	}

	c.JSON(200, output)
}
