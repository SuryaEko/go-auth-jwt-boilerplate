package controllers

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func (s *ControllerService) GetProfile(c *gin.Context) {
	userLogin, err := s.getUserLogin(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := s.Services.ProfileService.GetProfileByID(userLogin.ID)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, profile)
}

// UpdateProfile handles updating user profile
func (s *ControllerService) UpdateProfile(c *gin.Context) {
	// Get the user ID from the context
	userLogin, err := s.getUserLogin(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var input dto.UpdateProfileInput

	// Bind the input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := s.Services.ProfileService.UpdateProfile(userLogin.ID, &input)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, updatedUser)
}

// UpdatePasswordProfile handles updating user password
func (s *ControllerService) UpdatePasswordProfile(c *gin.Context) {
	// Get the user ID from the context
	userLogin, err := s.getUserLogin(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var input dto.UpdatePasswordProfileInput

	// Bind the input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := s.Services.ProfileService.UpdatePassProfile(userLogin.ID, &input); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password updated successfully"})
}
