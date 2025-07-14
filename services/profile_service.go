package services

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/dto"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/utils"
	"gorm.io/gorm"
)

type ProfileService struct {
	DB *gorm.DB
}

// GetProfileByID retrieves a user profile by ID
func (s *ProfileService) GetProfileByID(id uint) (*models.User, error) {
	var user models.User

	result := s.DB.First(&user, id)

	// If user not found, return an error
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// UpdateProfile updates an existing user profile
func (s *ProfileService) UpdateProfile(id uint, input *dto.UpdateProfileInput) (*models.User, error) {
	existingUser, err := s.GetProfileByID(id)
	if err != nil {
		return nil, err
	}

	// Update the user fields with the input data
	existingUser.Username = input.Username

	if err := s.DB.Save(existingUser).Error; err != nil {
		return nil, err
	}

	return existingUser, nil
}

// updatePassProfile updates the password of an existing user
func (s *ProfileService) UpdatePassProfile(id uint, input *dto.UpdatePasswordProfileInput) error {
	existingUser, err := s.GetProfileByID(id)
	if err != nil {
		return err
	}
	// Verify the old password
	if err := utils.IsPasswordValid(existingUser.Password, input.OldPassword); err != nil {
		return err // Old password does not match
	}

	// Hash the new password before saving
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return err
	}
	existingUser.Password = string(hashedPassword)

	return s.DB.Save(existingUser).Error
}
