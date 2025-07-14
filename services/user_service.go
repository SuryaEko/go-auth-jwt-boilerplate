package services

import (
	"errors"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/utils"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(input *dto.CreateUserInput) (*models.User, error) {
	// validate unique username
	var existingUser models.User
	if err := s.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Create a new user instance
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role, // Use the role from the input
	}

	// Save the user to the database
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	result := s.DB.First(&user, id)

	// If user not found, return an error
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// updateUser updates an existing user
func (s *UserService) UpdateUser(id uint, userInput *dto.UpdateUserInput) (*models.User, error) {
	existingUser, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if err := s.DB.Model(existingUser).Updates(userInput).Error; err != nil {
		return nil, err
	}

	return existingUser, nil
}

// updatePassUser updates the password of an existing user
func (s *UserService) UpdatePassUser(id uint, input *dto.UpdatePasswordUserInput) error {
	existingUser, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	// Hash the new password before saving
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return err
	}

	data := map[string]any{
		"password": string(hashedPassword),
	}

	return s.DB.Model(existingUser).Updates(data).Error
}
