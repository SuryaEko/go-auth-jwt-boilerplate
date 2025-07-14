package controllers

import (
	"net/http"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func (s *ControllerService) Register(c *gin.Context) {
	var input dto.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.Services.AuthService.Register(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login handles user login
func (s *ControllerService) Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.Services.AuthService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": *token})
}
