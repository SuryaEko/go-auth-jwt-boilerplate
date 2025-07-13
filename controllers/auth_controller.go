package controllers

import (
	"net/http"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func (s *ControllerService) Register(c *gin.Context) {
	var input struct {
		Username   string `json:"username" binding:"required,min=2,max=100"`
		Password   string `json:"password" binding:"required,min=6"`
		RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
	}

	if err := s.Services.UserService.CreateUser(&user, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login handles user login
func (s *ControllerService) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.Services.AuthService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": *token})
}
