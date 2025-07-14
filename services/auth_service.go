package services

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/utils"

	// "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

// Register creates a new user in the database
func (s *AuthService) Register(input dto.RegisterInput) (*models.User, error) {
	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Create a new user instance
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     "user", // Default role is 'user'
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Login checks if the user exists and returns an error if not
func (s *AuthService) Login(input dto.LoginInput) (*string, error) {
	var user models.User

	result := s.DB.Where("username = ?", input.Username).First(&user)

	// If user not found, return an error
	if result.Error != nil {
		return nil, result.Error
	}

	// check password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, err
	}

	// generate jwt token
	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
