package rest

import (
	"net/http"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) login(c *gin.Context) {
	var input usecase.LoginInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind",
		})
		return
	}

	loginUsecase := usecase.NewLoginUsecase(s.userRepo)

	output, err := loginUsecase.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot login: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}
