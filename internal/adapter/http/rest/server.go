package rest

import (
	"fmt"
	"todoapp/internal/usecase/repo"

	"github.com/gin-gonic/gin"
)

func NewServer(ur repo.UserRepo) *Server {
	server := &Server{userRepo: ur}
	server.setupRouter()
	return server
}

type Server struct {
	router   *gin.Engine
	userRepo repo.UserRepo
}

func (s *Server) setupRouter() {
	router := gin.Default()
	s.setupApi(router)
	s.router = router
}

func (s *Server) setupApi(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/auth/register", s.register)
	api.POST("/auth/login", s.login)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

func (s *Server) Run(port int) {
	p := fmt.Sprintf(":%v", port)
	s.router.Run(p)
}
