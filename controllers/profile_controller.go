package controllers

import (
	"net/http"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func (s *ControllerService) GetProfile(c *gin.Context) {
	userLogin, err := s.getUserLogin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := s.Services.ProfileService.GetProfileByID(userLogin.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfile handles updating user profile
func (s *ControllerService) UpdateProfile(c *gin.Context) {
	// Get the user ID from the context
	userLogin, err := s.getUserLogin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input dto.UpdateProfileInput

	// Bind the input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := s.Services.ProfileService.UpdateProfile(userLogin.ID, &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// UpdatePasswordProfile handles updating user password
func (s *ControllerService) UpdatePasswordProfile(c *gin.Context) {
	// Get the user ID from the context
	userLogin, err := s.getUserLogin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input dto.UpdatePasswordProfileInput

	// Bind the input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Services.ProfileService.UpdatePassProfile(userLogin.ID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
