package controllers

import (
	"net/http"
	"strconv"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/pkg"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

func (s *ControllerService) ListUsers(c *gin.Context) {
	// Create a Pagination object from query parameters
	limit, err := utils.GetQueryIntGin(c, "limit", 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	page, err := utils.GetQueryIntGin(c, "page", 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	pagination := &pkg.Pagination{
		Limit: limit,
		Page:  page,
	}
	// Call the service to list paginateUsers with pagination
	paginateUsers, err := s.Services.UserService.ListUsers(*pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, paginateUsers)
}

func (s *ControllerService) GetUserByID(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve the user by ID
	user, err := s.Services.UserService.GetUserByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *ControllerService) CreateUser(c *gin.Context) {
	// Bind the input to the struct
	var input dto.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user using the service
	createdUser, err := s.Services.UserService.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

func (s *ControllerService) UpdateUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind the input to the struct
	var input dto.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user using the service
	updatedUser, err := s.Services.UserService.UpdateUser(uint(uid), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User update failed"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (s *ControllerService) UpdatePasswordUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind the input to the struct
	var input dto.UpdatePasswordUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the password using the service
	if err := s.Services.UserService.UpdatePassUser(uint(uid), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
