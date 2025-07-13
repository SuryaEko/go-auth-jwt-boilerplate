package services

import (
	"gorm.io/gorm"
)

type ServiceContainer struct {
	AuthService *AuthService
	UserService *UserService
	// Add other services here as needed
}

func InitServiceContainer(db *gorm.DB) *ServiceContainer {
	return &ServiceContainer{
		AuthService: &AuthService{DB: db},
		UserService: &UserService{DB: db},
		// Initialize other services here
	}
}
