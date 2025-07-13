package controllers

import (
	"github.com/gin-gonic/gin"
)

// Register handles user registration
func (s *ControllerService) Profile(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := s.Services.UserService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, profile)
}
