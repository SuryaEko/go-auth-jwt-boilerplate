package services

import (
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
func (s *AuthService) Register(user *models.User, pass string) error {
	// set role to "user" by default
	user.Role = "user"

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.DB.Create(user).Error
}

// Login checks if the user exists and returns an error if not
func (s *AuthService) Login(username string, password string) (*string, error) {
	var user models.User

	result := s.DB.Where("username = ?", username).First(&user)

	// If user not found, return an error
	if result.Error != nil {
		return nil, result.Error
	}

	// check password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	// generate jwt token
	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
