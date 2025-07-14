package controllers

import (
	"errors"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/services"
	"github.com/gin-gonic/gin"
)

type ControllerService struct {
	Services *services.ServiceContainer
}

func (s *ControllerService) getUserLogin(c *gin.Context) (*models.User, error) {
	userID, exist := c.Get("userID")
	if !exist {
		return nil, errors.New("user ID not found")
	}

	user, err := s.Services.UserService.GetUserByID(userID.(uint))
	if err != nil {
		return nil, err
	}

	return user, nil
}
