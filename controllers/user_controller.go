package controllers

import (
	"strconv"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/gin-gonic/gin"
)

func (s *ControllerService) GetUserByID(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve the user by ID
	user, err := s.Services.UserService.GetUserByID(uint(uid))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func (s *ControllerService) CreateUser(c *gin.Context) {
	// Bind the input to the struct
	var input dto.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create the user using the service
	createdUser, err := s.Services.UserService.CreateUser(&input)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, createdUser)
}

func (s *ControllerService) UpdateUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind the input to the struct
	var input dto.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the user using the service
	updatedUser, err := s.Services.UserService.UpdateUser(uint(uid), &input)
	if err != nil {
		c.JSON(404, gin.H{"error": "User update failed"})
		return
	}
	c.JSON(200, updatedUser)
}

func (s *ControllerService) UpdatePasswordUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind the input to the struct
	var input dto.UpdatePasswordUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the password using the service
	if err := s.Services.UserService.UpdatePassUser(uint(uid), &input); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password updated successfully"})
}
