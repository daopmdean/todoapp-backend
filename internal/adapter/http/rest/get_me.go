package rest

import (
	"net/http"
	"time"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func (s *Server) getMe(c *gin.Context) {
	accessToken := s.getWebAccessToken(c)

	authenUsecase := usecase.NewAuthenticationUsecase()
	username, err := authenUsecase.AuthenticateWebAccessToken(accessToken, time.Now())
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	getMeUsecase := usecase.NewGetMeUsecase(s.userRepo, username)
	getMeOutput, err := getMeUsecase.GetMe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, getMeOutput)
}
