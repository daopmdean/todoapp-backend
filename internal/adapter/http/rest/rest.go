package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port int) {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/auth/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	p := fmt.Sprintf(":%v", port)
	router.Run(p)
}
