package rest

import (
	"fmt"
	"todoapp/internal/usecase/repo"

	"github.com/gin-gonic/gin"
)

func NewServer(ur repo.UserRepo) *Server {
	return &Server{ur}
}

type Server struct {
	userRepo repo.UserRepo
}

func (s *Server) Run(port int) {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/auth/register", s.register)

	api.POST("/auth/login", s.login)

	p := fmt.Sprintf(":%v", port)
	router.Run(p)
}