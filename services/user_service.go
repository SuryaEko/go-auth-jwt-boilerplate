package services

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user *models.User, pass string) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.DB.Create(user).Error
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

// GetUserByUsername retrieves a user by username
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	result := s.DB.Where("username = ?", username).First(&user)

	// If user not found, return an error
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
